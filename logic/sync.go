package logic

import (
	"encoding/json"
	"github.com/termermc/your-loss-sync/config"
	"github.com/termermc/your-loss-sync/util"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

type ffprobeResult struct {
	Streams []struct {
		CodecType string `json:"codec_type"` // We're looking for "audio"
		CodecName string `json:"codec_name"`
	} `json:"streams"`
}

func doFfprobe(bin string, filePath string) (ffprobeResult, error) {
	cmd := exec.Command(
		bin,
		"-print_format", "json",
		"-show_streams",
		filePath,
	)
	out, err := cmd.Output()
	if err != nil {
		return ffprobeResult{}, err
	}

	var res ffprobeResult
	err = json.Unmarshal(out, &res)
	if err != nil {
		return ffprobeResult{}, err
	}

	return res, nil
}

var audioExtensions = []string{
	"mp3",
	"flac",
	"wav",
	"opus",
	"m4a",
	"aif",
	"alac",
	"ape",
	"wma",
	"ogg",
}

// StartSync starts a sync.
// This function blocks until the sync is complete.
func StartSync(s *AppState, sync *config.SyncConfig, logOut chan string) {
	s.Progress.Sync.Store(sync)
	s.Progress.Completed.Store(0)
	s.Progress.Total.Store(0)
	s.Progress.Failed.Store(0)

	checkErr := func(err error) bool {
		if err == nil {
			return false
		}

		logOut <- s.Locale.Tr("general.error") + ": " + s.Locale.TrError(err)
		s.Progress.Failed.Add(1)
		return true
	}

	// TODO FFmpeg setting
	ffmpegBin := "ffmpeg"
	ffprobeBin := "ffprobe"

	logOut <- s.Locale.Tr("sync.scanning-source")

	srcPath := sync.SourceDir
	if !strings.HasSuffix(srcPath, "/") {
		srcPath += "/"
	}
	destPath := sync.DestDir
	if !strings.HasSuffix(destPath, "/") {
		destPath += "/"
	}

	// Assume that transcoding a file maxes out a single CPU thread
	concurrency := runtime.NumCPU()
	if concurrency < 1 {
		concurrency = 1
	}

	fileChan := make(chan string, 100_000) // Arbitrarily large buffer since we want to calculate the total number of files
	doneChan := make(chan struct{}, concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer func() {
				doneChan <- struct{}{}
			}()

			for srcFilePathFull := range fileChan {
				if s.Progress.Sync.Load() != sync {
					// Sync has been canceled
					return
				}

				fileRelative := srcFilePathFull[len(srcPath):]

				pathParts := strings.Split(fileRelative, string(os.PathSeparator))

				// Escape parts of the srcFilePathFull path if specified
				if sync.EscapeFilenames {
					for i, _ := range pathParts {
						pathParts[i] = util.EscapeFilename(pathParts[i])
					}

					fileRelative = strings.Join(pathParts, string(os.PathSeparator))
				}

				fnameFull := pathParts[len(pathParts)-1]
				ext := filepath.Ext(fnameFull)
				fnameNoExt := fnameFull[:len(fnameFull)-len(ext)]

				// Make dirs
				if len(pathParts) > 1 {
					err := os.MkdirAll(filepath.Join(destPath, filepath.Dir(fileRelative)), os.ModePerm)
					if checkErr(err) {
						continue
					}
				}

				shouldCopyRaw := true

				// Check if the srcFilePathFull is a supported audio srcFilePathFull
				if ext != "" && slices.Contains(audioExtensions, strings.ToLower(ext[1:])) {
					prof := sync.Profile

					destFilePath := filepath.Join(destPath, filepath.Dir(fileRelative), fnameNoExt+"."+prof.OutputFormat.Extension)
					destTmpPath := destFilePath + ".tmp." + prof.OutputFormat.Extension

					// Check if it already exists
					_, err := os.Stat(destFilePath)
					if err == nil {
						println(s.Locale.Tr("sync.path-already-exists", fileRelative))
						s.Progress.Completed.Add(1)
						continue
					}

					// Probe the file
					res, err := doFfprobe(ffprobeBin, srcFilePathFull)
					if checkErr(err) {
						continue
					}

					// Check for audio stream
					var audioFmt string
					for _, stream := range res.Streams {
						if stream.CodecType == "audio" {
							audioFmt = stream.CodecName
							break
						}
					}

					if audioFmt == "" {
						// No audio stream, copy the file
						shouldCopyRaw = true
					} else if !sync.ReencodeSameFormat && strings.Contains(sync.Profile.OutputFormat.FfmpegEncoder, audioFmt) {
						// Reencoding is disabled and the audio format matches the output format, copy the file
						shouldCopyRaw = true
					} else {
						// The file needs to be encoded
						shouldCopyRaw = false

						println(s.Locale.Tr("sync.transcoding", fileRelative))

						// Run FFmpeg
						cmd := exec.Command(
							ffmpegBin,
							"-i", srcFilePathFull,
							"-c:v", "copy",
							"-c:a", prof.OutputFormat.FfmpegEncoder,
							"-b:a", strconv.Itoa(int(prof.OutputFormat.SuggestedBitrate)),
							destTmpPath,
							"-y",
						)
						err = cmd.Run()
						if checkErr(err) {
							_ = os.Remove(destTmpPath)

							continue
						}

						// Successfully transcoded, rename the tmp file
						err = os.Rename(destTmpPath, destFilePath)
						if checkErr(err) {
							_ = os.Remove(destTmpPath)

							continue
						}

						s.Progress.Completed.Add(1)
					}
				}

				if shouldCopyRaw {
					destFilePath := filepath.Join(destPath, fileRelative)
					destTmpPath := destFilePath + ".tmp"

					// Check if it already exists
					_, err := os.Stat(destFilePath)
					if err == nil {
						println(s.Locale.Tr("sync.path-already-exists", fileRelative))
						s.Progress.Completed.Add(1)
						continue
					}

					println(s.Locale.Tr("sync.copying", fileRelative))

					// Simply copy the file
					osSrcFile, err := os.Open(srcFilePathFull)
					if checkErr(err) {
						_ = os.Remove(destTmpPath)
						_ = osSrcFile.Close()
						continue
					}
					osDestFile, err := os.Create(destTmpPath)
					if checkErr(err) {
						_ = os.Remove(destTmpPath)
						_ = osSrcFile.Close()
						_ = osDestFile.Close()
						continue
					}
					_, err = io.Copy(osDestFile, osSrcFile)
					if checkErr(err) {
						_ = os.Remove(destTmpPath)
						_ = osSrcFile.Close()
						_ = osDestFile.Close()
						continue
					}

					_ = osDestFile.Close()
					_ = osSrcFile.Close()

					// Successfully copied, rename the tmp file
					err = os.Rename(destTmpPath, destFilePath)
					if checkErr(err) {
						_ = os.Remove(destTmpPath)
						continue
					}

					s.Progress.Completed.Add(1)
				}
			}
		}()
	}

	// Walk source directory
	err := filepath.WalkDir(srcPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		s.Progress.Total.Add(1)
		fileChan <- path

		return nil
	})
	close(fileChan)
	if checkErr(err) {
		return
	}

	// Wait for all processes to finish
	for i := 0; i < concurrency; i++ {
		<-doneChan
	}

	logOut <- s.Locale.Tr(
		"sync.done",
		strconv.Itoa(int(s.Progress.Total.Load())),
		strconv.Itoa(int(s.Progress.Completed.Load())),
		strconv.Itoa(int(s.Progress.Failed.Load())),
	)
	s.Progress.Sync.Store(nil)
}

package ffmpeg

// Ffmpeg manages execution of FFmpeg and FFprobe
type Ffmpeg struct {
	// Path to the `ffmpeg` binary
	FfmpegPath string

	// Path to the `ffprobe` binary
	FfprobePath string
}

// New returns a new Ffmpeg instance
func New(ffmpegPath, ffprobePath string) Ffmpeg {
	return Ffmpeg{
		FfmpegPath:  ffmpegPath,
		FfprobePath: ffprobePath,
	}
}

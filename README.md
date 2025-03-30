# Your Loss! Sync
Sync and convert your music library easily

# About

Your Loss! Sync (YLS) is a tool that allows you to sync and convert your music library from a master source to multiple destinations, such as external drives, SD cards, or USBs.

Each "sync" is a set of rules for a source and destination that can be used to copy and convert the source's music.

Let's say you have a music library on your computer with a mixture of lossless FLAC and lossy MP3s. Normally, you would have to manually convert the FLAC files to MP3 to play them
on an MP3 player or use them with DJ equipment. YLS allows you to set up a sync that will automatically sync your source to your external drive, converting the FLACs to MP3s and
copying the existing MP3s unchanged to the external drive, all while keeping the same file names and metadata.

Adding new music to a source is no longer a hassle, as YLS will automatically detect new files and sync (and convert) them to the destination.

Unlike most other tools, YLS normalizes file names to support all file systems, including FAT32 and exFAT.
Often file names on MacOS and Linux are not compatible with Windows/MS-DOS file systems, so traditional syncing and conversion tools will often fail when encountering files with
unusual names. With YLS, you can be sure that your music will have no trouble being synced to a variety of storage devices.

I designed YLS to allow me to sync my large collection of lossless music to external drives and SD cards for DJing. Most CDJs made before 2020
don't support FLAC, so I previously needed to manually convert my FLACs to MP3s before exporting to USBs from Rekordbox.

# Download

You can download the latest release for Windows, Mac and Linux from the [releases page](https://github.com/termermc/your-loss-sync/releases).

# Features

 - Sync music to a destination
 - Convert music while syncing (supports a variety of formats)
 - Automatically detect new music
 - Skip files larger than a certain size (useful when the destination is on a FAT32 filesystem)
 - Normalize and sanitize file names to support all file systems
 - Reduce size of destination files by reducing bitrate
 - Re-encode files for a lower bitrate

# Building

## Linux

You need some GTK dependencies.
TODO Figure out which ones

Run `go build -o your-loss-sync cmd/main.go`

## Windows

TODO go build

## MacOS

TODO go build


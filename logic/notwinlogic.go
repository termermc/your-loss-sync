//go:build !windows

package logic

import "syscall"

func GetFFmpegSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{}
}

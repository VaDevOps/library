package info

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"syscall"
)

func Log(json, debug, color bool) {
	var handler slog.Handler
	if json {
		if debug {
			handler = slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
		} else {
			handler = slog.NewJSONHandler(os.Stderr, nil)
		}
	} else {
		if color {
			if debug {
				handler = tint.NewHandler(os.Stderr, &tint.Options{Level: slog.LevelDebug})
			} else {
				handler = tint.NewHandler(os.Stderr, nil)
			}
		} else {
			if debug {
				handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
			} else {
				handler = slog.NewTextHandler(os.Stderr, nil)
			}
		}
	}
	slog.SetDefault(slog.New(handler))
}

func CPU() string{
	return strconv.Itoa(runtime.NumCPU())
}

func Memory(sysinfoFunc func(*syscall.Sysinfo_t) error) (string, error) {
	var info syscall.Sysinfo_t
	err := sysinfoFunc(&info)
	if err != nil {
		return "", err
	}

	totalMemory := uint64(info.Totalram) * uint64(info.Unit)
	return strconv.FormatUint(totalMemory/1024/1024, 10), nil
}

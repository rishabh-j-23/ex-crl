package app

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Logger *slog.Logger

var (
	primaryLogDir = "/var/log/ex-crl"
	logDir        = ""
)

func getFallbackLogDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "./logs"
	}
	return filepath.Join(home, "ex-crl", "logs")
}

// InitLogger tries primary, then fallback, only prints to console if both fail
func InitLogger(debug bool) error {
	logDir = primaryLogDir
	if err := os.MkdirAll(logDir, 0755); err != nil {
		// Try fallback
		logDir = getFallbackLogDir()
		if err2 := os.MkdirAll(logDir, 0755); err2 != nil {
			// Only now print to console
			os.Stderr.WriteString("Failed to initialize logger: " + err2.Error() + "\n")
			return err2
		}
	}

	today := time.Now().Format("2006-01-02")
	logFileName := fmt.Sprintf("ex-crl_%s.log", today)
	logPath := filepath.Join(logDir, logFileName)

	// Rotate and clean up logs
	if err := rotateAndCleanupLogs(today); err != nil {
		return err
	}

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}

	handlers := []slog.Handler{slog.NewTextHandler(f, &slog.HandlerOptions{Level: slog.LevelInfo})}
	if debug {
		handlers = append(handlers, slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	var h slog.Handler
	if len(handlers) == 1 {
		h = handlers[0]
	} else {
		h = &multiHandler{handlers: handlers}
	}
	Logger = slog.New(h)
	slog.SetDefault(Logger)
	return nil
}

// --- End of logger setup ---

type multiHandler struct {
	handlers []slog.Handler
}

func (m *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *multiHandler) Handle(ctx context.Context, r slog.Record) error {
	var err error
	for _, h := range m.handlers {
		e := h.Handle(ctx, r)
		if e != nil {
			err = e
		}
	}
	return err
}

func (m *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		next[i] = h.WithAttrs(attrs)
	}
	return &multiHandler{handlers: next}
}

func (m *multiHandler) WithGroup(name string) slog.Handler {
	next := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		next[i] = h.WithGroup(name)
	}
	return &multiHandler{handlers: next}
}

func rotateAndCleanupLogs(today string) error {
	d, err := os.Open(logDir)
	if err != nil {
		return err
	}
	files, err := d.Readdirnames(-1)
	_ = d.Close()
	if err != nil {
		return err
	}
	now := time.Now()
	for _, name := range files {
		if strings.HasSuffix(name, ".log") && !strings.Contains(name, today) {
			// Compress old log
			logPath := filepath.Join(logDir, name)
			gzPath := logPath + ".gz"
			if err := compressFile(logPath, gzPath); err != nil {
				return err
			}
			_ = os.Remove(logPath)
		}
		if strings.HasSuffix(name, ".gz") {
			// Delete if older than 7 days
			fi, err := os.Stat(filepath.Join(logDir, name))
			if err == nil && now.Sub(fi.ModTime()) > 7*24*time.Hour {
				_ = os.Remove(filepath.Join(logDir, name))
			}
		}
	}
	return nil
}

func compressFile(src, dst string) error {
	srcF, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcF.Close()
	dstF, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstF.Close()
	gz := gzip.NewWriter(dstF)
	defer gz.Close()
	_, err = io.Copy(gz, srcF)
	return err
}

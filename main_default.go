//go:build !systemd
// +build !systemd

package main

import (
	"github.com/nxadm/tail"
	"log/slog"
	"log/syslog"
	"os"
)

func (e *Exporter) JournalTail(identifier string, priority syslog.Priority) chan *tail.Line {
	_ = identifier
	_ = priority
	slog.Error("Not compiled with systemd support. (-tags systemd)")
	os.Exit(1)
	return make(chan *tail.Line)
}

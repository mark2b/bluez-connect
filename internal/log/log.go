package log

import (
	"github.com/op/go-logging"
	"os"
)

var (
	Log = logging.MustGetLogger("blue-connect")
)

func SetOutputMode(debug bool, verbose bool, systemd bool) {
	if systemd {
		if verbose {
			initialize(systemdModeFormatter(), logging.DEBUG)
		} else {
			initialize(systemdModeFormatter(), logging.INFO)
		}
	} else if debug {
		if verbose {
			initialize(debugModeFormatter(), logging.DEBUG)
		} else {
			initialize(debugModeFormatter(), logging.INFO)
		}
	} else {
		if verbose {
			initialize(defaultModeFormatter(), logging.DEBUG)
		} else {
			initialize(defaultModeFormatter(), logging.INFO)
		}
	}
}

func initialize(formatter logging.Formatter, level logging.Level) {
	backend := logging.NewLogBackend(os.Stdout, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, formatter)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(level, "")
}

func init() {
	SetOutputMode(false, false, false)
}

func debugModeFormatter() logging.Formatter {
	return logging.MustStringFormatter(
		`%{color}%{level:.1s} %{module} %{time:15:04:05.000} %{shortfile} %{message}%{color:reset}`,
	)
}

func defaultModeFormatter() logging.Formatter {
	return logging.MustStringFormatter(
		`%{level:.1s} %{module} %{time:15:04:05.000} %{message}`,
	)
}

func systemdModeFormatter() logging.Formatter {
	return logging.MustStringFormatter(
		`%{level:.1s} %{message}`,
	)
}

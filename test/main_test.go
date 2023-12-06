package test

import (
	"logs"
	"testing"
)

func TestLogs(t *testing.T) {
	logger, err := logs.NewLogger(logs.InfoLevel, logs.LogOutputToTerminal())
	if err != nil {
		t.Error(err)
		return
	}
	logger.Info("info")
	logger.Debug("debug")
}

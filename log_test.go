package teapot

import "testing"

func TestDebug(t *testing.T) {
	log := New()
	log.Debug("hello world")
	log.Debug("another line")
}

func TestInfo(t *testing.T) {
	log := New()
	log.Info("hello world")
	log.Info("another line")
}

func TestError(t *testing.T) {
	log := New()
	log.Error("hello world")
	log.Error("another line")
}

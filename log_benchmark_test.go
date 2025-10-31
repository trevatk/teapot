package teapot

import (
	"io"
	"testing"
)

var bhLogger = New()

func init() {
	bhLogger.writer = io.Discard
	bhLogger.lvl = INFO
}

func BenchmarkLogger_InfoHotPath(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bhLogger.Info("Request started: user_id=%d, service=%s, elapsed=%s", 12345, "checkout", "25ms")
	}
}

package sm_yamux

import (
	"testing"

	test "github.com/dms3-p2p/go-stream-muxer/test"
)

func TestYamuxTransport(t *testing.T) {
	test.SubtestAll(t, DefaultTransport)
}

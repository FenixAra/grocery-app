package log

import (
	"testing"

	"github.com/RealImage/jt-utils/testHelpers"
)

func TestLogger(t *testing.T) {
	l := NewLogger("")
	testHelpers.AssertNotEqual(nil, l, t)
	l.Debug("This is Debug log")
	l.Error("This is error log")
	l.Fatal("This is fatal log")
	l.Info("This is info log")
	l.Warn("This is warn log")
}

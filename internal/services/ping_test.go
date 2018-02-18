package services

import (
	"testing"

	"github.com/FenixAra/grocery-app/internal/fixtures/test_helpers"
	"github.com/RealImage/jt-utils/testHelpers"
)

func TestPing(t *testing.T) {
	p := NewPing(test_helpers.TestInit())

	err := p.Ping()
	testHelpers.AssertNoError(err, t, "Unable to call ping service")
}

package response
 
import (
	"testing"
	"github.com/roxyash/go_ahead/pkg/zaplogger"
)
 
func TestNewResponse(t *testing.T) {
	logger := zaplogger.NewZapLogger("", "")

	resp := NewResponse(logger)

	if resp == nil {
		t.Errorf("Expected response to be not nil")
	}

	if resp.logger != logger {
		t.Errorf("Expected logger to be equal")
	}
}
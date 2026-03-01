package organization_test

import (
	"go-backend-stream/internal/utilities/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logger.Init()
	os.Exit(m.Run())
}

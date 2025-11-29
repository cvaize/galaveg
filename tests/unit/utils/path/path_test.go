package unit

import (
	"galaveg/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestCwd(t *testing.T) {
	assert.Contains(t, utils.Cwd(), "tests/unit/utils/path")
}

func TestFindModuleRoot(t *testing.T) {
	_, err := os.Stat(filepath.Join(utils.FindModuleRoot(utils.Cwd()), "go.mod"))
	assert.NoError(t, err)
}

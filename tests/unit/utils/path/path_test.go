package unit

import (
	"galaveg/utils/path"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestCwd(t *testing.T) {
	assert.Contains(t, path.Cwd(), "tests/unit/utils/path")
}

func TestFindModuleRoot(t *testing.T) {
	_, err := os.Stat(filepath.Join(path.FindModuleRoot(path.Cwd()), "go.mod"))
	assert.NoError(t, err)
}

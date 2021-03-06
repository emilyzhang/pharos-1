package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/lob/pharos/internal/test"
	"github.com/stretchr/testify/assert"
)

const (
	pharosConfig = "../testdata/pharosConfig"
	empty        = "../testdata/empty"
	nonexistent  = "../testdata/nonexistent"
)

func TestNew(t *testing.T) {
	t.Run("successfully creates reference to existing config file", func(tt *testing.T) {
		c, err := New(pharosConfig)
		assert.NoError(tt, err)
		assert.Equal(tt, pharosConfig, c.filePath)
	})

	t.Run("defaults to creating a config file reference at $HOME/.kube/pharos/config", func(tt *testing.T) {
		c, err := New("")
		assert.NoError(tt, err)
		assert.Equal(tt, fmt.Sprintf("%s/.kube/pharos/config", os.Getenv("HOME")), c.filePath)
	})
}

func TestLoad(t *testing.T) {
	t.Run("successfully loads existing config file", func(tt *testing.T) {
		// Create reference to config file.
		c, err := New(pharosConfig)
		assert.NoError(tt, err)

		// Loads file successfully into struct.
		err = c.Load()
		assert.NoError(tt, err)
		assert.Equal(tt, "http://localhost:7654", c.BaseURL)
	})

	t.Run("fails to load from nonexistent config", func(tt *testing.T) {
		c, err := New(nonexistent)
		assert.NoError(tt, err)
		assert.Equal(tt, nonexistent, c.filePath)

		err = c.Load()
		assert.Error(tt, err)
		assert.Contains(tt, err.Error(), "pharos hasn't been configured yet")
	})

	t.Run("fails to load from empty config", func(tt *testing.T) {
		c, err := New(empty)
		assert.NoError(tt, err)
		assert.Equal(tt, empty, c.filePath)

		err = c.Load()
		assert.Error(tt, err)
		assert.Contains(tt, err.Error(), "pharos hasn't been configured yet")
	})
}

func TestSave(t *testing.T) {
	t.Run("successfully saves to empty config", func(tt *testing.T) {
		// Create temporary empty test config file and defer cleanup.
		emptyConfigFile := test.CopyTestFile(tt, "../testdata", "config", empty)
		defer os.Remove(emptyConfigFile)

		// Create config reference to empty file.
		c, err := New(emptyConfigFile)
		assert.NoError(tt, err)
		assert.Equal(tt, emptyConfigFile, c.filePath)

		// Edit config and save it to file.
		c.AWSProfile = "egg"
		err = c.Save()
		assert.NoError(tt, err)

		// Check to see if file was saved successfully.
		c1, err := New(emptyConfigFile)
		assert.NoError(tt, err)
		err = c1.Load()
		assert.NoError(tt, err)
		assert.Equal(tt, "egg", c1.AWSProfile)
	})
}

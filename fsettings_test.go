package fsettings_test

import (
	"github.com/gtaylor/factorio-settings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorioSettings_String(t *testing.T) {
	fs := fsettings.NewFactorioServerSettings()
	output := fs.String()
	assert.NotEmpty(t, output)
}

func TestMapGenSettings_String(t *testing.T) {
	ms := fsettings.NewMapGenSettings()
	output := ms.String()
	assert.NotEmpty(t, output)
}

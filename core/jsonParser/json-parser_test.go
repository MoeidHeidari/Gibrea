package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeedFunction(t *testing.T) {
	err := ParseJson("./fixtures/repo-list-fixture.json")
	assert.NoError(t, err)
	assert.Equal(t, len(Dataset.Pages), 3)
	assert.Equal(t, len(Dataset.Pages[0].Repos), 10)
	assert.Equal(t, len(Dataset.Pages[1].Repos), 10)
	assert.Equal(t, len(Dataset.Pages[2].Repos), 10)
}

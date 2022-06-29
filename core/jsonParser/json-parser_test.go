package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeedFunction(t *testing.T) {
	err := Seed("./fixtures/repo-list-fixture.json")
	assert.NoError(t, err)
	assert.Equal(t, len(Seeder.Pages), 3)
	assert.Equal(t, len(Seeder.Pages[0].Repos), 10)
	assert.Equal(t, len(Seeder.Pages[1].Repos), 10)
	assert.Equal(t, len(Seeder.Pages[2].Repos), 10)
}

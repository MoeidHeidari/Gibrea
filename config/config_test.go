package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	InitFactory()
	config, error := Load("./fixtures/fixture-config.yaml")

	assert.NoError(t, error, "Parse example config failed.")

	assert.Equal(t, config.Config.AuthToken, ConfigFactory.auth_token)
	assert.Equal(t, config.Config.BaseAddress, ConfigFactory.base_url)
	assert.Equal(t, config.Config.Owner, ConfigFactory.owner)
	assert.Equal(t, len(config.Config.Tasks), 1)
	assert.Equal(t, config.Config.Tasks[0].Task.Name, ConfigFactory.Tasks[0].name)
	assert.Equal(t, config.Config.Tasks[0].Task.Service, ConfigFactory.Tasks[0].service)
	assert.Equal(t, config.Config.Tasks[0].Task.DataSet.MaxStars, ConfigFactory.Tasks[0].DataSet.MaxStars)
	assert.Equal(t, config.Config.Tasks[0].Task.DataSet.File, ConfigFactory.Tasks[0].DataSet.File)

}

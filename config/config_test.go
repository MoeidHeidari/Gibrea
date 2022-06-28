package config

import (
	"testing"

	"github.com/jinzhu/configor"
	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	InitFactory()
	result := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&ConfigStruct, "./fixtures/fixture-config.yaml")
	assert.NoError(t, result, "Parse example config failed.")

	assert.Equal(t, ConfigStruct.Config.AuthToken, ConfigFactory.auth_token)
	assert.Equal(t, ConfigStruct.Config.BaseAddress, ConfigFactory.base_url)
	assert.Equal(t, ConfigStruct.Config.Owner, ConfigFactory.owner)
	assert.Equal(t, len(ConfigStruct.Config.Tasks), 1)
	assert.Equal(t, ConfigStruct.Config.Tasks[0].Task.Name, ConfigFactory.Tasks[0].name)
	assert.Equal(t, ConfigStruct.Config.Tasks[0].Task.Service, ConfigFactory.Tasks[0].service)
	assert.Equal(t, ConfigStruct.Config.Tasks[0].Task.DataSet.MaxStars, ConfigFactory.Tasks[0].DataSet.MaxStars)
	assert.Equal(t, ConfigStruct.Config.Tasks[0].Task.DataSet.File, ConfigFactory.Tasks[0].DataSet.File)

}

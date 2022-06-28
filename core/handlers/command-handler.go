package core

import (
	"fmt"
	"main/config"

	"github.com/jinzhu/configor"
)

//########################################################################################################################
//handles start command with it's specified arguments
func StartHandler(configPath string, isService bool) (bool, error) {

	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&config.ConfigStruct, configPath)

	if err != nil {
		return false, err
	}
	fmt.Println("start handler", config.ConfigStruct.Config.BaseAddress, isService)
	return true, err

}

//========================================================================================================================
//handles stop command with it's specified arguments
func StopHandler() {
	fmt.Println("stop handler")
}

//========================================================================================================================
//handles pause command with it's specified arguments
func PauseHandler() {
	fmt.Println("pause handler")
}

//========================================================================================================================
//handles resum command with it's specified arguments
func ResumHandler() {
	fmt.Println("resum handler")
}

//========================================================================================================================
//handles status command with it's specified arguments
func StatusHandler() {
	fmt.Println("status handler")
}

//========================================================================================================================
//handles logs command with it's specified arguments
func LogsHandler(logFilePath string) {
	fmt.Println("logs handler", logFilePath)
}

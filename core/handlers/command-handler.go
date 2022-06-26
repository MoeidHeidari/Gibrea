package handlers

import (
	"fmt"
)

//handles start command with it's specified arguments
func StartHandler(config string, isService bool) {
	fmt.Println("start handler", config, isService)

}

//handles stop command with it's specified arguments
func StopHandler() {
	fmt.Println("stop handler")
}

//handles pause command with it's specified arguments
func PauseHandler() {
	fmt.Println("pause handler")
}

//handles resum command with it's specified arguments
func ResumHandler() {
	fmt.Println("resum handler")
}

//handles status command with it's specified arguments
func StatusHandler() {
	fmt.Println("status handler")
}

//handles logs command with it's specified arguments
func LogsHandler(logFilePath string) {
	fmt.Println("logs handler", logFilePath)
}

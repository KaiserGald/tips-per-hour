package error

import "log"

// HandleError handles basic errors
func HandleError(m string, e error) {
	if e != nil {
		log.Println(m, e)
	}
}

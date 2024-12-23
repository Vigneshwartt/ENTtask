package helpers

import (
	"fmt"
	"log"
	"os"
)

var (
	ErrorData *log.Logger
	WarnData  *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Error Occured", err)
		return
	}
	ErrorData = log.New(file, "ERROR: ", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	WarnData = log.New(file, "WARN:", log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
}

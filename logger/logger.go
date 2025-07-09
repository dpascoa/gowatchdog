package logger

import (
    "fmt"
    "os"
    "time"
)

var logFile *os.File
var Latest map[string]string = make(map[string]string)

func InitLog(filename string) error {
    var err error
    logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    return err
}

func CloseLog() {
    if logFile != nil {
        logFile.Close()
    }
}

func LogStatus(url string, statusCode int, alive bool) {
    status := "DOWN"
    if alive {
        status = "UP"
    }

    entry := fmt.Sprintf("[%s] %s - %d %s\n",
        time.Now().Format(time.RFC3339),
        url,
        statusCode,
        status,
    )

    fmt.Print(entry) // also show on terminal
    if logFile != nil {
        logFile.WriteString(entry)
    }
	
	Latest[url] = fmt.Sprintf("%d %s at %s", statusCode, status, time.Now().Format("15:04:05"))

}

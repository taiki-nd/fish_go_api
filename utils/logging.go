package utils

import (
	"io"
	"log"
	"os"
	"time"
)

func Logging(logFile string) {
	day := time.Now()
	const layout = "2006-01-02"

	logfile, err := os.OpenFile(logFile+"-"+day.Format(layout)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

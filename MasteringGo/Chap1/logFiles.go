package main

// On the macOS Catalina,
// $> cat /var/log/mail.log
// Nov  1 16:43:10 Raymonds-MacBook-Pro-2 Some program![26896]: 2019/11/01 16:43:10 LOG_MAIL: LOgging in Go!
// Nov  1 16:43:29 Raymonds-MacBook-Pro-2 Some program![26951]: 2019/11/01 16:43:29 LOG_MAIL: LOgging in Go!
//
import (
	"fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
	sysLog, err = syslog.New(syslog.LOG_MAIL, "Some program!")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}

	log.Println("LOG_MAIL: LOgging in Go!")
	fmt.Println("Will you see this?")

}

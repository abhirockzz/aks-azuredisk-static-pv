package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const fileLoc = "/mnt/logs/"
const fileName = "logz.out"

var logFile *os.File

func init() {
	var err error
	logFile, err = os.Create(fileLoc + fileName)
	if err != nil {
		panic(err)
	}
	log.Println("log file created")
}

func main() {
	ticker := time.NewTicker(3 * time.Second)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case t := <-ticker.C:
			logToFile(t.String())
		case <-exit:
			log.Println("exit signalled")
			err := os.Remove(fileLoc + fileName)
			if err != nil {
				log.Println("unable to delete log file")
			}
			log.Println("deleted log file")
			log.Println("exiting")
			os.Exit(1)
		}
	}
}

func logToFile(txt string) {
	_, err := logFile.WriteString(txt + "\n")
	if err != nil {
		log.Println("could not log " + txt + " due to err " + err.Error())
	}
}

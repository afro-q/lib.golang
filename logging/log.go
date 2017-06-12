package logging

import (
	"log"
	"os"

	"github.com/quinlanmorake/lib.golang/config"	
)

func SetupLog() {
	if config.SystemConfig.LogFile != "console" {
		f, fileErr := os.OpenFile(config.SystemConfig.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if fileErr != nil {
			log.Fatal("error opening file: %v", fileErr)
			return
		}
		
		defer f.Close()
		log.SetOutput(f)
	}	
}

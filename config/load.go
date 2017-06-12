package config

import (
	"io/ioutil"
	"path/filepath"
	"os"
)

// Function that will do the deserialisation of the config data
type DeserializeFunction func ([]byte) error

func Load(unMarshalFunction DeserializeFunction) error {
	appName := os.Args[0]
	
	var byConfigString []byte
		
	// Path config name is app name . config, it is expected that all config detail is contained herein		
	workingDir, _ := os.Getwd()
	configFilename := filepath.Join(workingDir, appName + ".config")
	if _, err := os.Stat(configFilename); err == nil {
		byConfigString, err = ioutil.ReadFile(configFilename)
		if err != nil {
			return err
		}
	}
	
	// If we have a string by now, load from that, otherwise, try AWS
	if len(byConfigString) > 0 {
		err := unMarshalFunction(byConfigString)
		if err != nil {
			return err
		}		
	}

	return nil
}

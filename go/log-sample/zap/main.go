package main

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
)

func main() {

	rawJSON := []byte(`{
	  "level": "debug",
      "development": true,
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		log.Fatal(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Info("info logger construction succeeded")
	logger.Debug("debug logger construction succeeded")
	logger.Error("error")
}

package main

import (
	"fmt"
	"os"

	"github.com/Volkov-D-A/gku-data-analisys/pkg/logger"
	"github.com/Volkov-D-A/gku-data-analisys/pkg/repos"
)

func main() {
	f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(nil)
	}

	log := logger.Get()
	log.SetOutput(f)

	log.Info("logger initialized successfully")

	data, err := repos.ImportDbfData("./testdata/chrg_144_92_202307.dbf", f)
	if err != nil {

	}

	fmt.Print(data)
}

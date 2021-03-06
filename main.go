package main

import (
	"github.com/sirupsen/logrus"
	"github.com/solvingj/credx/commands"
	"github.com/solvingj/credx/commands/common"
	"os"
)

// Override at command-line with:
//   go build -ldflags="-X main.Version=<SOMEVERSION>"
var Version = "dev"

func main() {
	var log = logrus.New()
	log.SetLevel(GetCliLogLevel())
	err := execMain()
	common.ExitOnErr(err)
}

func execMain() error {
	app := commands.GetApp(Version)
	args := os.Args
	err := app.Run(args)
	return err
}

func GetCliLogLevel() logrus.Level {
	switch os.Getenv("CREDX_LOG_LEVEL") {
	case "ERROR":
		return logrus.ErrorLevel
	case "WARN":
		return logrus.WarnLevel
	case "DEBUG":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
}
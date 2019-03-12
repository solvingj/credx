package commands

import (
	"github.com/codegangsta/cli"
	"github.com/solvingj/credx/commands/cred"
	docs "github.com/solvingj/credx/docs/credx"
)

func GetApp(version string) *cli.App {
	app := cli.NewApp()
	app.Name = "credx"
	app.Version = version
	app.Usage = docs.AppDescription
	app.Commands = GetCommands()
	app.HideHelp = true
	app.Flags = GlobalOptions()
	return app
}

func GlobalOptions() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "credx-home",
			Usage:  docs.HomeDirectoryUsage,
			EnvVar: "CREDX_HOME_DIR",
		},
		cli.StringFlag{
			Name:   "backend",
			Usage:  docs.BackendUsage,
			EnvVar: "CREDX_BACKEND",
		},
		cli.StringFlag{
			Name:   "log-level",
			Usage:  docs.LogLevelUsage,
			EnvVar: "CREDX_LOG_LEVEL",
		},
	}
}

func GetCommands() []cli.Command {
	return []cli.Command{
		cred.Command(),
	}
}



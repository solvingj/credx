package commands

import (
	"bytes"
	"flag"
	"github.com/codegangsta/cli"
	"strings"
)

func runAppSubCmd(command string) (string, error) {
	args := strings.Split(command, " ")
	app := GetApp("test")
	outBuffer := new(bytes.Buffer)
	app.Writer = outBuffer
	parFlagSet := flag.NewFlagSet("", 0)
	parContext := cli.NewContext(app, parFlagSet, nil)

	flagSet := flag.NewFlagSet("", 0)
	_ = flagSet.Parse(args)
	context := cli.NewContext(app, flagSet, parContext)
	err := app.RunAsSubcommand(context)

	if err != nil {
		return "", err
	}
	outBytes := outBuffer.Bytes()
	result := string(outBytes[:])
	return result, nil
}


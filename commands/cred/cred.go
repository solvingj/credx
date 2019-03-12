package cred

import (
	"fmt"
	"github.com/99designs/keyring"
	"github.com/codegangsta/cli"
	"github.com/solvingj/credx/commands/common"
	"github.com/solvingj/credx/cred"
	docs "github.com/solvingj/credx/docs/credx"
	creddocs "github.com/solvingj/credx/docs/credx/cred"
	"strings"
)

func Command() cli.Command {
	return cli.Command{
		Name:        "cred",
		Usage:       creddocs.Description,
		Subcommands: GetCommands(),
		HideHelp:    true,
		Flags:		 SharedCredFlags(),
	}
}

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:     "list",
			Usage:    creddocs.DescriptionList,
			HideHelp: true,
			Action: func(c *cli.Context) {
				ListCmd(c)
			},
		},
		{
			Name:      "set",
			Usage:     creddocs.DescriptionSet,
			ArgsUsage: "<variable=value>",
			HideHelp:  true,
			Action: func(c *cli.Context) {
				SetCmd(c)
			},
		},
		{
			Name:     "get",
			Usage:    creddocs.DescriptionGet,
			HideHelp: true,
			Flags:    GetCredFlags(),
			Action: func(c *cli.Context) {
				GetCmd(c)
			},
		},
		{
			Name:     "remove",
			Usage:    creddocs.DescriptionRemove,
			HideHelp: true,
			Action: func(c *cli.Context) {
				RemoveCmd(c)
			},
		},
	}
}

func ListCmd(c *cli.Context) {

	config := CreateConfig(c)
	kr, err := keyring.Open(config)

	if err != nil {
		common.ExitOnErr(err)
	}
	credNames, err := kr.Keys()
	if err != nil {
		common.ExitOnErr(err)
	}
	for _, k := range credNames {
		fmt.Println(k)
	}
}

func SetCmd(c *cli.Context) {
	config := CreateConfig(c)
	credentials := c.Args()
	if len(credentials) == 0 {
		common.PrintHelpAndExitWithError(`Not enough arguments.`, c)
	}
	for _, credential := range credentials {
		credSplit := strings.Split(credential, "=")
		if len(credSplit) < 2 {
			common.PrintHelpAndExitWithError(`credentials must be in form of "key=value" pairs.`, c)
		}
		credName := credSplit[0]
		credValue := strings.Join(credSplit[1:], "")
		err := cred.Set(config, credName, credValue)
		if err != nil {
			common.ExitOnErr(err)
		}
	}
}

func GetCmd(c *cli.Context) {
	config := CreateConfig(c)
	credNames := c.Args()
	if len(credNames) == 0 {
		common.PrintHelpAndExitWithError(`Not enough arguments.`, c)
	}
	for _, credName := range credNames {
		credItem, err := cred.Get(config, credName)
		credValue := string(credItem.Data)
		if err != nil {
			common.ExitOnErr(err)
		}
		showKeys := c.Bool("show-credNames")
		if showKeys {
			fmt.Println(credName + "=" + credValue)
		} else {
			fmt.Println(credValue)
		}
	}
}

func RemoveCmd(c *cli.Context) {
	config := CreateConfig(c)
	credNames := c.Args()
	if len(credNames) == 0 {
		common.PrintHelpAndExitWithError(`Not enough arguments.`, c)
	}
	for _, credName := range credNames {
		err := cred.Remove(config, credName)
		if err != nil {
			common.ExitOnErr(err)
		}
	}
}

func GetCredFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "show-keys",
			Usage: creddocs.ShowKeysFlagUsage,
		},
	}
}

func SharedCredFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "secret-service-name",
			Usage:  docs.SecretServiceNameUsage,
			EnvVar: "CREDX_SECRET_SERVICE_NAME",
		},
		cli.StringFlag{
			Name:   "creds-path",
			Usage:  docs.KeyPathUsage,
			EnvVar: "CREDX_CREDS_PATH",
		},
		cli.StringFlag{
			Name:   "keychain-name",
			Usage:  docs.KeychainNameUsage,
			EnvVar: "CREDX_KEYCHAIN_NAME",
		},
		cli.StringFlag{
			Name:   "libsecret-coll-name",
			Usage:  docs.LibsecretCollectionNameUsage,
			EnvVar: "CREDX_LIBSECRET_COLL_NAME",
		},
		cli.StringFlag{
			Name:   "kwallet-service-name",
			Usage:  docs.KWalletServiceNameUsage,
			EnvVar: "CREDX_KWALLET_SERVICE_NAME",
		},
		cli.StringFlag{
			Name:   "kwallet-app-id",
			Usage:  docs.KWalletAppIdUsage,
			EnvVar: "CREDX_KWALLET_APP_ID",
		},
		cli.StringFlag{
			Name:   "kwallet-folder",
			Usage:  docs.KWalletFolderUsage,
			EnvVar: "CREDX_KWALLET_FOLDER",
		},
		cli.StringFlag{
			Name:   "pass-cmd",
			Usage:  docs.PassCmdUsage,
			EnvVar: "CREDX_PASS_CMD",
		},
		cli.StringFlag{
			Name:   "pass-dir",
			Usage:  docs.PassDirUsage,
			EnvVar: "CREDX_PASS_DIR",
		},
		cli.StringFlag{
			Name:   "pass-prefix",
			Usage:  docs.PassPrefixUsage,
			EnvVar: "CREDX_PASS_PREFIX",
		},
	}
}

func CreateConfig(c *cli.Context) keyring.Config {
	backend := c.GlobalString("backend")
	if backend == "" {
		backend = cred.DefaultBackend()
	}
	if backend == "secret-service" {
		return SecretServiceKeyringConfig(c)
	} else if backend == "keychain" {
		return KeychainKeyringConfig(c)
	} else if backend == "kwallet" {
		return KWalletKeyringConfig(c)
	} else if backend == "wincred" {
		return WinCredKeyringConfig(c)
	} else if backend == "file" {
		return FileKeyringConfig(c)
	} else if backend == "pass" {
		return PassKeyringConfig(c)
	} else {
		common.PrintHelpAndExitWithError("Invalid backend specified: "+backend, c)
		return keyring.Config{}
	}
}

func SecretServiceKeyringConfig(c *cli.Context) keyring.Config {
	return cred.SecretServiceKeyringConfig(
		c.String("secret-service-name"),
		c.String("libsecret-coll-name"),
	)
}

func KeychainKeyringConfig(c *cli.Context) keyring.Config {
	return cred.KeychainKeyringConfig(
		c.String("keychain-name"),
	)
}

func KWalletKeyringConfig(c *cli.Context) keyring.Config {
	return cred.KWalletKeyringConfig(
		c.String("kwallet-service-name"),
		c.String("kwallet-app-id"),
		c.String("kwallet-folder"),
	)
}

func WinCredKeyringConfig(c *cli.Context) keyring.Config {
	return cred.WinCredKeyringConfig()
}

func FileKeyringConfig(c *cli.Context) keyring.Config {
	return cred.FileKeyringConfig(
		c.String("creds-path"),
	)
}

func PassKeyringConfig(c *cli.Context) keyring.Config {
	return cred.PassKeyringConfig(
		c.String("pass-dir"),
		c.String("pass-cmd"),
		c.String("pass-prefix"),
	)
}

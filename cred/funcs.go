package cred

import (
	"fmt"
	"github.com/99designs/keyring"
	"github.com/solvingj/credx/commands/common"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"runtime"
)

func Get(config keyring.Config, name string) (keyring.Item, error) {
	kr, err := keyring.Open(config)
	if err != nil {
		return keyring.Item{}, err
	}
	item, err := kr.Get(name)
	if err != nil {
		return keyring.Item{}, err
	}
	return item, nil
}

func Set(config keyring.Config, credName string, credValue string) error {
	kr, err := keyring.Open(config)
	item := keyring.Item{
		Key:  credName,
		Data: []byte(credValue),
	}
	err = kr.Set(item)
	return err
}

func Remove(config keyring.Config, credName string) error {
	kr, err := keyring.Open(config)
	if err != nil {
		return err
	}
	err = kr.Remove(credName)
	return err
}

func List(config keyring.Config) ([]string, error) {
	kr, err := keyring.Open(config)

	if err != nil {
		common.ExitOnErr(err)
	}
	keys, err := kr.Keys()
	return keys, err
}

func allowedBackends(backendName string) []keyring.BackendType {
	return []keyring.BackendType{
		keyring.BackendType(backendName),
	}
}

func SecretServiceKeyringConfig(serviceName string, collectionName string) keyring.Config {
	return keyring.Config{
		AllowedBackends:         allowedBackends("secret-service"),
		ServiceName:             serviceName,
		LibSecretCollectionName: collectionName,
	}
}

func KeychainKeyringConfig(keychainName string) keyring.Config {
	return keyring.Config{
		AllowedBackends:          allowedBackends("keychain"),
		KeychainName:             keychainName,
		KeychainTrustApplication: true,
	}
}

func KWalletKeyringConfig(serviceName string, appId string, folder string) keyring.Config {
	if appId == "" {
		appId = "credx"
	}
	if folder == "" {
		folder = "credx"
	}
	return keyring.Config{
		AllowedBackends: allowedBackends("kwallet"),
		ServiceName:     serviceName,
		KWalletAppID:    appId,
		KWalletFolder:   folder,
	}
}

func WinCredKeyringConfig() keyring.Config {
	return keyring.Config{
		AllowedBackends: allowedBackends("wincred"),
		ServiceName:     WinCredServiceName,
	}
}

func FileKeyringConfig(credsDir string) keyring.Config {
	if credsDir == "" {
		credsDir = CredxCredsDirDefault
	}
	return keyring.Config{
		AllowedBackends:  allowedBackends("file"),
		FileDir:          credsDir,
		FilePasswordFunc: FileKeyringPassphrasePrompt,
	}
}

func PassKeyringConfig(passDir string, passCmd string, passPrefix string) keyring.Config {
	return keyring.Config{
		AllowedBackends: allowedBackends("pass"),
		PassDir:         passDir,
		PassCmd:         passCmd,
		PassPrefix:      passPrefix,
	}
}

func DefaultBackend() string {
	switch osys := runtime.GOOS; osys {
	case "darwin":
		return "keyring"
	case "linux":
		return "secret-service"
	case "windows":
		return "wincred"
	default:
		return "file"
	}
}

func FileKeyringPassphrasePrompt(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	b, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
	return string(b), nil
}

package cred

import (
	"github.com/99designs/keyring"
	log "github.com/sirupsen/logrus"
	"github.com/solvingj/credx/commands/common"
	"testing"
)


func winCredConfig() keyring.Config {
	allowedBackends := []keyring.BackendType{
		keyring.BackendType("wincred"),
	}
	return keyring.Config{
		AllowedBackends:  allowedBackends,
		ServiceName:  "credx",
	}
}

func TestWinCredSetListGetRemoveKey(t *testing.T) {
	testCredName := "test-credName"
	preTestCredValue := "test-value"

	config := winCredConfig()
	_ = Set(config, testCredName, preTestCredValue)
	credNames, _ := List(config)
	for _, credName := range credNames {
		log.Println(credName)
	}
	postTestItem, _ := Get(config, testCredName)
	log.Println("postTestCredValue : " + string(postTestItem.Data))
	_ = Remove(config, testCredName)

}


func fileConfig() keyring.Config {
	allowedBackends := []keyring.BackendType{
		keyring.BackendType("file"),
	}
	return keyring.Config{
		AllowedBackends:  allowedBackends,
		FileDir:          "test_data/keys",
		FilePasswordFunc: func(string)(string, error) {
			return "testpassword", nil
		},
	}
}

func TestFileSetListGetRemoveKey(t *testing.T) {
	testCredName := "test-credName"
	preTestCredValue := "test-value"

	config := fileConfig()
	kr, err := keyring.Open(config)
	if err != nil {
		common.ExitOnErr(err)
	}
	preTestItem := keyring.Item{
		Key:  testCredName,
		Data: []byte(preTestCredValue),
	}

	err = kr.Set(preTestItem)

	if err != nil {
		common.ExitOnErr(err)
	}

	credNames, err := kr.Keys()

	if err != nil {
		common.ExitOnErr(err)
	}

	log.Print("Creds : " )

	for _, credName := range credNames {
		log.Println(credName)
	}

	postTestItem, err := kr.Get(testCredName)

	if err != nil {
		common.ExitOnErr(err)
	}

	log.Println("postTestCredValue : " + string(postTestItem.Data))

	err = kr.Remove(testCredName)

	if err != nil {
		common.ExitOnErr(err)
	}

}

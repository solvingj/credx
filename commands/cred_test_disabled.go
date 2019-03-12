package commands

import (
	"fmt"
	"testing"
)

func TestCommandCred(t *testing.T) {
	command := "credx cred"
	result, err := runAppSubCmd(command)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
	//%s `, command)
	//
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}

func TestCommandCredSetGetListRemove(t *testing.T) {
	testCredName := "test-name"
	testCredValue := "test-value"

	commandSet := fmt.Sprintf("credx cred set %s=%s", testCredName, testCredValue)
	result, err := runAppSubCmd(commandSet)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	commandGet := fmt.Sprintf("credx cred get %s", testCredName)
	result, err = runAppSubCmd(commandGet)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	commandList := fmt.Sprintf("credx cred list")
	result, err = runAppSubCmd(commandList)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	commandRemove := fmt.Sprintf("credx cred remove %s", testCredName)
	result, err = runAppSubCmd(commandRemove)
	if err != nil {
		t.Fatal()
	}

	t.Log(result)

	//expected := fmt.Sprintf(`USAGE:
	//%s `, command)
	//
	//if !strings.Contains(result, expected) {
	//	t.Errorf("result (%s) does not contain expected (%s)", result, expected)
	//}
}


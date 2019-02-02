package airq

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetKeyFile(t *testing.T) {

	err := GetKeyFile("/tmp/DoesntExit")
	if err == nil {
		t.Error("Expected Error")
		t.Fail()
	}

	file1 := []byte("")

	err = ioutil.WriteFile("/tmp/airq_short", file1, 0644)
	if err != nil {
		t.Error("Unexpected err while creating temp file:", err)
		t.Fail()
	}

	err = GetKeyFile("/tmp/airq_short")
	if err == nil {
		t.Error("Unexpected", err)
		t.Fail()
	}

	err = os.Remove("/tmp/airq_short")
	if err != nil {
		t.Error("Unexpected", err)
		t.Fail()
	}

	file2 := []byte("LongContent")
	err = ioutil.WriteFile("/tmp/airq_long", file2, 0644)
	if err != nil {
		t.Error("Unexpected", err)
		t.Fail()
	}

	err = GetKeyFile("/tmp/airq_long")
	if err != nil {
		t.Error("Unexpected", err)
		t.Fail()
	}

	err = os.Remove("/tmp/airq_long")
	if err != nil {
		t.Error("Unexpected", err)
		t.Fail()
	}

}

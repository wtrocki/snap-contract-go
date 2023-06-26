package examples

import (
	"testing"

	snapcontract "github.com/wtrocki/snap-contract-go/v1"
)

type TestStruct struct {
	Name string
}

func TestString(t *testing.T) {
	result := TestStruct{Name: "Hello world"}
	err := snapcontract.Snapshot(result)
	if err != nil {
		t.Errorf("This will pass because \"Hello world\" is in the snapshot %s", err)
	}
	result.Name = "Hello world!"
	err = snapcontract.Snapshot(result)
	if err == nil {
		t.Errorf("Now it will fail because the snapshot doesn't have an exclamation mark %s", err)
	}
}

 
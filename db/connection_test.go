package db

import (
	"fmt"
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	_, err := connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[TestConnect]: %v", err)
		t.Fail()
	}
}

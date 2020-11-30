package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	res := print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}

}

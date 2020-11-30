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

func TestPrint1(t *testing.T) {

	fmt.Println(t.Name())
	fmt.Println("test demo")
}

func testName(t *testing.T) {
	fmt.Println("test name")
}

//在其他函数中调用test
func TestAll(t *testing.T) {
	t.Run("testPrint", TestPrint)
	t.Run("TestPrint1", TestPrint1)
	t.Run("testName", testName)
}

// 自动调用的test main
func TestMain(m *testing.M) {
	m.Run()
}

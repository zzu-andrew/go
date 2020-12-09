package main

import "fmt"

type interAnimal interface {
	speakSt
	speakData
}

type speakSt interface {
	s()
}

type speakData interface {
	getData()
}

type interfaceCat struct {
	data int
	name string
}

func (x *interfaceCat) s() {

}

func (x *interfaceCat) getData() {
	fmt.Println("demo")
}

func main() {

	var tmp *interfaceCat
	tmp = new(interfaceCat)

	tmp.getData()
}

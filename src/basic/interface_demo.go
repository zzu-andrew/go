package main

import "fmt"

func main() {
	/*
				接口：interface
				在Go中，接口是一组方法的签名
				当某个类型为这个几口中的所有方法提供了方法的实现，它被称为实现接口
				Go语言中，接口和类型的实现关系，是非非侵入式
				如java中需要指明实现的接口 :class Mouse implements USB{}
			1. 当需要接口类型的对象时， 可以使用任意实现类对象代替
			2. 接口对象不能实现，实现类中的字段
		接口使用说明
			1. 如果一个函数接受接口类型作为参数， 那么实际上可以传入该接口的任意实现类型的对象作为参数
			2. 定义一个接口类型，实际上可以赋值为任意实现类的对象
		鸭子类型：

	*/
	// 1. 逻辑鼠标
	m1 := Mouse{name: "逻辑鼠标"}
	fmt.Println(m1.name)
	// 创建flash
	f1 := FlashDisk{name: "闪迪"}
	fmt.Println(f1.name)

	testInterface(m1)
	testInterface(f1)
	f1.delete()

	fmt.Println("interface demo")
}

// 1. 定义接口
// 实现接口的可以有很多
type USB interface {
	start() // USB设备开始工作
	end()
}

// 实现类
type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标，准备就绪，可以开始工作了， 点点点")
}

func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，可以安全退出...")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "准备开始工作， 可以进行数据的存储..")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出")
}

func (f FlashDisk) delete() {
	fmt.Println(f.name, "U盘删除数据")
}

// 3. 测试方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}

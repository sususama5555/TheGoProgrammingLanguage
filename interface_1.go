package main

import "fmt"

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "mouse is ready to work")
}

func (m Mouse) end() {
	fmt.Println(m.name, "is stop work")
}

func testInterface(usb USB) {
	usb.start()
	usb.end()
}

func main() {
	m := Mouse{"logic G304"}
	fmt.Println(m.name)

	var usb USB

	usb = m

	// 接口对象，不能访问实现类的属性
	// fmt.Println(usb.name)

	usb.start()
	usb.end()
	testInterface(m)

}

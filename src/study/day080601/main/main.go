package main

import "fmt"

//声明接口
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

type Computer struct {
}

func (p Phone) Start() {
	fmt.Println("phone start.......")
}

func (p Phone) Stop() {
	fmt.Println("phone stop.......")
}

func (c Camera) Start() {
	fmt.Println("camera start.....")
}
func (c Camera) Stop() {
	fmt.Println("camera stop......")
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	computer.Working(camera)

}

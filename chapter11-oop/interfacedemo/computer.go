package interfacedemo

type Usb interface {
	Start()
	Stop()
}

type Computer struct {
}

// 该方法接收一个Usb接口类型变量
func (c Computer) Working(usb Usb) {

	// 通过usb接口变量调用Start和Stop方法
	usb.Start()

	// 如果usb是指向Phone结构体变量，则还需要调用Call()方法
	// 这里用到了 类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}

	usb.Stop()

}

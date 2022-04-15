package function

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaphone NokiaPhone) call() {
	fmt.Println("nokia is calling")
}

type Iphone struct {
}

func Interface() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

}

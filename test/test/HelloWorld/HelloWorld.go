package HelloWorld

import "fmt"

type HelloWorld struct {
}

func NewHelloWorld() (rcvr *HelloWorld) {
	rcvr = &HelloWorld{}
	return
}
func HelloWorld_main() {
	fmt.Println("Hello World")
}

type inner struct {
}

func NewInner() (rcvr *inner) {
	rcvr = &inner{}
	return
}
func (rcvr *inner) Test() (int) {
	fmt.Println("test1")
	return 1
}
func (rcvr *inner) TestStringint(s string, a int) (string) {
	fmt.Println("test4")
	return "returntest4"
}
func (rcvr *inner) Testint(a int) {
	fmt.Println("test2")
}
func (rcvr *inner) TestintString(a int, s string) (string) {
	fmt.Println("test3")
	return "returntest3"
}
func inner_main() {
	o := NewInner()
	fmt.Println(o.Test())
	o.Testint(1)
	fmt.Println(o.TestintString(1, "test3"))
	fmt.Println(o.TestStringint("test4", 1))
}

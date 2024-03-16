package Inter

import "java2go/test/utils"

type A interface {
	Eat()
	Sleep()
}
type B interface {
	Show()
}
type Inter struct {
}

func (rcvr *Inter) Eat() {
}
func NewInter() (rcvr *Inter) {
	rcvr = &Inter{}
	return
}
func (rcvr *Inter) Show() {
}
func (rcvr *Inter) Sleep() {
}

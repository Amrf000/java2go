package Comments

import "java2go/test/utils"

type Comments struct {
	X int
}

func NewComments() (rcvr *Comments) {
	rcvr = &Comments{}
	rcvr.X = 10
	y := 10
	z := 10
	return
}
func NewCommentsComments(that *Comments) (rcvr *Comments) {
	rcvr = &Comments{}
	rcvr.X = 2
	return
}
func NewCommentsString(name string) (rcvr *Comments) {
	rcvr = &Comments{}
	rcvr.X = 2
	return
}
func (rcvr *Comments) Test() (string) {
	y := 10
	z := 10
	rcvr.X = 3
	rcvr.X++
	w := utils.Plus(&rcvr.X)
	z := utils.Ternary(w == 2, 1, 5)
	return ""
}
func (rcvr *Comments) TestArrint(a []int) (string) {
}
func Comments_main() {
	o := NewComments()
	b := NewCommentsString("test")
	c := NewCommentsComments(o)
	o.Test()
	b.Test()
	aa := []int{0}
	o.TestArrint(aa)
}

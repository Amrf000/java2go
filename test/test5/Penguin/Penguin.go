package Mouse

import "fmt"

type Animal struct {
	name string
	id   int
}

func (rcvr *Animal) Eat() {
}
func (rcvr *Animal) Move() {
	fmt.Println("动物可以移动")
}
func NewAnimal() (rcvr *Animal) {
	rcvr = &Animal{}
	return
}
func NewAnimalStringint(myName string, myid int) (rcvr *Animal) {
	rcvr = &Animal{}
	return
}
func (rcvr *Animal) Sleep() {
}

type Dog struct {
	*Animal
}

func NewDog() (rcvr *Dog) {
	rcvr = &Dog{}
	return
}
func (rcvr *Dog) Move() {
	rcvr.Animal.Move()
	fmt.Println("狗可以跑和走")
}

type Mouse struct {
	*Animal
}

func NewMouseStringint(myName string, myid int) (rcvr *Mouse) {
	rcvr = &Mouse{}
	rcvr.Animal = NewAnimalStringint(myName, myid)
	return
}

type Penguin struct {
	*Animal
}

func NewPenguinStringint(myName string, myid int) (rcvr *Penguin) {
	rcvr = &Penguin{}
	rcvr.Animal = NewAnimalStringint(myName, myid)
	return
}
func Penguin_main() {
	a := NewAnimal()
	b := NewDog()
	a.Move()
	b.Move()
}

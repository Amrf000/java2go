package JavaLoop

import "fmt"
import "java2go/test/utils"

type JavaLoop struct {
}

func NewJavaLoop() (rcvr *JavaLoop) {
	rcvr = &JavaLoop{}
	return
}
func JavaLoop_main() {
	x := 10
	for x < 20 {
		fmt.Print(fmt.Sprintf("%v%v", "value of x : ", x))
		x++
		fmt.Print("\n")
	}
	x = 10
	for {
		fmt.Print(fmt.Sprintf("%v%v", "value of x : ", x))
		x++
		fmt.Print("\n")
		if !(x < 20) {
			break
		}
	}
	for x := 10; x < 20; x = x + 1 {
		fmt.Print(fmt.Sprintf("%v%v", "value of x : ", x))
		fmt.Print("\n")
	}
	numbers := []int{10, 20, 30, 40, 50}
	for _, x := range numbers {
		fmt.Print(x)
		fmt.Print(",")
	}
	fmt.Print("\n")
	names := []string{"James", "Larry", "Tom", "Lacy"}
	for _, name := range names {
		fmt.Print(name)
		fmt.Print(",")
	}
	for _, x := range numbers {
		if x == 30 {
			break
		}
		fmt.Print(x)
		fmt.Print("\n")
	}
	for _, x := range numbers {
		if x == 30 {
			continue
		}
		fmt.Print(x)
		fmt.Print("\n")
	}
}

package Test

import "fmt"
import "java2go/test/utils"

type Test struct {
}

func NewTest() (rcvr *Test) {
	rcvr = &Test{}
	return
}
func Test_main() {
	grade := 'C'
	switch grade {
	case 'A':
		fmt.Println("优秀")
	case 'B','C':
		fmt.Println("良好")
	case 'D':
		fmt.Println("及格")
	case 'F':
		fmt.Println("你需要再努力努力")
	default:
		fmt.Println("未知等级")
	}
	fmt.Println(fmt.Sprintf("%v%v", "你的等级是 ", grade))
}

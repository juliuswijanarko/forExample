//Package scnd waw
package scnd

import "fmt"

//CalcErr is user defined type in package scnd
type CalcErr struct {
	len   int
	heigh int
	err   error
}

//Error is awesome in package scnd
func (n CalcErr) Error() string {
	return fmt.Sprintf("Minimum value is: %d, %d, %v", n.len, n.heigh, n.err)
}

func main() {
	v, err := calc(0, 1)
	if err != nil {
		fmt.Println(v)
		fmt.Println(err)
	} else {
		fmt.Println(v, err)
	}
}

func calc(len, heigh int) (int, error) {
	if len < 1 || heigh < 1 {
		msg := fmt.Errorf("Input number must higher than: %d and %d", len, heigh)
		return 0, CalcErr{1, 1, msg}
	}
	return len * heigh, nil

}

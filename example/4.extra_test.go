package main

import (
	"log"
	"testing"

	"fmt"

	"github.com/yezihack/e"
)

func div(x, y int) (int, error) {
	if y == 0 {
		e1 := e.ErrorF("x:%d, y:%d", x, y)
		err := e.New("除数不能等于0", e1)
		return 0, err
	}
	return x / y, nil
}
func TestDiv(t *testing.T) {
	a, err := div(1, 0)
	if e.Assert(err) {
		// 输出 extra 里的错误
		fmt.Printf("extra-error:%v\n", e.Convert(err).Errs())
		// 输出堆栈信息
		log.Fatalln(e.Convert(err).ToStr())
	}
	fmt.Println("a", a)
}

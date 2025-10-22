package goroutines

import (
	"fmt"
	"strconv"
	"time"
)

func showMessage(msg string) {
	fmt.Printf("%v \n", msg)
}

func RunSample01() {

	for i := 0; i < 10; i++ {
		go showMessage(strconv.Itoa(i))
	}

	time.Sleep(1 * time.Second)
}

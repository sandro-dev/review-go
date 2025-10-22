package goroutines

import (
	"fmt"
	"sync"
)

func RunSample03() {
	var i int
	var wg sync.WaitGroup
	var m sync.Mutex

	nGoroutines := 100_000
	wg.Add(nGoroutines)

	for x := 0; x < nGoroutines; x++ {
		go func() {
			defer wg.Done()

			m.Lock()
			i++
			m.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("O valor final de i é:", i)
}

package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func callDatabase(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Iniciando chamada a banco de dados...")
	time.Sleep(1 * time.Second)
	fmt.Println("FINALIZADO: callDatabase")
}

func callAPI(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Iniciando chamada a API...")
	time.Sleep(1 * time.Second)
	fmt.Println("FINALIZADO: callAPI")
}

func processInterno(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Iniciando processamento interno...")
	time.Sleep(1 * time.Second)
	fmt.Println("FINALIZADO: processInterno")
}

func Sample01() {

	fmt.Println("Iniciando processos concorrentes...")

	var wg sync.WaitGroup
	wg.Add(3)

	go callDatabase(&wg)
	go callAPI(&wg)
	go processInterno(&wg)

	wg.Wait()

	fmt.Println("Todos os processos finalizados. Encerrando o programa.")

	// for i := 0; i < 10; i++ {
	// 	go showMessage(strconv.Itoa(i))
	// }

	// time.Sleep(time.Second * 1)
}

// func showMessage(msg string) {
// 	fmt.Println(msg)
// }

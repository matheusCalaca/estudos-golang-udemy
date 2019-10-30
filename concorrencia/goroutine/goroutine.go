package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for index := 0; index < qtde; index++ {
		time.Sleep(time.Second)
		fmt.Printf("%s : '%s' %d \n", pessoa, texto, index)
	}
}

func main() {
	// fale("matheus", "falou", 4)
	// fale("calaça", "teste 2", 2)

	// com o GO antes da função ela vira uma go rotine mais nescessita do chane para se utilizar, mais afrente tera o exemplo com chanel

	// go fale("matheus", "falou", 4)
	// go fale("calaça", "teste 2", 2)

}

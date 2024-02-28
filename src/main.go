package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoVictor := ContaCorrente{titular: "Victor", numeroAgencia: 555, numeroConta: 9911, saldo: 250.50}

	contadaBruna := ContaCorrente{"Bruna", 555, 9191, 125}

	fmt.Println(contaDoVictor)
	fmt.Println(contadaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.numeroAgencia = 575
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)
}

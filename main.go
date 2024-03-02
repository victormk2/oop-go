package main

import( 
	"fmt"
	c "oop-go/conta"
)

func main() {
	contaDoVictor := c.ContaCorrente{Titular: "Victor", NumeroAgencia: 555, NumeroConta: 9911, Saldo: 500}
	fmt.Println(contaDoVictor)
	fmt.Println(contaDoVictor.Saldo)

	fmt.Println(contaDoVictor.Depositar(300))
	status, valor := contaDoVictor.Depositar(300)
	fmt.Println(status, valor)
	fmt.Println(contaDoVictor.Saldo)
}

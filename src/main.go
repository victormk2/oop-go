package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso, saldo atual:", valorDeposito
	}

	return "Valor inválido, saldo atual:", valorDeposito
}

func main() {
	contaDoVictor := ContaCorrente{titular: "Victor", numeroAgencia: 555, numeroConta: 9911, saldo: 500}
	fmt.Println(contaDoVictor)
	fmt.Println(contaDoVictor.saldo)

	fmt.Println(contaDoVictor.depositar(300))
	status, valor := contaDoVictor.depositar(300)
	fmt.Println(status, valor)
	fmt.Println(contaDoVictor.saldo)
}

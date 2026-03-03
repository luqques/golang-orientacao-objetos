package contas

import (
	"fmt"

	"github.com/luques/golang-orientacao-objetos/clientes"
)

type ContaCorrente struct {
	Saldo       float64
	NumeroConta int
	Titular     clientes.Titular
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	contaDestino.Saldo += valorTransferencia
	c.Saldo -= valorTransferencia

	fmt.Printf("Conta destino %v recebeu %v reais, totalizando %v de saldo.\n", contaDestino.NumeroConta, valorTransferencia, contaDestino.Saldo)
	fmt.Printf("Conta origem %v está com o saldo total de %v.\n", c.NumeroConta, c.Saldo)
	return true
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	c.Saldo += valorDeposito
	fmt.Printf("Conta %v recebeu um depósito de %v reais, totalizando %v de saldo.\n", c.NumeroConta, valorDeposito, c.Saldo)
}

func (c *ContaCorrente) Sacar(valorSaque float64) bool {
	if valorSaque > c.Saldo {
		fmt.Printf("Saldo insuficiente para saque de %v reais. Saldo atual: %v reais.\n", valorSaque, c.Saldo)
		return false
	}
	c.Saldo -= valorSaque
	fmt.Printf("Conta %v realizou um saque de %v reais, saldo restante: %v reais.\n", c.NumeroConta, valorSaque, c.Saldo)
	return true
}

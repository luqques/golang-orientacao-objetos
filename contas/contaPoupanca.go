package contas

import (
	"fmt"

	"github.com/luques/golang-orientacao-objetos/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroConta, NumeroAgencia, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {
	contaDestino.saldo += valorTransferencia
	c.saldo -= valorTransferencia

	fmt.Printf("Conta destino %v recebeu %v reais, totalizando %v de saldo.\n", contaDestino.NumeroConta, valorTransferencia, contaDestino.saldo)
	fmt.Printf("Conta origem %v está com o saldo total de %v.\n", c.NumeroConta, c.saldo)
	return true
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) {
	c.saldo += valorDeposito
	fmt.Printf("Conta %v recebeu um depósito de %v reais, totalizando %v de saldo.\n", c.NumeroConta, valorDeposito, c.saldo)
}

func (c *ContaPoupanca) Sacar(valorSaque float64) bool {
	if valorSaque > c.saldo {
		fmt.Printf("saldo insuficiente para saque de %v reais. saldo atual: %v reais.\n", valorSaque, c.saldo)
		return false
	}
	c.saldo -= valorSaque
	fmt.Printf("Conta %v realizou um saque de %v reais, saldo restante: %v reais.\n", c.NumeroConta, valorSaque, c.saldo)
	return true
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

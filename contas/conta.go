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

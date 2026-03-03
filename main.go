package main

import "fmt"

type Conta struct {
	Saldo       float64
	NumeroConta int
}

func (c *Conta) Transferir(valorTransferencia float64, contaDestino *Conta) bool {
	contaDestino.Saldo += valorTransferencia
	c.Saldo -= valorTransferencia

	fmt.Printf("Conta destino %v recebeu %v reais, totalizando %v de saldo.\n", contaDestino.NumeroConta, valorTransferencia, contaDestino.Saldo)
	fmt.Printf("Conta origem %v está com o saldo total de %v.\n", c.NumeroConta, c.Saldo)
	return true
}

func main() {
	contaDestino := Conta{Saldo: 10, NumeroConta: 12345}
	contaOrigem := Conta{Saldo: 20, NumeroConta: 54321}

	contaOrigem.Transferir(5, &contaDestino)
}

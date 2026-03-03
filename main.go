package main

import "github.com/luques/golang-orientacao-objetos/contas"

func main() {
	contaDestino := contas.ContaCorrente{Saldo: 10, NumeroConta: 12345}
	contaOrigem := contas.ContaCorrente{Saldo: 20, NumeroConta: 54321}

	contaOrigem.Transferir(5, &contaDestino)
}

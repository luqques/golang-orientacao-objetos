package main

import (
	"fmt"

	"github.com/luques/golang-orientacao-objetos/clientes"
	"github.com/luques/golang-orientacao-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valorSaque float64) bool
}

func main() {
	lucas := clientes.Titular{
		Nome:      "Lucas",
		Cpf:       "123.456.789-00",
		Profissao: "Desenvolvedor",
	}

	contaLucas := contas.ContaCorrente{
		NumeroConta: 123456,
		Titular:     lucas,
	}

	contaLucas.Depositar(100)
	contaLucas.Sacar(30)
	PagarBoleto(&contaLucas, 10)

	fmt.Print(contaLucas.ObterSaldo())
}

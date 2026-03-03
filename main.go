package main

import (
	"github.com/luques/golang-orientacao-objetos/clientes"
	"github.com/luques/golang-orientacao-objetos/contas"
)

func main() {
	lucas := clientes.Titular{
		Nome:      "Lucas",
		Cpf:       "123.456.789-00",
		Profissao: "Desenvolvedor",
	}

	contaLucas := contas.ContaCorrente{
		Saldo:       1000,
		NumeroConta: 123456,
		Titular:     lucas,
	}

	print(contaLucas)
}

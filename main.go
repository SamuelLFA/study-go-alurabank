package main

import (
	"fmt"

	"github.com/samuellfa/study-go-alurabank/clientes"
	"github.com/samuellfa/study-go-alurabank/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) bool
}

func main() {
	titular := clientes.Titular{Nome: "Samuel", CFF: "12715497601", Profissao: "Desenvolvedor"}
	conta := contas.ContaPoupanca{Titular: titular, NumeroAgencia: 123, Operacao: 13, NumeroConta: 112233}
	conta2 := contas.ContaCorrente{Titular: titular, NumeroAgencia: 123, NumeroConta: 112233}

	conta.Depositar(200)

	PagarBoleto(&conta, 60)

	fmt.Println(conta.ObterSaldo())
	fmt.Println(conta2.ObterSaldo())
}

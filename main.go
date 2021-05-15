package main

import (
	"fmt"

	"github.com/samuellfa/study-go-alurabank/contas"
)

func main() {
	conta := contas.ContaCorrente{Titular: "Samuel", NumeroAgencia: 123, NumeroConta: 112233, Saldo: 200.0}
	conta2 := contas.ContaCorrente{Titular: "Douglas", NumeroAgencia: 123, NumeroConta: 112233, Saldo: 300.0}

	conta.Transferir(100, &conta2)

	fmt.Println(conta)
	fmt.Println(conta2)
}

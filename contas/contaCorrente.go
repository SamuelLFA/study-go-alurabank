package contas

import "github.com/samuellfa/study-go-alurabank/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return true
	}

	return false
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) bool {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return true
	}

	return false
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)

		return true
	}

	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

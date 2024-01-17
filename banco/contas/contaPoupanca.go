package contas;

import(
	"fmt";
	"banco/clientes";
)


type ContaPoupanca struct{

	Titular clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64

};


func (c* ContaPoupanca) Sacar(valorSaque float64) string{

	saquePossivel := valorSaque > 0 && c.saldo >= valorSaque;

	if(saquePossivel){
		c.saldo -= valorSaque;
		return "Saque efetuado";
	}else {
		return "Não foi possível efetuar o Saque";
	}
}

func (c* ContaPoupanca) Depositar(valorDeposito float64) (string, float64){

	depositoPossivel := valorDeposito > 0;

	if(depositoPossivel){
		c.saldo += valorDeposito;
		return "Depósito feito com sucesso", c.saldo;
	} else{
		return "Não foi possível efetuar o Saque", c.saldo;
	}

}


func (c* ContaPoupanca) ExibirSaldo(){

	fmt.Println("Saldo da conta corrente do ", c.Titular.Nome, ": ", c.saldo);

}


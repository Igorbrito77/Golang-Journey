package contas;

import(
	"strconv";
	"fmt";
	"banco/clientes";
)


type ContaCorrente struct{

	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	saldo float64

};


func (c* ContaCorrente) Sacar(valorSaque float64) string{

	saquePossivel := valorSaque > 0 && c.saldo >= valorSaque;

	if(saquePossivel){
		c.saldo -= valorSaque;
		return "Saque efetuado";
	}else {
		return "Não foi possível efetuar o Saque";
	}
}

func (c* ContaCorrente) Depositar(valorDeposito float64) (string, float64){

	depositoPossivel := valorDeposito > 0;

	if(depositoPossivel){
		c.saldo += valorDeposito;
		return "Depósito feito com sucesso", c.saldo;
	} else{
		return "Não foi possível efetuar o Saque", c.saldo;
	}

}


func (c* ContaCorrente ) Transferir(valorTransferencia float64, contaDestino* ContaCorrente) (string, float64){

	transferenciaPossivel := valorTransferencia > 0 && valorTransferencia <= c.saldo;

	if(transferenciaPossivel){
		contaDestino.saldo += valorTransferencia;
		c.saldo -= valorTransferencia;
		return "Transferência no valor de  "+ strconv.FormatFloat(valorTransferencia, 'f', -1, 64)+ " de " + c.Titular.Nome + " para o "+ contaDestino.Titular.Nome+  " feita com sucesso", c.saldo;
	} else{
		return "Não foi possível efetuar a transição", c.saldo;
	}

}

func (c* ContaCorrente) ExibirSaldo(){

	fmt.Println("Saldo da conta corrente do ", c.Titular.Nome, ": ", c.saldo);

}


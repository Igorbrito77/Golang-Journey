package contas;

import(
	"strconv";
	"banco/clientes";
)


type ContaCorrente struct{

	Titular clientes.Titular
	NumeroAgencia int
	NumeroConta int
	Saldo float64

};


func (c* ContaCorrente) sacar(valorSaque float64) string{

	saquePossivel := valorSaque > 0 && c.Saldo >= valorSaque;

	if(saquePossivel){
		c.Saldo -= valorSaque;
		return "Saque efetuado";
	}else {
		return "Não foi possível efetuar o Saque";
	}
}

func (c* ContaCorrente) depositar(valorDeposito float64) (string, float64){

	depositoPossivel := valorDeposito > 0;

	if(depositoPossivel){
		c.Saldo += valorDeposito;
		return "Depósito feito com sucesso", c.Saldo;
	} else{
		return "Não foi possível efetuar o Saque", c.Saldo;
	}

}


func (c* ContaCorrente ) Transferir(valorTransferencia float64, contaDestino* ContaCorrente) (string, float64){

	transferenciaPossivel := valorTransferencia > 0 && valorTransferencia <= c.Saldo;

	if(transferenciaPossivel){
		contaDestino.Saldo += valorTransferencia;
		c.Saldo -= valorTransferencia;
		return "Transferência no valor de  "+ strconv.FormatFloat(valorTransferencia, 'f', -1, 64)+ " de " + c.Titular + " para o "+ contaDestino.Titular+  " feita com sucesso", c.Saldo;
	} else{
		return "Não foi possível efetuar a transição", c.Saldo;
	}

}


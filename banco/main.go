package main;

import "fmt";

type ContaCorrente struct{

	titular string
	numeroAgencia int
	numeroConta int
	saldo float64

};

func (c* ContaCorrente) sacar(saque float64) string{

	saquePossivel := saque > 0 && c.saldo >= saque;

	if(saquePossivel){
		c.saldo -= saque;
		return "Saque efetuado";
	}else {
		return "Não foi possível efetuar o saque";
	}
}

func (c* ContaCorrente) depositar(valorDeposito float64) (string, float64){

	depositoPossivel := valorDeposito > 0;

	if(depositoPossivel){
		c.saldo += valorDeposito;
		return "Depósito feito com sucesso", c.saldo;
	} else{
		return "Não foi possível efetuar o saque", c.saldo;
	}

}

func main(){

	var contaJorel *ContaCorrente;

	contaJorel = new(ContaCorrente);

	contaJorel.titular = "Jorel";
	contaJorel.saldo = 1000;

	fmt.Println(*contaJorel);

//	fmt.Println(contaJorel.sacar(-14));

	fmt.Println(contaJorel.depositar(-400));


	fmt.Println(*contaJorel);

}


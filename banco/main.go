package main;

import (
	"fmt";
	"banco/contas";
);




func main(){

	/*var contaJorel *contas.ContaCorrente;

	contaJorel = new(contas.ContaCorrente);

	contaJorel.Titular = "Jorel";
	contaJorel.Saldo = 1000;*/

	var contaLurdes = contas.ContaCorrente{Titular: "Lurdes", Saldo: 3500};



//	fmt.Println(*contaJorel);
	fmt.Println(contaLurdes);


	fmt.Println(contaLurdes.sacar(-14));

	//fmt.Println(contaJorel.depositar(-400));

	//fmt.Println(contaLurdes.Transferir(500.2412154545, contaJorel));


	//fmt.Println(*contaJorel);
	fmt.Println(contaLurdes);


}


package main;

import (
	"fmt";
	"banco/contas";
	"banco/clientes";
);




func main(){

	/*var contaJorel *contas.ContaCorrente;

	contaJorel = new(contas.ContaCorrente);

	contaJorel.Titular = "Jorel";
	contaJorel.Saldo = 1000;*/

	contaLurdes := contas.ContaCorrente{Titular: clientes.Titular{ 
														Nome: "Lurdes",
														Cpf: "123.456.789-12",
														Profissao: "Bancária" }, Saldo: 3500};



//	fmt.Println(*contaJorel);
	fmt.Println(contaLurdes);


	fmt.Println(contaLurdes.Sacar(200));

	//fmt.Println(contaJorel.depositar(-400));

	//fmt.Println(contaLurdes.Transferir(500.2412154545, contaJorel));


	//fmt.Println(*contaJorel);
	fmt.Println(contaLurdes);


}


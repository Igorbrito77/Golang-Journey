package main;

import (
	"fmt";
	"banco/contas";
	//"banco/clientes";
);


func pagarBoleto(conta verificarConta, valorBoleto float64){

	conta.Sacar(valorBoleto);
}

type verificarConta interface {

	Sacar(valor float64) string;
}


func main(){

/*
	var titularJorel *clientes.Titular;
	titularJorel = new(clientes.Titular);

	titularJorel.Nome = "Jorel";
	titularJorel.Cpf = "147.258.369-14";
	titularJorel.Profissao = "Dev";

	var contaJorel *contas.ContaCorrente;
	contaJorel = new(contas.ContaCorrente);
	contaJorel.Titular = *titularJorel;

	contaJorel.Depositar(1000);
	

	contaLurdes := contas.ContaCorrente{Titular: clientes.Titular{ 
														Nome: "Lurdes",
														Cpf: "123.456.789-12",
														Profissao: "Bancária" }};


	contaLurdes.Depositar(3500);

	fmt.Println(*contaJorel);
	fmt.Println(contaLurdes);


	fmt.Println(contaLurdes.Sacar(200));

	//fmt.Println(contaJorel.depositar(-400));

	//fmt.Println(contaLurdes.Transferir(500.2412154545, contaJorel));


	//fmt.Println(*contaJorel);
	fmt.Println(contaLurdes);



	contaExemplo := contas.ContaCorrente{};
	contaExemplo.Depositar(241);

	contaExemplo.ExibirSaldo()*/
	
	


	minhaCorrente := contas.ContaCorrente{};

	minhaCorrente.Depositar(1200.25);

	minhaPoupanca := contas.ContaPoupanca{};

	minhaPoupanca.Depositar(640.50);

	//fmt.Println("minhaCorrente : ", minhaCorrente)

	fmt.Println("minhaPoupanca : ", minhaPoupanca)

	pagarBoleto(&minhaPoupanca, 200);

	fmt.Println("minhaPoupanca : ", minhaPoupanca)


	fmt.Println("minhaCorrente : ", minhaCorrente)

	pagarBoleto(&minhaCorrente, 300.13);

	fmt.Println("minhaCorrente : ", minhaCorrente)


}


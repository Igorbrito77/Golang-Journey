package main;

import "fmt";

type ContaCorrente struct{

	titular string
	numeroAgencia int
	numeroConta int
	saldo float64

};


func main(){

	contaIgor := ContaCorrente{titular: "Igor", numeroAgencia: 1234, numeroConta: 7896, saldo: 40000.74} 

	contaMaria := ContaCorrente{titular: "Maria", numeroAgencia: 4567, numeroConta: 7412, saldo: 345.41} 

	var contaJorel *ContaCorrente;

	contaJorel = new(ContaCorrente);

	contaJorel.titular = "Jorel";
	contaJorel.saldo = 5000;
	
	fmt.Println(contaIgor);

	fmt.Println(contaMaria);

	fmt.Println(*contaJorel);


}


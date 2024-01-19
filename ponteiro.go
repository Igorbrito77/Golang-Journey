package main

import(
	"fmt";
	"time";
	"math/rand";
)

func sortear_ganho( saldo *float32){

	fmt.Println("  INICIANDO SORTEIO ");

  	rand.Seed(time.Now().UnixNano());

	fmt.Println("na função = ", *saldo, " ", &saldo);


	*saldo = rand.Float32()

	fmt.Println("Você ganhou: R$", *saldo );

	fmt.Println("na função = ", *saldo, " ", &saldo);


}


func main(){
	
	var saldo float32 = 0.;


	fmt.Println("Saldo inicial: ", saldo, " Endereço: ", &saldo);

	sortear_ganho(&saldo)

	fmt.Println("Você tem agora: R$", saldo, " Endereço: ", &saldo);



}
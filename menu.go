package main

import "fmt"
import "os"

func main(){

	exibir_introducao();

	comando := inserir_comando();


	switch comando {

		case 1: 
			fmt.Println("Monitorando ...")
		case 2:
			fmt.Println("Exibindo logs ...")
		case 0:
			fmt.Println("Saindo ...")
			os.Exit(0);
		default	:
			fmt.Println("Comando inválido ...")
			os.Exit(-1);
	}
/*
	if comando == 1 {
		fmt.Println("Monitorando ...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs ...")
	} else {
		fmt.Println("Saindo ...")
	}*/

}


func exibir_introducao(){

	nome_usuario := "Igor";
	var versao_software float32 = 1.00;
	
	fmt.Println("Saudações, ", nome_usuario, "! Tenha um excelente dia, meu nobre.")
	fmt.Println("A versão deste software: ", versao_software)

}

func inserir_comando() int{

	var comando int;

	fmt.Println("1- Inicar monitoramento");
	fmt.Println("2- Exibir logs");
	fmt.Println("0- Sair do programa");

	fmt.Scan(&comando);

	// fmt.Scanf("%d", &comando)

	// fmt.Println("O endereço da minha variável 'comando' é: ", &comando);
	//fmt.Println("Opção selecionada: ", comando);

	return comando;

}
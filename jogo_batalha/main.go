package main;

import (
	"fmt";
	"jogo_batalha/personagem";
	"jogo_batalha/barbaro";
);

type Arqueiro struct{

	Atk int;
	Defesa int;
	Numero_flechas int;

}

func (arqueiro Arqueiro) Atacar(){

	fmt.Println("Arqueiro atirou uma flecha. Dano: ", arqueiro.Atk);
}


func (arqueiro Arqueiro) AumentarAtk(incremento_atk int){

	arqueiro.Atk += incremento_atk;
}

type Jogador struct{

	nome string;
	personagem personagem.Personagem;
}

func main(){


	fmt.Println("Escolha uma classe de personagem: 1- Bárbaro, 2- Arqueiro");

	var escolha int; 

	fmt.Scan(&escolha);

	var jogador  personagem.Personagem;

	if(escolha == 1){
		//jogador = new( personagem.Personagem); 
		//(*jogador.Atk) = 95; // Defesa: 45, Numero_machados: 3};

		jogador = &barbaro.Barbaro{Atk: 95, Defesa: 45, Numero_machados: 3};

	}else{
		jogador = Arqueiro{Atk: 70, Defesa: 35, Numero_flechas: 40};
	}


	fmt.Println(jogador);
    fmt.Printf("jogador: %T\n", jogador)


	jogador.Atacar();

	jogar(&jogador);

}


func jogar(jogador * personagem.Personagem){



	var comando int;

	for{

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** MENU PRINCIPAL **** \n")
		fmt.Println("\n\n Você está andando pela floresta e escuta um som atrás de uma árvore. O que você faz?");
		fmt.Println("1- Conferir o que há atrás da árvore");
		fmt.Println("2- Ignorar e seguir o seu caminho");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);


		if(comando == 1){

			(*jogador).AumentarAtk(10);
		}




	}


}
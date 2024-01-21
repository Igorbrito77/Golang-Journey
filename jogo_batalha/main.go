package main;

import (
	"fmt";
	"jogo_batalha/personagem";
	"jogo_batalha/barbaro";
);

type Arqueiro struct{

	atk int;
	defesa int;
	numero_flechas int;

}

func (arqueiro Arqueiro) Atacar(){

	fmt.Println("Arqueiro atirou uma flecha. Dano: ", arqueiro.atk);
}

type Jogador struct{

	nome string;
	personagem personagem.Personagem;
}

func main(){


	fmt.Println("Escolha uma classe de personagem: 1- Bárbaro, 2- Arqueiro");

	var escolha int; 

	fmt.Scan(&escolha);

	var jogador personagem.Personagem;

	if(escolha == 1){
		jogador = barbaro.Barbaro{Atk: 95, Defesa: 45, Numero_machados: 3};

	}else{
		jogador = Arqueiro{atk: 70, defesa: 35, numero_flechas: 40};
	}


	fmt.Println(jogador);

	jogador.Atacar();

}

func jogar(){

}
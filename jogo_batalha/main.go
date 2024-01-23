package main;

import (
	"fmt";
	"jogo_batalha/personagem";
	"jogo_batalha/jogador";
	"jogo_batalha/inimigo";
	//"jogo_batalha/barbaro";
	//"jogo_batalha/arqueiro";
);


// jogo de fases. com ataques aleatpórios. toda vez que passar de uma fase, criar uma fase nova trazendo um novo tipo de inimigo, e somando mais atk ao jogagor e ao inimigo seguinte;

func main(){

	fmt.Println("Escolha uma classe de personagem: 1- Bárbaro, 2- Arqueiro");

	var escolha int; 

	fmt.Scan(&escolha);

	//player := jogador.Jogador{};

	var player * jogador.Jogador = new(jogador.Jogador);


	
	if(escolha == 1){
		player.Inicializar(personagem.Barbaro);

	}else{
		player.Inicializar(personagem.Arqueiro);
	}


	//fmt.Println(player);
  //  fmt.Printf("Tipo do jogador: %T\n", player)
 
	jogar(player);

}

// a interface já funciona como um ponteiro ???


func faseAtacar(p personagem.Personagem){

	p.Atacar();

}

func jogar(player * jogador.Jogador){

	var comando int;

	inimigo := inimigo.Inimigo{};

	inimigo.Inicializar(personagem.Arqueiro);

	for{

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** MENU PRINCIPAL **** \n")
		fmt.Println("\n\n Você está andando pela floresta e escuta um som atrás de uma árvore. O que você faz?");
		fmt.Println("1- Conferir o que há atrás da árvore");
		fmt.Println("2- Ignorar e seguir o seu caminho");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);

		switch(comando){

			case 1: 
				player.AumentarAtk(10);
			
			default:
				return;

		}


		faseAtacar(player);
		faseAtacar(&inimigo);

		player.AddItens("Bomba");
		player.ExibirStatus();

	}


}
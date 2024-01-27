package main;

import (
	"fmt";
	"os";
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

	var player * jogador.Jogador = new(jogador.Jogador);

	
	if(escolha == 1){
		player.Inicializar(personagem.Barbaro);

	}else{
		player.Inicializar(personagem.Arqueiro);
	}

  //  fmt.Printf("Tipo do jogador: %T\n", player)
 
	jogar(player);

}

func jogar(player * jogador.Jogador){

	resultado := fase1(player);

	if(resultado == -1){
		os.Exit(0);
	}

	resultado = fase2(player);

	

}

func fase1(player * jogador.Jogador) int{

	var comando int;

	inimigo := inimigo.Inimigo{};

	inimigo.Inicializar(personagem.Arqueiro);

	fmt.Println("\n\n Um inimigo ", inimigo.TipoPersonagem, " apareceu. Lute ou morra!");

	inimigo.ExibirStatus();


	for inimigo.Vida >= 0 && player.Vida >= 0 {

		// loop de batalha

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** FASE 1! LUTE OU MORRA! **** \n")
		fmt.Println("1- Atacar");
		fmt.Println("2- Procurar item especial");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);

		switch(comando){

			case 1: 
				player.Atacar(&inimigo);
			case 2: 
				player.AddItens("Bomba");
				//player.AumentarAtk(10);
			
			default:
				return -1;

		}

		inimigo.Atacar(player);
		
		inimigo.ExibirStatus();
		player.ExibirStatus();

	}

	if(inimigo.Vida <= 0){
		fmt.Println("Vitória! Você derrotou o inimigo!");

		player.AumentarAtk(25);

		return 1;
	}else{
		fmt.Println("Fim de Jogo! Você foi derrotado!");
		return -1;
	}
}

func fase2(player * jogador.Jogador) int{

	var comando int;

	inimigo := inimigo.Inimigo{};

	inimigo.Inicializar(personagem.Barbaro);

	fmt.Println("\n\n Um inimigo ", inimigo.TipoPersonagem, " apareceu. Lute ou morra!");

	inimigo.ExibirStatus();


	for inimigo.Vida >= 0 && player.Vida >= 0 {

		// loop de batalha

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** FASE 2 ! LUTE OU MORRA! **** \n")
		fmt.Println("1- Atacar");
		fmt.Println("2- Procurar item especial");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);

		switch(comando){

			case 1: 
				player.Atacar(&inimigo);
			case 2: 
				player.AddItens("Bomba");
				//player.AumentarAtk(10);
			
			default:
				return -1;

		}

		inimigo.Atacar(player);
		
		inimigo.ExibirStatus();
		player.ExibirStatus();

	}

	if(inimigo.Vida <= 0){
		fmt.Println("Vitória! Você derrotou o inimigo!");

		player.AumentarAtk(25);

		return 1;
	}else{
		fmt.Println("Fim de Jogo! Você foi derrotado!");
		return -1;
	}
}
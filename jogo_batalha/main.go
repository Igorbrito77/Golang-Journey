package main;

import (
	"fmt";
	"os";
	"jogo_batalha/personagem";
	"jogo_batalha/jogador";
	"jogo_batalha/inimigo";
);


type Fase struct{

	inimigo inimigo.Inimigo;
	numero int;

}

func (fase * Fase ) iniciarFase( player* jogador.Jogador) int{

	var comando int;


	fmt.Println("\n\n Um inimigo ", fase.inimigo.TipoPersonagem, " apareceu. Lute ou morra!");

	fase.inimigo.ExibirStatus();
	player.ExibirStatus();

	for fase.inimigo.Vida >= 0 && player.Vida >= 0 {

		// loop de batalha

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** FASE ", fase.numero, "! LUTE OU MORRA! **** \n")
		fmt.Println("1- Atacar");
		fmt.Println("2- Procurar item especial");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);

		switch(comando){

			case 1: 
				player.Atacar(&fase.inimigo);
			case 2: 
				player.AddItens("Bomba");
				//player.AumentarAtk(10);
			
			default:
				return -1;

		}

		fase.inimigo.Atacar(player);
		
		fase.inimigo.ExibirStatus();
		player.ExibirStatus();

	}

	if(fase.inimigo.Vida <= 0){
		fmt.Println("Vitória! Você derrotou o inimigo!");

		player.AumentarAtk(25);

		return 1;
	}else{
		fmt.Println("Fim de Jogo! Você foi derrotado!");
		return -1;
	}
	

}

type Livro struct{

	fases [] Fase;

}

func (livro * Livro ) iniciarHistoria(player* jogador.Jogador){


	resultado:= livro.fases[0].iniciarFase(player);

	if(resultado == -1){
		return;
	}

	resultado= livro.fases[1].iniciarFase(player);


}
 

// jogo de fases. com ataques aleatpórios. toda vez que passar de uma fase, criar uma fase nova trazendo um novo tipo de inimigo, e somando mais atk ao jogagor e ao inimigo seguinte;

func main(){


	inimigoFase1 := inimigo.Inimigo{};

	inimigoFase1.Inicializar(personagem.Arqueiro);

	fase1:= Fase{ inimigo : inimigoFase1, numero: 1}

	//fase1:= Fase{ inimigos : [] inimigo.Inimigo {inimigoFase1}  }

	fmt.Println(fase1);

	livro := &Livro{  fases: [] Fase{fase1} };
	
	
	fmt.Println(livro);

	

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

  	livro.iniciarHistoria(player);
 
	//jogar(player);

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
	player.ExibirStatus();

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
	player.ExibirStatus();

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
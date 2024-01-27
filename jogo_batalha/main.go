package main;

import (
	"fmt";
	//"os";
	"jogo_batalha/personagem";
	"jogo_batalha/jogador";
	"jogo_batalha/inimigo";
	"jogo_batalha/livro";
	"jogo_batalha/fase";
);




// jogo de fases. com ataques aleatpórios. toda vez que passar de uma fase, criar uma fase nova trazendo um novo tipo de inimigo, e somando mais atk ao jogagor e ao inimigo seguinte;


func criarLivro() livro.Livro{

	inimigoFase1 := inimigo.Inimigo{};

	inimigoFase1.Inicializar(personagem.Arqueiro);

	fase1:= fase.Fase{ Inimigo : inimigoFase1, Numero: 1};


	inimigoFase2 := inimigo.Inimigo{};

	inimigoFase2.Inicializar(personagem.Barbaro);

	fase2:= fase.Fase{ Inimigo : inimigoFase2, Numero: 2};


	livro := livro.Livro{  Fases: [] fase.Fase{fase1, fase2} };
	
	return livro;

}

func main(){

	livro := criarLivro();

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

  	livro.IniciarHistoria(player);
 
}
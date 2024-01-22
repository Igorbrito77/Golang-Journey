package jogador;

import(
	"fmt"
)

type Jogador struct{

	Atk int;
	Defesa int;

	TipoPersonagem string;

}
/*
type TipoPersonagem struct{

	

}*/

func (jogador * Jogador) Atacar(){

	fmt.Println("Bárbaro atacou com o machado. Dano: ", jogador.Atk);
}

func (jogador * Jogador) AumentarAtk(incremento_atk int){

	jogador.Atk += incremento_atk;

	fmt.Println("atk agora == ", jogador.Atk );
}

func (jogador * Jogador) Inicializar(tipo_personagem string){

	jogador.TipoPersonagem = tipo_personagem;

	if(tipo_personagem == "Bárbaro"){
		jogador.Atk = 95;
		jogador.Defesa = 50;
	}else{
		jogador.Atk = 70;
		jogador.Defesa = 35;
	}


	fmt.Println("Jogador ", jogador.TipoPersonagem , " criado !");
}

func (jogador * Jogador) ExibirStatus(){

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println(" 								**** STATUS DO JOGADOR **** \n")
	fmt.Println("Clase: ", jogador.TipoPersonagem);
	fmt.Println("Ataque: ", jogador.Atk);
	fmt.Println("Defesa: ", jogador.Defesa);
}

func (jogador * Jogador) AddItens(){


}
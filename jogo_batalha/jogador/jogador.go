package jogador;

import(
	"fmt";
	"jogo_batalha/personagem";
)

type Jogador struct{

	Atk int;
	Defesa int;

	TipoPersonagem string;

	Itens [] string;
	
}


func (jogador * Jogador) Atacar(){

	fmt.Println("-> Você (", jogador.TipoPersonagem, ") atacou . Dano: ", jogador.Atk);
}

func (jogador * Jogador) AumentarAtk(incremento_atk int){

	jogador.Atk += incremento_atk;

	fmt.Println("Seu ataque subiu ", incremento_atk, " pontos . Seu ataque agora é de ", jogador.Atk, " pontos");
}

func (jogador * Jogador) Inicializar(tipo_personagem string){

	jogador.TipoPersonagem = tipo_personagem;

	if(tipo_personagem == personagem.Barbaro){
		jogador.Atk = 95;
		jogador.Defesa = 50;
	}else{
		jogador.Atk = 70;
		jogador.Defesa = 35;
	}

}

func (jogador * Jogador) ExibirStatus(){

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println(" 								**** STATUS DO JOGADOR **** \n")
	fmt.Println("Clase: ", jogador.TipoPersonagem);
	fmt.Println("Ataque: ", jogador.Atk);
	fmt.Println("Defesa: ", jogador.Defesa);
	fmt.Println("Itens: ", jogador.Itens);
}

func (jogador * Jogador) AddItens(novo_item string){

	jogador.Itens = append(jogador.Itens, novo_item);

}
package inimigo;

import(
	"fmt";
	"jogo_batalha/personagem";
)

type Inimigo struct{

	Atk int;
	Defesa int;
	Vida int;
	TipoPersonagem string;

}

func (inimigo * Inimigo) Atacar(jogador personagem.Personagem){

	danoCausado := jogador.LevarDano(inimigo.Atk);

	fmt.Println("-> O inimigo", inimigo.TipoPersonagem, " te atacou. Dano sofrido: ", danoCausado);
}

func (inimigo * Inimigo) AumentarAtk(incremento_atk int){

	inimigo.Atk += incremento_atk;

	fmt.Println("atk agora == ", inimigo.Atk );
}

func (inimigo * Inimigo) Inicializar(tipo_personagem string){

	inimigo.TipoPersonagem = tipo_personagem;

	if(tipo_personagem == "Bárbaro"){
		inimigo.Atk = 95;
		inimigo.Defesa = 50;
		inimigo.Vida = 100;
	}else{
		inimigo.Atk = 70;
		inimigo.Defesa = 35;
		inimigo.Vida = 100;
	}

}

func (inimigo * Inimigo) ExibirStatus(){

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println(" 								**** STATUS DO INIMIGO **** \n")
	fmt.Println("Clase: ", inimigo.TipoPersonagem);

	var barraVida = "";

	for i := 0; i < inimigo.Vida; i++{
		barraVida += "|";
	} 

	fmt.Println("Vida: ", barraVida, " (", inimigo.Vida, ")");
	fmt.Println("Ataque: ", inimigo.Atk);
	fmt.Println("Defesa: ", inimigo.Defesa);

}

func (inimigo* Inimigo) LevarDano(atkPersonagem int) int{

	dano := atkPersonagem - inimigo.Defesa;

	inimigo.Vida -= dano;

	return dano;

}
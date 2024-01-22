package inimigo;

import(
	"fmt"
)

type Inimigo struct{

	Atk int;
	Defesa int;

	TipoPersonagem string;

}

func (inimigo * Inimigo) Atacar(){

	fmt.Println("Bárbaro atacou com o machado. Dano: ", inimigo.Atk);
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
	}else{
		inimigo.Atk = 70;
		inimigo.Defesa = 35;
	}


	fmt.Println("inimigo ", inimigo.TipoPersonagem , " criado !");
}

func (inimigo * Inimigo) ExibirStatus(){

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println(" 								**** STATUS DO INIMIGO **** \n")
	fmt.Println("Clase: ", inimigo.TipoPersonagem);
	fmt.Println("Ataque: ", inimigo.Atk);
	fmt.Println("Defesa: ", inimigo.Defesa);
}


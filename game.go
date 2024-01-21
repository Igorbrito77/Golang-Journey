package main;

import (
	"fmt"
);

type Barbaro struct{

	atk int;
	defesa int;
	numero_machados int;

}


type Arqueiro struct{

	atk int;
	defesa int;
	numero_flechas int;

}

type Personagem interface{

	atacar();
	//caminhar();

};

func (barbaro Barbaro) atacar(){

	fmt.Println("Bárbaro atacou com o machado. Dano: ", barbaro.atk);
}

func (arqueiro Arqueiro) atacar(){

	fmt.Println("Arqueiro atirou uma flecha. Dano: ", arqueiro.atk);
}

func destruir(p Personagem){

	p.atacar();
}


/*
func (Personagem p) atacar2(){



}*/

type Jogador struct{

	nome string;
	personagem Personagem;
}


func main(){

	barbaro := Barbaro{atk: 95, defesa: 45, numero_machados: 3};

	arqueiro := Arqueiro{atk: 70, defesa: 35, numero_flechas: 40};


	igor := Jogador{ nome: "Igor", personagem: barbaro};

	bob := Jogador{nome: "Bob", personagem: arqueiro}


	fmt.Println(igor)

	fmt.Println(bob)


	destruir(barbaro);

	destruir(arqueiro);

}
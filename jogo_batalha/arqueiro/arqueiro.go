package arqueiro;

import(
	"fmt"
)

type Arqueiro struct{

	Atk int;
	Defesa int;
	Numero_flechas int;

}

func (arqueiro *Arqueiro) Atacar(){

	fmt.Println("Arqueiro atirou uma flecha. Dano: ", arqueiro.Atk);
}


func (arqueiro *Arqueiro) AumentarAtk(incremento_atk int){

	arqueiro.Atk += incremento_atk;
}
package personagem;

type Personagem interface{
	Atacar();
	//caminhar();

	AumentarAtk(int);

};


func Destruir(p Personagem){

	p.Atacar();
}

func exibirStatus(){



}

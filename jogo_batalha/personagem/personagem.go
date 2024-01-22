package personagem;

type Personagem interface{
	Atacar();
	//caminhar();

	AumentarAtk(int);

	Inicializar(string);

	ExibirStatus()

};


func Destruir(p Personagem){

	p.Atacar();
}

func exibirStatus(){



}

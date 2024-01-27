package personagem;

type Personagem interface{
	Atacar( Personagem);
	//caminhar();
	LevarDano(int) int;

	AumentarAtk(int);

	Inicializar(string);

	ExibirStatus()

};


func Destruir(p Personagem){

	//p.Atacar();
}

func exibirStatus(){



}

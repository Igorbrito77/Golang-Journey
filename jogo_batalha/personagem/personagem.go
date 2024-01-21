package personagem;

type Personagem interface{

	Atacar();
	//caminhar();

};


func Destruir(p Personagem){

	p.Atacar();
}
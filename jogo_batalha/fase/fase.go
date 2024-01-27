package fase;

import(
	"fmt";
	//"jogo_batalha/personagem";
	"jogo_batalha/jogador";
	"jogo_batalha/inimigo";

)


type Fase struct{

	Inimigo inimigo.Inimigo;
	Numero int;

}


func ImprimirEspaco(){
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}

func ExibirStatusDuelo(player * jogador.Jogador, inimigo inimigo.Inimigo){

	inimigo.ExibirStatus();
	player.ExibirStatus();
}

func (fase * Fase ) IniciarFase( player* jogador.Jogador) int{

	var comando int;

	ImprimirEspaco();

	fmt.Println("-----------------------------------------------------------------------------------------------------\n")
	fmt.Println(" ------------------------------------------ FASE ", fase.Numero, " ---------------------------------\n")
	fmt.Println("-----------------------------------------------------------------------------------------------------\n")

	fmt.Println("\n\n Um inimigo ", fase.Inimigo.TipoPersonagem, " apareceu. Lute ou morra!");


	for fase.Inimigo.Vida >= 0 && player.Vida >= 0 {
		

		ExibirStatusDuelo(player, fase.Inimigo);

		// loop de batalha

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** FASE ", fase.Numero, "! LUTE OU MORRA! **** \n")
		fmt.Println(" 					1- Atacar | 2- Procurar item especial | 0- Sair do jogo  \n");

		fmt.Scan(&comando);

		switch(comando){

			case 1: 
				player.Atacar(&fase.Inimigo);
			case 2: 
				player.AddItens("Bomba");
				//player.AumentarAtk(10);
			
			default:
				return -1;

		}
	
		if(fase.Inimigo.Vida <= 0){
			break;
		}

		fase.Inimigo.Atacar(player);

		/*if(player.Vida <= 0){
			break;
		}*/

		fmt.Println("Digite algo para continuar a batalha\n");
		fmt.Scan(&comando);

		ImprimirEspaco();

	}


	resultado := 1;

	//ImprimirEspaco();

	ExibirStatusDuelo(player, fase.Inimigo);


	if(fase.Inimigo.Vida <= 0){
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n               Vitória! Você derrotou o inimigo!								");

		player.AumentarAtk(25);

	}else{
		fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n								Fim de Jogo! Você foi derrotado!");
		resultado = -1;
	}



	fmt.Println("Digite algo para seguir para a próxima fase\n");
	fmt.Scan(&comando);

	return resultado;
	
}

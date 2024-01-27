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


func (fase * Fase ) IniciarFase( player* jogador.Jogador) int{

	var comando int;


	fmt.Println("\n\n Um inimigo ", fase.Inimigo.TipoPersonagem, " apareceu. Lute ou morra!");

	fase.Inimigo.ExibirStatus();
	player.ExibirStatus();

	for fase.Inimigo.Vida >= 0 && player.Vida >= 0 {

		// loop de batalha

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** FASE ", fase.Numero, "! LUTE OU MORRA! **** \n")
		fmt.Println("1- Atacar");
		fmt.Println("2- Procurar item especial");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

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

		fase.Inimigo.Atacar(player);
		
		fase.Inimigo.ExibirStatus();
		player.ExibirStatus();

	}

	if(fase.Inimigo.Vida <= 0){
		fmt.Println("Vitória! Você derrotou o inimigo!");

		player.AumentarAtk(25);

		return 1;
	}else{
		fmt.Println("Fim de Jogo! Você foi derrotado!");
		return -1;
	}
	

}

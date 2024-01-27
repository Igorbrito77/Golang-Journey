package fase;

import(
	"fmt";
	"jogo_batalha/personagem";
	"jogo_batalha/jogador";
	"jogo_batalha/inimigo";

)


type Fase struct{

	inimigo inimigo.Inimigo;
	numero int;

}

func (fase * Fase ) iniciarFase( player* jogador.Jogador) int{

	var comando int;


	fmt.Println("\n\n Um inimigo ", fase.inimigo.TipoPersonagem, " apareceu. Lute ou morra!");

	fase.inimigo.ExibirStatus();
	player.ExibirStatus();

	for fase.inimigo.Vida >= 0 && player.Vida >= 0 {

		// loop de batalha

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** FASE ", fase.numero, "! LUTE OU MORRA! **** \n")
		fmt.Println("1- Atacar");
		fmt.Println("2- Procurar item especial");
		fmt.Println("0- Sair do jogo");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);

		switch(comando){

			case 1: 
				player.Atacar(&fase.inimigo);
			case 2: 
				player.AddItens("Bomba");
				//player.AumentarAtk(10);
			
			default:
				return -1;

		}

		fase.inimigo.Atacar(player);
		
		fase.inimigo.ExibirStatus();
		player.ExibirStatus();

	}

	if(fase.inimigo.Vida <= 0){
		fmt.Println("Vitória! Você derrotou o inimigo!");

		player.AumentarAtk(25);

		return 1;
	}else{
		fmt.Println("Fim de Jogo! Você foi derrotado!");
		return -1;
	}
	

}

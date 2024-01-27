package fase;

import(
	"fmt";
	"jogo_batalha/jogador";
)



type Livro struct{

	fases [] Fase;

}

func (livro * Livro ) iniciarHistoria(player* jogador.Jogador){


	resultado:= livro.fases[0].iniciarFase(player);

	if(resultado == -1){
		return;
	}

	resultado= livro.fases[1].iniciarFase(player);


}
 
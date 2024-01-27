package livro;

import(
	"jogo_batalha/jogador";
	"jogo_batalha/fase";
)



type Livro struct{

	Fases [] fase.Fase;

}

func (livro * Livro ) IniciarHistoria(player* jogador.Jogador){


	resultado:= livro.Fases[0].IniciarFase(player);

	if(resultado == -1){
		return;
	}

	resultado= livro.Fases[1].IniciarFase(player);

}
 
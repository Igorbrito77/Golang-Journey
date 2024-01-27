package livro;

import(
	"fmt",
	"jogo_batalha/jogador";
	"jogo_batalha/fase";
)



type Livro struct{

	Fases [] fase.Fase;

}

func (livro * Livro ) IniciarHistoria(player* jogador.Jogador){

	var resultado int;

	for _, fase:= range livro.Fases {

		resultado = fase.IniciarFase(player);

		if(resultado == -1){
			return;
		}

	}

	fmt.Println("                          PARABÉNS: JOGO CONCLUÍDO !!	");

}
 
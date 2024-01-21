package barbaro;

import(
	"fmt"
)

type Barbaro struct{

	Atk int;
	Defesa int;
	Numero_machados int;

}

func (barbaro Barbaro) Atacar(){

	fmt.Println("Bárbaro atacou com o machado. Dano: ", barbaro.Atk);
}

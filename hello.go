package main

import "fmt"
import "reflect"

func main() {

	var nome string = "Igor Brito";
	apelido:= "Igord";
	var idade int = 25;
	var taxa_acerto float32 = 91.2;

	fmt.Println("Nome do usuário logado: ", nome, "\nApelido: ", apelido, " \nIdade: ", idade, "\nTaxa de acerto: ", taxa_acerto)

	fmt.Println("Tipo da variável 'apelido:' ", reflect.TypeOf(apelido));

}
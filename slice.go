package main;

import(
	"fmt"
)

func main() {

	// ARRAY

	fmt.Println("          ***** ARRAY **** ");

	//var array_nomes [2] string;
	array_nomes := [2] string{"Igor", "Maria"};

	fmt.Println("Nomes: ", array_nomes);
	fmt.Println("Tamanho: ", len(array_nomes), " Capacidade: ", cap(array_nomes));

	// SLICE

	fmt.Println("\n\n          ***** SLICE **** ");


	nomes := []	string{"Messi", "Cr7", "Neymar"};

	fmt.Println("Nomes: ", nomes);
	fmt.Println("Tamanho: ", len(nomes), " Capacidade: ", cap(nomes));


	nomes = append(nomes, "Pirlo"); // À cada insert, o go dobra a capacidade do slice

	fmt.Println("Nomes: ", nomes);
	fmt.Println("Tamanho: ", len(nomes), " Capacidade: ", cap(nomes));

}

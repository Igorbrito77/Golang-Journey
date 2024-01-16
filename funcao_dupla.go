package main
import "fmt"


func main(){

	nome, versao := retornarDados();

	fmt.Println("Usuário : ", nome, ". Versão: ", versao);

	_, versao2 := retornarDados(); // _ usado pra ignorar uma variável 

	fmt.Println("Versão: ", versao2);

	nome2, _ := retornarDados();

	fmt.Println("Usuário : ", nome2);


}


func retornarDados() (string, float32) {

	return "I.A", 1.289

}
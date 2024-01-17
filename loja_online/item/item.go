package item;

import(
	"fmt";
	"loja_online/usuario";
)

type Item struct{

	Nome string;
	Preco float64;
	Codigo int;

};


func (item* Item) ExibirCodigo(usuario usuario.Usuario){

	if(usuario.Perfil == "admin"){
		fmt.Println("Código do produto ", item.Nome, ": ", item.Codigo)
	}else{
		fmt.Println("Você não tem permissão para visualizar códigos de itens!");
	}	
}
package item;

import(
	"fmt";
	"loja_online/usuario";
)

type Item struct{

	Nome string;
	Preco float64;
	codigo int;

};


func (item* Item) ExibirCodigo(usuario usuario.Usuario){

	if(usuario.Perfil == "admin"){
		fmt.Println("Código do produto ", item.Nome, ": ", item.codigo)
	}else{
		fmt.Println("Você não tem permissão para visualizar códigos de itens!");
	}	
}

func (item* Item) CadastrarCodigo(usuario usuario.Usuario, codigo int){

	if(usuario.Perfil == "admin"){
		item.codigo = codigo
		fmt.Println("Código do produto cadastrado")
	}else{
		fmt.Println("Você não tem permissão para cadastrar códigos de itens!");
	}	
}
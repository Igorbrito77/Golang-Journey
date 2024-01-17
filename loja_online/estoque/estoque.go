package estoque;

import(
	"fmt";
	"loja_online/item";
	"loja_online/usuario";
);

// depois ter associação pra se ter o numero de produtos no estoque

type Estoque struct{

	itens [] item.Item;

}


func (estoque* Estoque) CadastraItem ( novo_item item.Item, usuario usuario.Usuario){

	if(usuario.Perfil == "admin"){

		fmt.Println(novo_item.Nome , " adicionado ao estoque");
		estoque.itens = append(estoque.itens, novo_item);

	}else{
		fmt.Println("Você não tem permissão para cadastrar um item");
	}	

}

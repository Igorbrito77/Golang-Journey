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


func (estoque* Estoque) CadastraItens ( novos_itens [] item.Item, usuario usuario.Usuario){

	if(usuario.Perfil == "admin"){

		fmt.Println("\n_________________________________ \n");

		for _, novo_item := range novos_itens{
			fmt.Println(novo_item.Nome , " adicionado ao estoque por ", usuario.Nome);
			estoque.itens = append(estoque.itens, novo_item);
		} 

		fmt.Println("_________________________________ \n\n");

	}else{
		fmt.Println("Você não tem permissão para cadastrar um item");
	}	
}

func (estoque* Estoque) ExibirEstoque( usuario usuario.Usuario){

	// depois listar apenas os produtos em estoque com quantidade maior que 0
	fmt.Println("\n_________________________________ \n");
	fmt.Println("     ESTOQUE:         \n");
	
	if(usuario.Perfil == "admin"){

		for _, item := range estoque.itens{
			fmt.Println(item.Nome , ". Preço: ", item.Preco, ". Código: ", item.Codigo);
		} 

	}else{
		
		for _, item := range estoque.itens{
			fmt.Println(item.Nome , ". Preço: ", item.Preco);
		} 
	
	}
	
	fmt.Println("_________________________________ \n\n");

}

package estoque;

import(
	"fmt";
	"loja_online/item";
	"loja_online/usuario";
);


type ItemEstoque struct{

	Item item.Item;
	UnidadesDisponiveis int;
};


type Estoque struct{

	ItensEstoque [] ItemEstoque;
}


func (estoque* Estoque) CadastraItens ( itens_estoque [] ItemEstoque, usuario usuario.Usuario){

	if(usuario.Perfil == "admin"){

		fmt.Println("\n_________________________________ \n");

		for i, item_estoque := range itens_estoque{
			item_estoque.Item.Id = i + 1;
			fmt.Println(item_estoque.Item.Nome , " adicionado ao estoque por ", usuario.Nome);
			estoque.ItensEstoque = append(estoque.ItensEstoque, item_estoque);
			} 

		fmt.Println("_________________________________ \n\n");

	}else{
		fmt.Println("Você não tem permissão para cadastrar um item");
	}	
}

func (estoque* Estoque) ExibirEstoque( usuario usuario.Usuario){ // talvez criar uma interface para usar essa função de acordo com o perifl do usuário

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println("\n 								**** ESTOQUE ****         \n");
	
	if(usuario.Perfil == "admin"){

		for _, itemEstoque := range estoque.ItensEstoque{

			if(itemEstoque.UnidadesDisponiveis > 0){
				fmt.Println(itemEstoque.Item.Id, ". ",  itemEstoque.Item.Nome , ". Preço: ", itemEstoque.Item.Preco, ". Código: ", itemEstoque.Item.Codigo, "Unidades disponíveis: ", itemEstoque.UnidadesDisponiveis);
			}

		} 

	}else{
		
		for _, itemEstoque := range estoque.ItensEstoque{

			if(itemEstoque.UnidadesDisponiveis > 0){
				fmt.Println(itemEstoque.Item.Id, ". ",  itemEstoque.Item.Nome , ". Preço: ", itemEstoque.Item.Preco, "Unidades disponíveis: ", itemEstoque.UnidadesDisponiveis);
			}
		} 
	
	}
	
}


func (estoque* Estoque) InicializarEstoque(userDefault usuario.Usuario) Estoque{
	
	estoqueInicio := Estoque{};

	cafe:= item.Item{Nome: "Café", Preco: 4.50, Codigo: 1};
	livro:= item.Item{Nome: "Livro", Preco: 10.50, Codigo: 2};
	kit_ferramentas:= item.Item{Nome: "Kit de Ferramentas", Preco: 74.99, Codigo: 3};


	cafe_qtd := ItemEstoque{ Item: cafe, UnidadesDisponiveis: 100 };
	livro_qtd := ItemEstoque{ Item: livro, UnidadesDisponiveis: 50 };
	kit_ferramentas_qtd := ItemEstoque{ Item: kit_ferramentas, UnidadesDisponiveis: 200 };


	itens_novos := [] ItemEstoque{ cafe_qtd, livro_qtd, kit_ferramentas_qtd};

	fmt.Println(" ... Inicializando Estoque ...");

	estoqueInicio.CadastraItens(itens_novos, userDefault);

	return estoqueInicio;
}


func (estoque* Estoque) RetornarItensDisponiveis() [] ItemEstoque{

	
	var itens_disponiveis [] ItemEstoque;

	for _, itemEstoque := range estoque.ItensEstoque{

		if(itemEstoque.UnidadesDisponiveis > 0){
			itens_disponiveis = append(itens_disponiveis, itemEstoque);
		}
	} 
	
	return itens_disponiveis;
}
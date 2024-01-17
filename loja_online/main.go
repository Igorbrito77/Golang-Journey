package main;

import(
	"fmt";
	"loja_online/usuario";
	"loja_online/item";
	"loja_online/estoque";
);


func main(){

	estoque := estoque.Estoque{};

	fmt.Println(estoque);

	var igor *usuario.Usuario = new(usuario.Usuario);

	igor.Nome = "Igor";
	igor.Perfil = "admin";

	fmt.Println(*igor);


	cafe:= item.Item{Nome: "Café", Preco: 4.50};
	//livro:= item.Item{Nome: "Livro", Preco: 10.50};

	cafe.CadastrarCodigo(*igor, 1234);
	cafe.ExibirCodigo(*igor)


	estoque.CadastraItem(cafe, *igor);

	fmt.Println(estoque);




	//meu_carrinho := Carrinho{ Itens: []Item{cafe, livro} };

	/*exibir_carrinho(meu_carrinho);*/

}
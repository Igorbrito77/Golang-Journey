package main;

import(
	//"fmt";
	"loja_online/usuario";
	"loja_online/item";
	"loja_online/estoque";
);


func cadastarItensIniciais(user usuario.Usuario) estoque.Estoque{
	
	estoque := estoque.Estoque{};

	cafe:= item.Item{Nome: "Café", Preco: 4.50, Codigo: 1};
	livro:= item.Item{Nome: "Livro", Preco: 10.50, Codigo: 2};
	kit_ferramentas:= item.Item{Nome: "Kit de Ferramentas", Preco: 74.99, Codigo: 3};

	itens_novos := [] item.Item{ cafe, livro, kit_ferramentas}

	estoque.CadastraItens(itens_novos, user);

	return estoque;
}

func main(){

	user := usuario.Usuario{Nome: "User",  Perfil: "admin"};

	estoque := cadastarItensIniciais(user);

	var igor *usuario.Usuario = new(usuario.Usuario);

	igor.Nome = "Igor";
	igor.Perfil = "admin";

	estoque.ExibirEstoque(*igor);

	cafe:= item.Item{Nome: "Café", Preco: 4.50, Codigo: 1};

	igor.AdicionarItemCarrinho(cafe, 2);

	igor.ExibirCarrinho();

}
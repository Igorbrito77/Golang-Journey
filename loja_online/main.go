package main;

import(
	"fmt";
	"os";
	"loja_online/loja";
	//"loja_online/usuario";
	"loja_online/item";
	//"loja_online/estoque";
);

/*
func cadastarItensIniciais(user usuario.Usuario) estoque.Estoque{
	
	estoque := estoque.Estoque{};

	cafe:= item.Item{Nome: "Café", Preco: 4.50, Codigo: 1};
	livro:= item.Item{Nome: "Livro", Preco: 10.50, Codigo: 2};
	kit_ferramentas:= item.Item{Nome: "Kit de Ferramentas", Preco: 74.99, Codigo: 3};

	itens_novos := [] item.Item{ cafe, livro, kit_ferramentas}

	estoque.CadastraItens(itens_novos, user);

	return estoque;
}*/

func main(){	
	
	loja:= loja.Loja{};

	loja.CarregarDadosIniciais();

	fmt.Println(loja)

	exibirMenuCliente(loja);

/*
	var igor *usuario.Usuario = new(usuario.Usuario);

	igor.Nome = "Igor";
	igor.Perfil = "admin";

	estoque.ExibirEstoque(*igor);

	cafe:= item.Item{Nome: "Café", Preco: 4.50, Codigo: 1};
	livro:= item.Item{Nome: "Livro", Preco: 10.50, Codigo: 2};
	kit_ferramentas:= item.Item{Nome: "Kit de Ferramentas", Preco: 74.99, Codigo: 3};

	igor.AdicionarItemCarrinho(cafe, 2);
	igor.AdicionarItemCarrinho(livro, 4);
	igor.AdicionarItemCarrinho(kit_ferramentas, 1);


	igor.ExibirCarrinho();*/

}

func exibirMenuCliente(loja loja.Loja){


	fmt.Println("'**** LOJA ONLINE *** \n");

	novoUsuario := loja.CadastrarCliente();	

	loja.ListarClientes();

	var comando int;

	for{

		fmt.Println("Escolha uma opção: \n")
		fmt.Println("1- Listar itens para compra");
		fmt.Println("2- Exibir carrinho de compras");
		fmt.Println("0- Sair");

		fmt.Scan(&comando);


		switch(comando){
			
			case 1:
				loja.Estoque.ExibirEstoque(novoUsuario);

				var desejaEsolher string; 

				fmt.Println("Deseja adicionar algum item ao carrinho de compras? (S = sim, N = Não)");

				fmt.Scan(&desejaEsolher);

				if(desejaEsolher == "S"){

					fmt.Println("Digite o id do item que você deseja: ");

					var idItemEscolhido int;
					var idItemValido bool = false;
					var itemEscolhido item.Item;

					fmt.Scan(&idItemEscolhido);


					var idsItens [] int;

					for _, item:= range loja.Estoque.Itens{

						idsItens = append(idsItens, item.Id); 

						if(item.Id == idItemEscolhido){
							itemEscolhido = item;
							idItemValido = true;
						}
					}

					//fmt.Println("ids === ", idsItens);
					fmt.Println("item Escolhido === ", itemEscolhido);

					if(idItemValido){

						var quantidadeItem int;

						fmt.Println("Digita e quantidade deseja deste item: ");
						fmt.Scan(&quantidadeItem);

						novoUsuario.AdicionarItemCarrinho(itemEscolhido, quantidadeItem);

						fmt.Println("novoUsuario == ", novoUsuario);

						fmt.Println("loja == ", loja);


					}else{
						fmt.Println("Id inválido");
					}


				}else{
					fmt.Println("Otário tu, hein?");
				}

				//fmt.Scan(&escolhaItem);



				

			case 3:
				return;
			default:
				os.Exit(0);

		}

	}





}


func exibirMenu(){


	fmt.Println("'**** LOJA ONLINE *** \n");

	var comando int;

	for{

		fmt.Println("Escolha uma opção: \n")
		fmt.Println("1- Cadastrar novo item");
		fmt.Println("2- Atualizar quanridade de itens");
		fmt.Scan(&comando);


		switch(comando){
			
			case 1:
				fmt.Println("deu bom");
			case 3:
				return;
			default:
				os.Exit(0);

		}

	}





}


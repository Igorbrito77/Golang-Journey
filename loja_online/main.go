package main;

import(
	"fmt";
	"loja_online/loja";
	//"loja_online/usuario";
	//"loja_online/item";
	"loja_online/estoque";
);

/*
	Novas funcionalidades que podem ser adiconadas futuramente:

	* Mostrar dados do usuário logado
	* Listar as compras finalizadas pelo usuário
	* Permitir usuários de perfil administrador adicionarem itens ao estoque durante a execução do programa
	* Permitir trocas de usuário durante a execução do programa 

	* Mostrar mensagem de compra efeutado com sucesso
	* Formatar status da loja 
*/


func main(){	
	
	loja:= loja.Loja{};

	loja.CarregarDadosIniciais();

	fmt.Println("ESTADO INICIAL DA LOJA ->", loja);

	exibirMenuCliente(&loja);

	fmt.Println("ESTADO FINAL DA LOJA -> ", loja);
}

func exibirMenuCliente(loja * loja.Loja){

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println("\n 								**** LOJA ONLINE *** \n");

	loja.CadastrarCliente();	
	
	var novoUsuario = &(loja.Clientes[len(loja.Clientes) - 1] ); // referencia o endereço do último cliente cadastrado na loja
	//loja.ListarClientes();

	var comando int;

	for{

		fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
		fmt.Println(" 								**** MENU PRINCIPAL **** \n")
		fmt.Println("1- Listar itens para compra");
		fmt.Println("2- Exibir carrinho de compras");
		fmt.Println("3- Finalizar compra");
		fmt.Println("0- Sair");
		fmt.Println(" 								Escolha uma opção: \n")

		fmt.Scan(&comando);

		switch(comando){
			
			case 1:

				loja.Estoque.ExibirEstoque(*novoUsuario);

				var desejaEsolher string; 

				fmt.Println("\n\nDeseja adicionar algum item ao carrinho de compras? (S = sim, N = Não)");
				fmt.Scan(&desejaEsolher);

				if(desejaEsolher == "S"){

					fmt.Println("Digite o id do item que você deseja: ");

					var idItemEscolhido int;
					var idItemValido bool = false;
					var itemEscolhidoEstoque estoque.ItemEstoque;

					fmt.Scan(&idItemEscolhido);

					var itens_disponiveis_estoque = loja.Estoque.RetornarItensDisponiveis();

					for _, item_disponivel_estoque:= range itens_disponiveis_estoque{

						if(item_disponivel_estoque.Item.Id == idItemEscolhido){
							itemEscolhidoEstoque = item_disponivel_estoque;
							idItemValido = true;
						}
					}


					if(idItemValido){

						var quantidadeItem int;

						fmt.Println("Digite a quantidade desejada deste item: ");
						fmt.Scan(&quantidadeItem);

						if(quantidadeItem <= itemEscolhidoEstoque.UnidadesDisponiveis){
							novoUsuario.AdicionarItemCarrinho(itemEscolhidoEstoque.Item, quantidadeItem);
						}else{
							fmt.Println("\n\n                                    ... Não há tantos produtos em estoque, digite um número igual ao inferior ao disponível do estoque ...");
						}
				

					}else{
						fmt.Println("Id inválido");
					}

				}

			case 2:

				novoUsuario.ExibirCarrinho();

			case 3:

				var carrrinhoCheio = novoUsuario.VerificarCarrinhoCheio();

				if(carrrinhoCheio){
					loja.FinalizarCompra(novoUsuario);
				}else{
					fmt.Println("                                    ... Seu carrinho de compras está vazio. Selecione algum item para finalizar uma compra ...")
				}
				
			default:
				fmt.Println("                                                       ... Saindo da Loja Online. Até breve! ...")
				fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
				return;

		}

	}

}


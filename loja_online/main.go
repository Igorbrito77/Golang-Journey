package main;

import(
	"fmt";
	"os";
	"loja_online/loja";
	//"loja_online/usuario";
	//"loja_online/item";
	"loja_online/estoque";
);



func main(){	
	
	loja:= loja.Loja{};

	loja.CarregarDadosIniciais();

	fmt.Println(loja) // ESTADO INICIAL DA LOJA

	exibirMenuCliente(&loja);

	fmt.Println("na main -> ", loja); // ESTADO FINAL DA LOJA
}

func exibirMenuCliente(loja * loja.Loja){

	fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")
	fmt.Println("\n 								**** LOJA ONLINE *** \n");

	loja.CadastrarCliente();	
	
	//var novoUsuario = &(loja.UltimoCliente);
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

		//fmt.Println("\n\n____________________________________________________________________________________________________________________________________________\n")

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
							fmt.Println("Não há tantos produtos em estoque, digite um número igual ao inferior ao disponível do estoque");
						}
				
						//fmt.Println("loja == ", loja);


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
					fmt.Println("Seu carrinho de compras está vazio. Selecione algum item para finalizar uma compra");
				}
				
			default:
				return;

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


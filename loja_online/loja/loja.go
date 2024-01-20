package loja;

import(
	"fmt";
	"loja_online/estoque";
	"loja_online/usuario";
);

// depois ter associação pra se ter o numero de produtos no estoque

type Loja struct{

	Estoque estoque.Estoque;
	Clientes [] usuario.Usuario;
	admins [] usuario.Usuario;
	UltimoCliente usuario.Usuario;
}

func (loja* Loja) CarregarDadosIniciais(){
	
	user := usuario.Usuario{Nome: "Default User",  Perfil: "admin"};

	loja.admins = append(loja.admins, user);

	estoque:= estoque.Estoque{};

	loja.Estoque = estoque.InicializarEstoque(user);

}


func (loja* Loja) CadastrarCliente(){

	var novoCliente = usuario.Usuario{};

	//var nome_usuario string;

	fmt.Println("Digite seu nome: \n")
	fmt.Scan(&novoCliente.Nome);

	fmt.Println("\n Bem-vindo(a), ", novoCliente.Nome, "\n\n _________________");

	loja.Clientes = append(loja.Clientes, novoCliente);	

	loja.UltimoCliente = novoCliente;

	//return novoCliente;
}

func (loja* Loja) ListarClientes(){

	for i, cliente := range loja.Clientes{

		fmt.Println("Clientes cadastrados: \n")

		fmt.Println("id: ", i+1, ". Nome: ", cliente.Nome);

	}

	fmt.Println("\n _________________\n\n");

}

func (loja * Loja) FinalizarCompra(novoUsuario usuario.Usuario){


	for _, item_escolhido := range novoUsuario.ItensEscolhidos{

		for _, loja_item := range loja.Estoque.ItensEstoque {

			if(loja_item.Item.Id == item_escolhido.Item.Id){

				fmt.Println("item_escolhido== ", item_escolhido);

				fmt.Println("loja_item== ", loja_item);


				var alterador = &(loja.Estoque.ItensEstoque[loja_item.Item.Id - 1] );

				fmt.Println("alterador == ", alterador);

				(*alterador).UnidadesDisponiveis -= item_escolhido.Quantidade;
			}
		}
	}

	novoUsuario.ItensEscolhidos = nil;

	

}
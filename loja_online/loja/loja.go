package loja;

import(
	"fmt";
	"loja_online/estoque";
	"loja_online/usuario";
);

// depois ter associação pra se ter o numero de produtos no estoque

type Loja struct{

	Estoque estoque.Estoque;
	clientes [] usuario.Usuario;
	admins [] usuario.Usuario;
}

func (loja* Loja) CarregarDadosIniciais(){
	
	user := usuario.Usuario{Nome: "Default User",  Perfil: "admin"};

	loja.admins = append(loja.admins, user);

	estoque:= estoque.Estoque{};

	loja.Estoque = estoque.InicializarEstoque(user);

}


func (loja* Loja) CadastrarCliente() usuario.Usuario{

	var novoCliente = usuario.Usuario{};

	//var nome_usuario string;

	fmt.Println("Digite seu nome: \n")
	fmt.Scan(&novoCliente.Nome);

	fmt.Println("\n Bem-vindo(a), ", novoCliente.Nome, "\n\n _________________");

	loja.clientes = append(loja.clientes, novoCliente);	

	return novoCliente;
}

func (loja* Loja) ListarClientes(){

	for i, cliente := range loja.clientes{

		fmt.Println("Clientes cadastrados: \n")

		fmt.Println("id: ", i+1, ". Nome: ", cliente.Nome);

	}

	fmt.Println("\n _________________\n\n");

}
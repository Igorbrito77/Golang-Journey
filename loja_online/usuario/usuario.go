package usuario;

import(
	"fmt";
	"loja_online/item";
);

type ItemUsuario struct{

	Item item.Item;
	Quantidade int;
};

type Usuario struct{

	Nome string;
	cpf string;
	Perfil string;

	ItensEscolhidos [] ItemUsuario;
};

func (usuario* Usuario) ExibirCarrinho(){

	for _, item_escolhido := range usuario.ItensEscolhidos{

		fmt.Println(" **** CARRINHO DE COMPRAS DO ", usuario.Nome, " ****    \n");

		fmt.Println("* ", item_escolhido.Item.Nome, " . Quantidade: ", item_escolhido.Quantidade, " . Preço: ", item_escolhido.Item.Preco);

	}

}
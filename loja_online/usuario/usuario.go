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



func (usuario* Usuario) AdicionarItemCarrinho(item_escolhido item.Item, quantidade int){

	novo_item_escolhido := ItemUsuario{ Item: item_escolhido, Quantidade: quantidade};

	usuario.ItensEscolhidos = append(usuario.ItensEscolhidos, novo_item_escolhido);
	
}

func (usuario* Usuario) ExibirCarrinho(){

	fmt.Println("\n_________________________________ \n");
	fmt.Println(" **** CARRINHO DE COMPRAS DO ", usuario.Nome, " ****    \n");

	var valor_total = 0.;

	if(len(usuario.ItensEscolhidos) == 0 ){
		fmt.Println(" ... Carrinho de compras vazio ...")
	}

	for _, item_escolhido := range usuario.ItensEscolhidos{
		
		fmt.Println("* ", item_escolhido.Item.Nome, " . Preço: ", item_escolhido.Item.Preco, " . Quantidade: ", item_escolhido.Quantidade);
		valor_total +=  ( item_escolhido.Item.Preco * float64(item_escolhido.Quantidade) );
	}

	fmt.Println("Valor total: ", valor_total);

	fmt.Println("_________________________________ \n\n");
}
package main;

import(
	"fmt"
);

type Produto struct{
	Nome string;
	Preco float64;
};

type Carrinho struct{
	Produtos [] Produto;
};




func exibir_carrinho(meu_carrinho Carrinho){

	fmt.Println("         *** Carrinho de Compras *** \n")

	valor_total := 0.000;

	for _, produto:= range meu_carrinho.Produtos{
		fmt.Println(produto.Nome, ". Preco: R$", produto.Preco);
		valor_total += produto.Preco;
	}


	fmt.Println("TOTAL: R$", valor_total);

}

func main(){

	cafe:= Produto{Nome: "Café", Preco: 4.50};
	livro:= Produto{Nome: "Livro", Preco: 10.50};


	meu_carrinho := Carrinho{ Produtos: []Produto{cafe, livro} };

	exibir_carrinho(meu_carrinho);

}
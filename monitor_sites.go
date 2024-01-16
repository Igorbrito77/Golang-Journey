package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibir_introducao()

	for {

		comando := inserir_comando()

		switch comando {

		case 1:
			iniciar_monitoramento()
		case 2:
			fmt.Println("Exibindo logs ...")
		case 0:
			fmt.Println("Saindo ...")
			os.Exit(0)
		default:
			fmt.Println("Comando inválido ...")
			os.Exit(-1)
		}
	}

}

func exibir_introducao() {

	nome_usuario := "Igor"
	var versao_software float32 = 1.00

	fmt.Println("Saudações, ", nome_usuario, "! Tenha um excelente dia, meu nobre.")
	fmt.Println("A versão deste software: ", versao_software)

}

func inserir_comando() int {

	var comando int

	fmt.Println("1- Inicar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	fmt.Scan(&comando)

	return comando
}

func iniciar_monitoramento() {

	fmt.Println("Monitorando ...")
	
	sites := []string{"https://random-status-code.herokuapp.com/app", "https://www.alura.com.br", 
						"https://www.caelum.com.br"}


	for i, site:= range sites{
		fmt.Println("Testando site ", i, ": ", site);

		testar_site(site);

	}


	fmt.Println("sites: ", sites)

}


func testar_site(site string){

	//resp, err := http.Get(site);
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}
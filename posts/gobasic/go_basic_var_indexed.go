package main

import "fmt"

func main(){

	var (
		listaNumStr [5]string
		listaNumInt [5]int
	)

	listaNumStr[0] = "um"
	listaNumStr[1] = "dois"
	listaNumStr[2] = "três"
	listaNumStr[3] = "quatro"
	listaNumStr[4] = "cinco"

	listaNumInt[0] = 1
	listaNumInt[1] = 2
	listaNumInt[2] = 3
	listaNumInt[3] = 4
	listaNumInt[4] = 5

	fmt.Print("---Arrays-----------------------------------------------------------\n")
	/*
		Aqui vemos uma forma diferente de usar for para interagir com variáveis indexadas
	*/
	for _, valor := range listaNumStr{
		fmt.Println(valor)
	}

	/*
		Caso queira ter acesso aos indices das listas substitua "_" por uma variável
	*/
	for ind, valor := range listaNumInt{
		fmt.Println("listaNumInt[",ind,"] = ",valor)
	}

	fmt.Print("---Slices-----------------------------------------------------------")

	/*
		Exibindo uma fatia criada apartir de um array "listaNumStr"
		[1:4]
	    começando de 1 até o indice 4-1
		listaNumStr[1]
		listaNumStr[2]
		listaNumStr[3]
	*/
    var fatia []string = listaNumStr[1:4]
	fmt.Println("\n")
    for ind, valor := range fatia{
    	fmt.Println("fatia[",ind,"] = ",valor)
	}

	fmt.Print("---Maps-----------------------------------------------------------")

    /*
		No mapa nós declaramos tanto o tipo do indice quanto o tipo do valor das variáveis.
    	Também não precisamos determinar o tamanho do mapa como fazemos com o array.
    */

	var m map[string]int
	m = make(map[string]int)
	m["um"] = 1;
	m["dois"] = 2;
	m["três"] = 3;
	m["quatro"] = 4;
	m["cinco"] = 5;

	fmt.Println("\n")
	for ind, valor := range m{
		fmt.Println("mapa[",ind,"] = ",valor)
	}


}
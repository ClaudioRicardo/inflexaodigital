package main

import "fmt"

func main(){

	/*
		Essa é uma forma bem simples de se usar for, uma estrutura de repetição
	    que executará um bloco de instruções até que a condição explicitada na declaração seja
	    satisfeita.
		Quando o valor de i, que está sendo incrementando ao final de cada interação i++, chegar até
		10 as repetições vão parar.
	*/
	for i := 0; i < 10; i++ {
		fmt.Println("Instrução nº",i )
	}
}
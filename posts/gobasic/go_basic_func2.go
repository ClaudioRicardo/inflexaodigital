package main

import "fmt"


func main(){

	var x int
	/*
		Encapsulando uma lógica em uma função dentro de outra função,
		no intuito de reaproveitamento de código.
	*/
	incrementa := func() int{
		x++
		return x
	}

	/*
		Função anonima.
		Os 2 exemplos a seguir mostram a utilização de funções anônimas.
	*/
	func(){
		fmt.Println(incrementa())
	}() //O uso dos parenteses no final da função fazer a execução imediata da função.

	fmt.Println(func() string{return "Função anônima incrementa:"}(),incrementa())

	/*
		Função com entrada variádicas.
		Teste utilizando quantidades diferentes de argumentos
	*/
	somatudo := func (args ...int) int{
		total := 0
		for _, v := range args{
			total +=v
		}
		return total
	}


	fmt.Println(somatudo(1,2,3,4,5))




}

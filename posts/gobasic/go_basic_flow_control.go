package main

import "fmt"

func main(){
	/*
		Declarando variáveis
		Dois inteiros e uma variável que representa a operação
	*/
	var(
		ValA int8 = 10
		ValB int8 = 32
		Op string = ""
	)

	/*
	  Usando condicionais if e else para criar dois fluxos
	  respeitando uma simples sentença comparativa (valA > valB)
	  se a comparação for verdadeira será exibida uma soma senão uma subtração
	*/
	if(ValA > ValB){
		fmt.Println("Soma: ", ValA + ValB)
	}else{
		fmt.Println("Subtração: ", ValB - ValA)
	}


	/*
	 Usando switch para criar diversar possibilidades condicionais e uma padrão para responder caso o teste
	 não passe por nehuma das "case"

	Caso Op seja igual a "+" : Uma soma será realizada
	Caso Op seja igual a "-" : Uma subtração será realizada
	Caso Op não correspondam a nenhum dos caso listados a operação default será executada
	*/

	switch Op {
		case "+":
			fmt.Println("(Switch) Soma: ", ValA + ValB)
		case "-":
			fmt.Println("(Switch) Subtração: ", ValA - ValB)
		default:
			/*
			   Imprimi uma resposta default(padrão) caso a variável verificada não tenha um caso que corresponda ao valor.
			   Aqui eu fiz uso de uma conversão de tipo para int16, porque a multiplicação de 10 * 32 = 320 que supera
			   a representação de um tipo int8 (-128 até 127)
			*/
			fmt.Println("(Switch) Multiplicação: ", int16(ValA) * int16(ValB) )
	}

	/*
		"defer" é uma instrução que serve para mudar a ordem de execução de um instrução postergando a sua execução sempre
		para o final da pilha de instruções.
	*/
	fmt.Println("------------------------------------------------------------------------------")
	fmt.Println("(defer) Soma: ", ValA + ValB)
	defer fmt.Println("(defer) Subtração: ", ValA - ValB)
	fmt.Println("(defer) Multiplicação ", int16(ValA) * int16(ValB))
	fmt.Println("(defer) Divisão: ", ValB % ValA)

	/*
		Essa é a saída desse ultimo bloco de código.
		Repare que a subtração que deveria ser a segunda instrução a ser executada
		foi executada por ultimo graças ao comando "defer" que mudou a ordem e a pos no final da
		pilha de execução.
		------------------------------------------------------------------------------
		(defer) Soma:  42
		(defer) Multiplicação  320
		(defer) Divisão:  2
		(defer) Subtração:  -22

	*/






}

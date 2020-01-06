package main

import (
	"fmt"
	"time"
)

func main(){

	/*
	Declarando um canal.
	Um canal dever ser declarado utilizando
	"chan"(canal) "tipo"(tipo de informação que trafegará nesse canal)
	*/
	var c chan string = make(chan string)
	var x string = ""
	/*
	Criando rotinas concorrentes anônimas para exemplificar
	uma comunicação simples.
	*/

	go func() {
		/*
			Enviando informação pelo canal "c" utilizando operador de envio/recebimento  "<-"
		*/
		c <- "Passando por A"
		x = "Teste 3"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		/*
			Imprimindo o recebimento da informação do canal "c"
		*/
		fmt.Println("Imprimindo x: ", x)
		fmt.Println("Rotina 2 Imprimindo mensagem da Rotina 1: \"",<-c, "\"")
	}()

	time.Sleep(1000 * time.Millisecond)
}




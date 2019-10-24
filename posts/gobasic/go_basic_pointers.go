package main

import "fmt"


func main(){
	fmt.Println("Ponteiros----------------------------")

	var (
		a int = 5
		b int = 10
	)

	/*
		Criando os ponteiros p1 e p2 referências
		respectivas as variáveis  a e b
		Essa forma de declaração utilizando utilizando "&" significa
		que estamos criando em "p1" uma referência a variável "a"
	*/
	p1 := &a
	p2 := &b

	/* Entendendo as diferenças */
	fmt.Println("Variável a =", a)
	fmt.Println("Variável b =", b)

	/*
		Imprimindo os valores  ponteiros
		Para visualizar e trabalhar com os valores dos ponteiros tempos
		que utilizar o "*" como mostrado abaixo
	*/
	fmt.Println("Ponteiro Valor a =", *p1)
	fmt.Println("Ponteiro Valor b =", *p2)

	/*
		Para melhor entender o conceito de referência, a impressão a
		seguir vai exibir no console as referências a memória de a e b
		Quando utilizamos o ponteiro se o uso de "*" estaremos mostrando
		O endenreço de memória da variável original.
	*/
	fmt.Println("Ponteiro referência a =", p1)
	fmt.Println("Ponteiro referência b =", p2)

	/*
		Para mostrar a ligação entre variável origem
		e ponteiro o código abaixo ira mostrar , que
		quando alteramos o valor no ponteiro a alteração
		acontece na origem e vice e versa.
	*/

	*p1 = 7

	fmt.Println("Variável a =", a)
	fmt.Println("Ponteiro referência a =", *p1)

	a = 9

	fmt.Println("Variável a =", a)
	fmt.Println("Ponteiro referência a =", *p1)





}



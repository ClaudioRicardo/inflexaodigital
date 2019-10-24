package main

import "fmt"

type Animal interface{
	falar() string
	getNome() string
}

type Cachorro struct {
	Nome string
}

func (c Cachorro) falar() string{
	return "AuAu!!"
}

func (c Cachorro) getNome() string{
	return c.Nome
}

type Gato struct{
	Nome string
}

func (g Gato) getNome() string{
	return g.Nome
}

func (g Gato) falar() string{
	return "Miauu!!"
}

func main(){
	/*
		Vamos utilizar um dos exemplos mais clássicos do uso de interfaces
	    e polimorfismo da programação.

	    "Animal"
		"Cachorro"
		"Gato"

		Repare que a mesma variável do tipo Animal(Interface) aceita os tipos
	    Cachorro e Gato.
	    Quando "animal" assume a identidade cachorro ao usar os métodos
	    getNome e falar automáticamente elas assumem o comportamento implementado pela
	    estrutura Cachorro.

		E o mesmo acontece quando "animal" assume a identidade Gato os métodos respondem a
	    implementação da estrutura gato.

		Para melhor entedermos a importância desse conceito tente declar a variável "animal"
	    como sendo "Cachorro" e tente inicializá -la como "Gato".O sistema vai dizer que
	    são tipos diferente e que o tipo "Cachorro" não pode receber "Gato".

		Se por exemplo você estiver desenvolvendo um sistema veterinário para cadastrar as fichas
	    de diversos animais, criar uma interface facilitaria a inserção de um novo tipo de animal
	    não só como um novo registro, mas como novo tipo com comportamentos particulares.

	*/

		fmt.Println("Interface -----------------------------------------")
		var animal Animal = Cachorro{"Rex"}
		fmt.Println("Animal ",animal.getNome(), animal.falar())
		animal = Gato{"Tigrão"}
		fmt.Println("Animal ",animal.getNome(), animal.falar())



}

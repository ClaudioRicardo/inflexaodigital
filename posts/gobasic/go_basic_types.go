package main

import "fmt"

/*
Podemos declarar variáveis de diversas formas
agrupadas usando var(...), ou individualmente usando var v string , ou simplesmente
atribuindo um valor v:="Teste"
Observe abaixo:
*/

// Em grupo
var (
	nome string = "José" // determinando um tipo e atribuindo um valor
	idade int  // determinando um tipo ,sem atribuição de valor
	casado bool
)

func main(){

	idade = 32
	casado = true
	/*
		carteira não foi declarado previamente.Nesse caso
	    ao atribuirmos um valor usando := a variável passa a ser do tipo atribuido
	    carteira := 15.35
	    carteira agora é uma variável tipo float
	*/
	carteira := 15.35

	//O mesmo ocorre com estado_civil agora é do tipo string
	estado_civil := "solteiro"

	/*
		Mais a frente falaremos sobre estruturas de controle então
		não se incomode com esse bloco agora
	*/
	if(casado == true){
		estado_civil = "casado"
	}


	fmt.Println(nome ," tem " , idade , " anos de idade e é ", estado_civil,".")
	fmt.Println(nome, " tem R$",carteira, " na carteira.")
}

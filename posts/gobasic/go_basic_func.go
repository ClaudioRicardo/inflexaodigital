package main

import "fmt"

func main(){
	/*
		Funções servem para encapsular conjuntos de instruções que podem ser reutilizadas no código do seu sistemas.
	*/

	resultado := soma(32,13)
	fmt.Println("Soma de dois valores = ", resultado )

	v, p := par(6)
	fmt.Println("Valor ", v ," é par? ",p)

	msg("Nem todas as variáveis vão ter valor de retorno.\n")
}

/*
	Quando declaramos uma função, devemos indicar suas variáveis de entrada e seus respectivos tipos ,caso existam.
	Também devemos indicar o tipo da sáida caso exista.
	val1 e val2 são do tipo int e são as entradas, em seguida após as entradas entre, ) int {, indicamos o tipo da saída.
*/
func soma(val1 int, val2 int) int{
	return  val1 + val2
}
/*
 Função com multiplos retornos.
 Vai checar se o valor é par e vai retornar o valor passado e o booleano indicando se o valor é
 par ou não.
*/

func par(val uint) (uint, bool){
	var(
		teste uint
		p bool = false
	)

	teste = val % 2

	if(teste == 0){
		p = true
	}
	return val, p
}

/*
	Nem todas as funções vão ter um valor de retorno
*/

func msg(m string){
	fmt.Println(m)
}

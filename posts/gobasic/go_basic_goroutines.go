package main
import(
	"fmt"
	"time"
)
/*
	Vamos tentar entender, basicamente como funciona concorrência simulando uma corrida.
*/

func corrida(carro string){
	/*
		Essa função vai simular o comportamento de uma carro em uma corrida de 30 voltas.
		Vamos usar a função Sleep para gerar um "delay" de 1 segundo entre os prints.
		Esses delays também serão utilizados pelo escalonador da linguagem como
		critério de interrupção para alternancia entre as rotinas.
	*/
	for i:=0; i <=30; i++  {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(fmt.Sprintf( "%s : %d voltas", carro, i))
	}
}

func main(){
	/*
		Em Go para executar uma chamada concorrente usamos o operador "go" antecedendo
		a chamada da função.Abaixo fizemos uma chamada concorrente para cada carro.
		Quando as impressões começarem, vamos conseguir entender que a cada volta de
		cada carro a ordem do carros muda indicando o que seriam as ultrapassagens da corrida.
	    Assim deixando claro como as funções estão concorrendo para conseguir
		imprimir seu nome e o númeor de voltas na tela.
	*/
	go corrida("Ferrari")
	go corrida("Mercedes")
	go corrida("RBR")
	go corrida("BMW")
	go corrida("Renault")


	/*
		Para conseguir ver os prints no final da execução temos que vazer com que a função principal "main"
	    espere o termino das "goroutines".
		Podemos fazer isso de diversos modos.Aqui nesse exemplo eu optei por deixar um delay com tempo um
	    pouco superior ao tamanho da corrida 35 segundos
	    Uma segunda opção comentada abaixo é utilizar um Scanln que gera uma operação de entrada que
	    espera que o usuário digite algo para que o programa seja concluido.
	*/

	/*
		Execute o programa e observe o comportamento da saída.
	*/

	//var input string
	//fmt.Scanln(&input)
	time.Sleep(35 * time.Second)
}

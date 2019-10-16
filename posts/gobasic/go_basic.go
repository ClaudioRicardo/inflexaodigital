package main
/*
  Go é modular e usa o conceito de pacotes para organizar o sitema e
  ajudar na reutilização do código do mesmo.
  Então a primeira coisa a se fazer quando criamos um arquivo .go e declarar
  o pacote do qual ele faz parte.No caso desse arquivo "package main"
*/
import "fmt"
/*
  import é utilizado para chamar outros pacotes da linguagem,de terceiros e pacotes
  que você tenha criado.Nesse caso "fmt" é um pacote para formatação de dados de saída.
*/

/*
  main() é a função principal de um arquivo .go ela é a espinha dorsal do programa.
  Atravez dela que as demais funções,declarações de variáveis, tipo e métodos serão executados.
*/
func main(){
	fmt.Println("Hello, World")
}
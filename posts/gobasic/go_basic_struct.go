package main

import "fmt"

/*
	Declarando o tipo:
       Autor é uma estrutura composta
       por um campo "Codigo" int
       e   um campo "Nome" string
*/
type Autor struct{
	Codigo int
	Nome string
}

/*
	Declarando o tipo:
       Livro é uma estrutura composta
       por um campo "ISBN" string
           um campo "Nome" string
       e   uma composição da estrutura "Autor"
*/
type Livro struct {
	ISBN string
	Nome string
	Autor
}

/*
	Métodos setAutor e getAutor referentes a estrutura Livro
	Eu implementei esses dois métodos simplesmente para ilustrar
    o conceito.
*/
func (l *Livro) setAutor(a Autor){
	//l.Autor.Nome = a.Nome
	//l.Autor.Codigo = a.Codigo

	l.Autor = a
}

func (l *Livro) getAutor() Autor{
	return l.Autor
}

func main(){
	/*
		A função new() cria uma nova "instância" da
		estrutura passada como parâmetro.
	*/
	autor := Autor{1,"Isaac Asimov"}


	/*
		Criando duas variáveis do tipo estrutura Livro
		que tem em sua composição uma variável do
		Autor
	*/
	livro := Livro{"000000002","Eu robô",Autor{} }
	/* Setando o autor */
	livro.setAutor(autor)

	livro2 := Livro{"000000002","Eu robô",autor}

	fmt.Println(livro)
	fmt.Println(livro2)
}

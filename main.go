package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======CALCULADORA DE MEDIA DE NOTAS DE ALUNOS======")

	fmt.Print("Digite o nome do aluno: ")
	nome_aluno, _ := reader.ReadString('\n')
	nome_aluno = strings.TrimSpace(nome_aluno)

	fmt.Println("digite 2(duas) notas do aluno: ") // pede as notas
	var nota1 float64
	fmt.Scanln(&nota1)

	if nota1 > 10 || nota1 < 0 { // verificacao da nota
		fmt.Println("nota invalida!")
		return
	}

	var nota2 float64
	fmt.Scanln(&nota2)

	if nota2 > 10 || nota2 < 0 { // passa por outra verificacao com a segunda nota
		fmt.Println("nota invalida!")
		return
	}

	resultado := (nota1 + nota2) / 2 // calcula a media

	fmt.Printf("media de %s: %.2f\n", nome_aluno, resultado) // fala a media do aluno

	if resultado >= 7 { // fala o estado  do aluno
		fmt.Println("aprovado. Continue assim!")
	} else if resultado < 7 && resultado >= 5 {
		fmt.Println("recuperacao. Você ainda tem chances!")
	} else {
		fmt.Println("reprovado. Tente de novo!")
	}
	

}
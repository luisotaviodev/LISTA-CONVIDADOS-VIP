package main

import (
	"fmt"
	"strings"
)

var stlConvidados = []string{"Luis", "Bruna", "João", "Bruno"}

func main() {
	verificarConvidado()
}

func verificarConvidado() {
	var cNome string
	var cSimNao string

	fmt.Print("Olá! Por favor, insira seu nome para verificarmos se você está na lista VIP: ")
	fmt.Scan(&cNome)
	if !ListaConvidados(cNome, stlConvidados) {
		fmt.Print("Desculpe, mas seu nome não está na lista de Convidados VIP.\n")
		fmt.Print("Gostaria de comprar o ingresso na portaria do evento, por apenas R$ 150,00? (Digite SIM ou NAO): ")
		fmt.Scan(&cSimNao)
		CompraIngressoVIP(cSimNao)
	} else {
		fmt.Printf("Muito bem-vindo Sr. %s, pode entrar.\n", cNome)
	}
}

func CompraIngressoVIP(cSimNao string) {
	var cNome string
	if strings.ToUpper(cSimNao) == "SIM" {
		fmt.Print("Por favor, digite seu nome para que ele seja incluído na lista VIP: ")
		fmt.Scan(&cNome)
		stlConvidados = append(stlConvidados, cNome)
		fmt.Print("Muito bem, agora dirija-se à entrada do evento.\n")
		verificarConvidado()
	} else {
		fmt.Print("Tudo bem, obrigado. Tenha uma ótima noite.\n")
	}
}

func ListaConvidados(cNome string, stlConvidados []string) bool {
	for _, nome := range stlConvidados {
		if nome == cNome {
			return true
		}
	}
	return false
}

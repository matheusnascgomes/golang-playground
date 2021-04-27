package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoringTimes = 5
const delay = 3

func main() {
	handleIntro()

	command := readCommand()

	switch command {
	case 1:
		monitoringSite()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando desconhecido")
	}
}

func handleIntro() {
	fmt.Println("SISTEMA DE MONITORAMENTO, ESCOLHA UMA OPÇÃO: ")
	fmt.Println("1 - Monitorar")
	fmt.Println("2 - Exibir log")
	fmt.Println("0 - Sair do programa")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func monitoringSite() {
	fmt.Println("Monitorando...")

	sites := []string{"http://www.google.com.br", "https://golang.org/", "https://github.com/"}

	i := 0
	for {
		i++
		fmt.Println("MONITORAMENTO #", i)
		for j, site := range sites {
			fmt.Println("Monitoramento Núm.", j+1, ":", site)
			handleMonitoring(site)
			fmt.Println("")
		}
		time.Sleep(delay * time.Second)
	}
}

func handleMonitoring(site string) {
	resp, _ := http.Get(site)
	statusCode := resp.StatusCode

	if statusCode == 200 {
		fmt.Println("Site: ", site, " está funcionando perfeitamente. StatusCode: ", statusCode)
	} else {
		fmt.Println("Site ", site, " está com algum problema. StatusCode: ", statusCode)
	}
}

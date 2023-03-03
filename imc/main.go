package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func printConversionError(err error) {
	fmt.Printf("Vixe, você precisa informar um numero inteiro, somente caracteres numéricos. 1 - 9. %v", err.Error())
}

func parseFloat(value string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(value, "\r", ""), "\n", ""), 64)
}

func main() {
	fmt.Println("Bem vinds a CALCULADORA IMC")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Por favor insira sua altura em centímetros, ex: 175")
	heightInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Oh oh, deu ruim lendo sua altura %v", err.Error())
		return
	}

	height, err := parseFloat(heightInput)
	if err != nil {
		printConversionError(err)
		return
	}

	fmt.Println("Por favor insira seu peso(foi mal a indelicadeza rsrsrs) em kg, ex: 65")
	weightInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Oh oh, deu ruim lendo seu peso %v", err.Error())
		return
	}

	weight, err := parseFloat(weightInput)
	if err != nil {
		printConversionError(err)
		return
	}

	imc := weight / math.Pow(height/100, 2)
	fmt.Printf("Seu IMC é %.1f \n", imc)

	if imc < 18.5 {
		fmt.Println("Abaixo do peso, capa da gaita")
	} else if imc < 25 {
		fmt.Println("Peso normal, tanquinho no grau")
	} else if imc < 30 {
		fmt.Println("Pré obesidade, tem pochete")
	} else if imc < 35 {
		fmt.Println("Obesidade grau 1, \"roliço\"")
	} else if imc < 40 {
		fmt.Println("Obesidade grau 2, tua vó não te acha magro mais")
	} else if imc < 25 {
		fmt.Println("Obesidade grau 3, joga de tanker né?")
	}

	fmt.Println("Aperte enter para sair")
	reader.ReadString('\n')
}

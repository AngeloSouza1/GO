
package main

import (
	"fmt"
)

func main() {
	var kelvin float64

	// Solicita ao usuário para inserir a temperatura em Kelvin
	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanln(&kelvin)

	// Converte a temperatura de Kelvin para Celsius
	celsius := kelvin - 273.15

	// Exibe o resultado
	fmt.Printf("A temperatura em Celsius é: %.2f°C\n", celsius)
}

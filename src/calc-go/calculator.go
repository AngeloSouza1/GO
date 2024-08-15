
package main

import (
    "fmt"
)

// Função para adicionar dois números
func add(a, b float64) float64 {
    return a + b
}

// Função para subtrair dois números
func subtract(a, b float64) float64 {
    return a - b
}

// Função para multiplicar dois números
func multiply(a, b float64) float64 {
    return a * b
}

// Função para dividir dois números
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero")
    }
    return a / b, nil
}

func main() {
    var a, b float64
    var operation string

    fmt.Println("Digite o primeiro número:")
    fmt.Scanln(&a)

    fmt.Println("Digite o segundo número:")
    fmt.Scanln(&b)

    fmt.Println("Escolha a operação (+, -, *, /):")
    fmt.Scanln(&operation)

    switch operation {
    case "+":
        fmt.Printf("Resultado: %.2f\n", add(a, b))
    case "-":
        fmt.Printf("Resultado: %.2f\n", subtract(a, b))
    case "*":
        fmt.Printf("Resultado: %.2f\n", multiply(a, b))
    case "/":
        result, err := divide(a, b)
        if err != nil {
            fmt.Println("Erro:", err)
        } else {
            fmt.Printf("Resultado: %.2f\n", result)
        }
    default:
        fmt.Println("Operação inválida")
    }
}

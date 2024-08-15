package main

import (
    "testing"
)

// Teste para a função add
func TestAdd(t *testing.T) {
    result := add(2, 3)
    expected := 5.0
    if result != expected {
        t.Errorf("add(2, 3) = %.2f; esperado %.2f", result, expected)
    }
}

// Teste para a função subtract
func TestSubtract(t *testing.T) {
    result := subtract(5, 3)
    expected := 2.0
    if result != expected {
        t.Errorf("subtract(5, 3) = %.2f; esperado %.2f", result, expected)
    }
}

// Teste para a função multiply
func TestMultiply(t *testing.T) {
    result := multiply(4, 2)
    expected := 8.0
    if result != expected {
        t.Errorf("multiply(4, 2) = %.2f; esperado %.2f", result, expected)
    }
}

// Teste para a função divide com valor normal
func TestDivide(t *testing.T) {
    result, err := divide(10, 2)
    if err != nil {
        t.Errorf("Erro inesperado: %v", err)
    }
    expected := 5.0
    if result != expected {
        t.Errorf("divide(10, 2) = %.2f; esperado %.2f", result, expected)
    }
}

// Teste para a função divide com divisão por zero
func TestDivideByZero(t *testing.T) {
    _, err := divide(10, 0)
    if err == nil {
        t.Error("Esperado erro na divisão por zero, mas não ocorreu erro.")
    }
}

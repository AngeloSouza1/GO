# Conversor de Temperatura Kelvin para Celsius

## Descrição

Este programa em Go realiza a conversão de temperaturas de Kelvin para Celsius. É uma ferramenta útil para transformar a temperatura no ponto de ebulição ou qualquer outra temperatura em Kelvin para a escala Celsius.

## Como Funciona

O programa solicita ao usuário que insira uma temperatura em Kelvin e, em seguida, converte esse valor para Celsius usando a fórmula:

\[ C = K - 273.15 \]

Onde:
- \( C \) é a temperatura em Celsius
- \( K \) é a temperatura em Kelvin

O resultado é exibido ao usuário em um formato de ponto flutuante com duas casas decimais.

## Estrutura do Código

- **`escalas_termo.go`**: Arquivo principal contendo a função `main`, onde o programa solicita a entrada do usuário, realiza a conversão e exibe o resultado.

### Conteúdo do Arquivo `escalas_termo.go`

```go
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
```

## Compilação e Execução

Para compilar e executar o programa, siga estas etapas:

#### 1. Salve o Código:
  -  Salve o código acima em um arquivo chamado `escalas_termo.go`.

#### 2. Compile o Programa:
  -  Abra o terminal e navegue até o diretório onde o arquivo `escalas_termo.go` está localizado.
  - Execute o comando de compilação:
   ```sh
   go build escalas_termo.go`
   ```
#### 3. Execute o Programa:
  -  Abra o terminal e navegue até o diretório onde o arquivo `escalas_termo.go` está localizado.
  - Execute o comando de compilação:
   ```sh
   go run escalas_termo.go`
   ou 
   ./escalas_termo 
   ```

#### 4. Entrada e Saída:

- Insira uma temperatura em Kelvin quando solicitado.
- O programa exibirá a temperatura correspondente em Celsius.

#### Contribuições

Se você deseja contribuir com melhorias ou correções, sinta-se à vontade para criar um pull request ou abrir uma issue no repositório.
   
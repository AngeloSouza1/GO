package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Estruturas para mapear a resposta da API
type PokedexResponse struct {
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
	Descriptions   []Description  `json:"descriptions"`
}

type PokemonEntry struct {
	EntryNumber    int            `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Description struct {
	Description string   `json:"description"`
	Language    Language `json:"language"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Estrutura para mapear os detalhes de um Pokémon
type PokemonDetail struct {
	Weight    int        `json:"weight"`
	Height    int        `json:"height"`
	Abilities []Ability  `json:"abilities"`
	Types     []TypeSlot `json:"types"`
}

type Ability struct {
	Ability AbilityInfo `json:"ability"`
}

type AbilityInfo struct {
	Name string `json:"name"`
}

type TypeSlot struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}

// Estrutura para os dados que serão passados ao template
type PokemonWithImage struct {
	EntryNumber int
	Name        string
	URL         string
	ImageURL    string
	Weight      float64
	Height      float64
	Abilities   string
	Type        string
}

type PageData struct {
	Description string
	Language    string
	LanguageURL string
	PokemonList []PokemonWithImage
}

func main() {
	// Fazer a requisição HTTP para a Pokédex de Kanto
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Fatalf("Erro ao fazer a requisição: %v", err)
	}
	defer response.Body.Close()

	// Ler o corpo da resposta
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Erro ao ler a resposta: %v", err)
	}

	// Decodificar o JSON na estrutura PokedexResponse
	var pokedex PokedexResponse
	err = json.Unmarshal(responseData, &pokedex)
	if err != nil {
		log.Fatalf("Erro ao decodificar o JSON: %v", err)
	}

	// Preparar os dados para o template HTML
	var pageData PageData
	for _, desc := range pokedex.Descriptions {
		if desc.Language.Name == "fr" {
			pageData.Description = desc.Description
			pageData.Language = desc.Language.Name
			pageData.LanguageURL = desc.Language.URL
			break
		}
	}

	for _, entry := range pokedex.PokemonEntries {
		// Obter o URL da imagem usando o ID do Pokémon
		imageURL := fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%d.png", entry.EntryNumber)

		// Fazer uma requisição adicional para obter detalhes do Pokémon
		detailResponse, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", entry.EntryNumber))
		if err != nil {
			log.Printf("Erro ao obter detalhes do Pokémon: %v", err)
			continue
		}
		defer detailResponse.Body.Close()

		var pokemonDetail PokemonDetail
		detailData, err := ioutil.ReadAll(detailResponse.Body)
		if err != nil {
			log.Printf("Erro ao ler detalhes do Pokémon: %v", err)
			continue
		}

		err = json.Unmarshal(detailData, &pokemonDetail)
		if err != nil {
			log.Printf("Erro ao decodificar detalhes do Pokémon: %v", err)
			continue
		}

		// Construir as strings para habilidades e tipos
		abilities := ""
		for i, ability := range pokemonDetail.Abilities {
			if i > 0 {
				abilities += ", "
			}
			abilities += ability.Ability.Name
		}

		types := ""
		for i, pokeType := range pokemonDetail.Types {
			if i > 0 {
				types += ", "
			}
			types += pokeType.Type.Name
		}

		// Preencher os dados do Pokémon com os detalhes adicionais
		pokemon := PokemonWithImage{
			EntryNumber: entry.EntryNumber,
			Name:        entry.PokemonSpecies.Name,
			URL:         entry.PokemonSpecies.URL,
			ImageURL:    imageURL,
			Weight:      float64(pokemonDetail.Weight) / 10.0, // Convertendo hectogramas para quilogramas
			Height:      float64(pokemonDetail.Height) / 10.0, // Convertendo decímetros para metros
			Abilities:   abilities,
			Type:        types,
		}
		pageData.PokemonList = append(pageData.PokemonList, pokemon)
	}

	// Gerar o arquivo HTML
	err = generateHTML(pageData)
	if err != nil {
		log.Fatalf("Erro ao gerar o arquivo HTML: %v", err)
	}

	fmt.Println("Arquivo HTML gerado com sucesso: pokedex.html")
}

func generateHTML(data PageData) error {
	// Ler o template HTML do arquivo
	tmplContent, err := ioutil.ReadFile("template.html")
	if err != nil {
		return fmt.Errorf("erro ao ler o template HTML: %v", err)
	}

	// Criar o template a partir do conteúdo lido
	tmpl, err := template.New("pokedex").Parse(string(tmplContent))
	if err != nil {
		return fmt.Errorf("erro ao analisar o template HTML: %v", err)
	}

	// Criar o arquivo HTML de saída
	file, err := os.Create("pokedex.html")
	if err != nil {
		return fmt.Errorf("erro ao criar o arquivo HTML de saída: %v", err)
	}
	defer file.Close()

	// Executar o template com os dados fornecidos e escrever no arquivo
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("erro ao executar o template: %v", err)
	}

	return nil
}

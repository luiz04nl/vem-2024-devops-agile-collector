package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/luiz04nl/devops-ic-collector/src/repository-dataset/shared"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// var repositories []shared.RepositoryDto
	var repositories *shared.GitHubGraphQLRepositoriesResponseDto
	var err error

	repositories, err = shared.GetRepositories()
	if err != nil {
		fmt.Println(err)
		return
	}

	repositoriesJsonData, err := json.Marshal(&repositories)
	if err != nil {
		log.Fatalf("Erro ao converter para JSON: %v", err)
	}
	log.Println("repositoriesJsonData: ", string(repositoriesJsonData))

	//GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto

}

package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/luiz04nl/devops-ic-collector/src/repository-dataset/shared"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var repositories []shared.RepositoryDto
	var err error

	repositories, err = shared.GetRepositories()
	if err != nil {
		fmt.Println(err)
		return
	}

	// repositoriesJsonData, err := json.Marshal(&repositories)
	// if err != nil {
	// 	log.Fatalf("Erro ao converter para JSON: %v", err)
	// }

	// log.Println("repositoriesJsonData: ", string(repositoriesJsonData))

	var dataSourceName = "../../database/sqlite/repository-dataset.db"

	createTableSQL := `
  CREATE TABLE IF NOT EXISTS repositories (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      url TEXT UNIQUE NOT NULL,
      name TEXT NOT NULL,
      starsTotalCount INTEGER,

      useAgile INTEGER DEFAULT NULL,
      useDevOps INTEGER DEFAULT NULL,
      qualityRanking INTEGER DEFAULT NULL
  );
  `

	db, err := sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela repositories: %v", err)
	}

	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados:", err)
	}

	err = newSQLiteRepository.SaveMultiple(repositories)
	if err != nil {
		log.Fatal("Erro ao salvar o repositório:", err)
	}
}

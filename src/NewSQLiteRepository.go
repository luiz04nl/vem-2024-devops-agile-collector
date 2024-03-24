package main

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db *sqlx.DB
}

func NewSQLiteRepository(dataSourceName string) (*SQLiteRepository, error) {
	db, err := sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &SQLiteRepository{db: db}, nil
}

func (sQLiteRepository *SQLiteRepository) Save(repo *Repository) error {
	query := `INSERT INTO repositories (name, url, starsTotalCount) VALUES (:name, :url, :starsTotalCount)`
	_, err := sQLiteRepository.db.NamedExec(query, repo)
	return err
}

// func (sQLiteRepository *SQLiteRepository) SaveMultiple(repos []Repository) error {
// 	// Iniciando uma transação
// 	tx, err := r.db.Beginx()
// 	if err != nil {
// 		return err
// 	}

// 	// Preparando a consulta SQL para inserir repositórios
// 	query := `INSERT INTO repositories (name, description) VALUES (:name, :description)`

// 	// Iterando sobre o slice de repositórios e inserindo um a um
// 	for _, repo := range repos {
// 		_, err := tx.NamedExec(query, &repo)
// 		if err != nil {
// 			// Se ocorrer um erro, fazemos um rollback da transação e retornamos o erro
// 			tx.Rollback() // Ignorando erro de rollback aqui para simplicidade
// 			return err
// 		}
// 	}

// 	// Se tudo correr bem, commitamos a transação
// 	return tx.Commit()
// }

func (sQLiteRepository *SQLiteRepository) SaveMultiple(repos []Repository) error {
	// Iniciando uma transação
	tx, err := sQLiteRepository.db.Beginx()
	if err != nil {
		return err
	}

	// Construindo a instrução SQL para inserção em massa
	valueStrings := make([]string, 0, len(repos))
	valueArgs := make([]interface{}, 0, len(repos)*2) // 2 campos por repositório
	for _, repo := range repos {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, repo.Name, repo.URL, repo.StarsTotalCount)
	}
	stmt := fmt.Sprintf("INSERT INTO repositories (name, url, starsTotalCount) VALUES %s",
		strings.Join(valueStrings, ","))

	// Executando a inserção em massa
	_, err = tx.Exec(stmt, valueArgs...)
	if err != nil {
		tx.Rollback() // Ignorando erro de rollback aqui para simplicidade
		return err
	}

	// Se tudo correr bem, commitamos a transação
	return tx.Commit()
}

func (sQLiteRepository *SQLiteRepository) FindByID(id int64) (*Repository, error) {
	var repo Repository
	err := sQLiteRepository.db.Get(&repo, "SELECT * FROM repositories WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (sQLiteRepository *SQLiteRepository) FindByName(name string) (*Repository, error) {
	var repo Repository
	err := sQLiteRepository.db.Get(&repo, "SELECT * FROM repositories WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

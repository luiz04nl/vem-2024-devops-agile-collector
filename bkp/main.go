package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// // Historico de commits e pull requets mais recentos em repositorio
	// jsonMapInstance := map[string]string{
	// 	"query": `
	//   {
	//     repository(owner: "luiz04nl", name: "devops-ic-collector") {
	//       defaultBranchRef {
	//         target {
	//           ... on Commit {
	//             history(first: 10) {
	//               edges {
	//                 node {
	//                   committedDate
	//                   oid
	//                   messageHeadline
	//                   author {
	//                     name
	//                     date
	//                   }
	//                 }
	//               }
	//             }
	//           }
	//         }
	//       }
	//       pullRequests(last: 10, orderBy: {field: CREATED_AT, direction: DESC}) {
	//         edges {
	//           node {
	//             title
	//             state
	//             author {
	//               login
	//             }
	//             createdAt
	//           }
	//         }
	//       }
	//     }
	//   }
	// `,
	// }

	// Historico de commits e pull requets mais recentos em repositorio
	query := `
{
  search(query: "is:public stars:>=100", type: REPOSITORY, first: 10) {
    repositoryCount
    pageInfo {
      endCursor
      startCursor
    }
    edges {
      node {
        ... on Repository {
          name
          url
          stargazers {
            totalCount
          }
          collaborators(first: 3) {
            totalCount
            edges {
              permission
              node {
                id
                name
                email
              }
            }
            pageInfo {
              hasNextPage
              endCursor
            }
          }
        }
      }
    }
  }
}
`
	GitHubGraphQLRepositoriesResponseDto, err := executeGraphQLQuery(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	var repositories = GitHubGraphQLRepositoriesResponseDtoToRepositoriesDto(GitHubGraphQLRepositoriesResponseDto)

	// Criando a tabela repositories
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS repositories (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            url TEXT NOT NULL,
            starsTotalCount INTEGER
        );
        `

	var dataSourceName1 = "../database/sqlite/extraction.db"
	db, err := sqlx.Connect("sqlite3", dataSourceName1)
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Erro ao criar tabela repositories: %v", err)
	}

	// Inicialização do repositório
	var dataSourceName2 = "../database/sqlite/extraction.db"
	newSQLiteRepository, err := NewSQLiteRepository(dataSourceName2)
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco de dados:", err)
	}

	err = newSQLiteRepository.Save(&repositories[0])
	if err != nil {
		log.Fatal("Erro ao salvar o repositório:", err)
	}

	// jsonData, err := json.Marshal(repository)
	// if err != nil {
	// 	log.Fatalf("Erro ao converter para JSON: %v", err)
	// }

	// fmt.Println(string(jsonData))

	// URL do repositório Git que você deseja clonar
	// repoURL := repositories[0].URL
	repoURL := "https://github.com/luiz04nl/devops-ic-collector.git"

	// Diretório onde o repositório será clonado
	// Pode ser um caminho absoluto ou relativo
	dir := "../repos/repo1"

	// Clonando o repositório
	cmd := exec.Command("git", "clone", repoURL, dir)
	if err := cmd.Run(); err != nil {
		// log.Fatalf("Erro ao clonar o repositório: %v", err)

		// // cmd2 := exec.Command("cd")
		// // cmd2.Dir = "../repos/repo1"
		// err := os.Chdir("repos/repo1")

		// // if err := cmd2.Run(); err != nil {
		// if err != nil {
		// 	log.Fatalf("Erro 2: %v", err)
		// }

		cmd3 := exec.Command("git", "pull")
		cmd3.Dir = "../repos/repo1"
		if err := cmd3.Run(); err != nil {
			log.Fatalf("Erro 3: %v", err)
		}

		// cat <<EOF > sonar-project.properties
		// sonar.projectKey=meu_projeto
		// sonar.projectName=Meu Projeto
		// sonar.projectVersion=1.0
		// sonar.sources=.
		// sonar.sourceEncoding=UTF-8
		// EOF

	}

	log.Println("Repositório clonado com sucesso.")

	//   CircleCI = .cicleci/config.yaml
	// https://circleci.com/docs/sample-config/
	// Github action = qualquer arquivo yaml ou yml dentro de .github/workflows
	// https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions
	// Jenkins =  Jenkinsfile na raiz do projeto
	// https://www.jenkins.io/doc/book/pipeline/getting-started/#:~:text=The%20default%20value%20of%20this,the%20root%20of%20the%20repository.
	// Azure Devops =  azure-pipelines.yml na raiz do projeto
	// https://learn.microsoft.com/en-us/azure/devops/pipelines/customize-pipeline?view=azure-devops#understand-the-azure-pipelinesyml-file

	// GitLab = .gitlab-ci.yml no diretorio root do projeto
	// https://docs.gitlab.com/ee/ci/pipelines/settings.html
	// BitBucket =  bitbucket-pipelines.yml no diretorio root do projeto
	// https://chrisfrewin.medium.com/the-last-bitbucket-pipelines-tutorial-youll-ever-need-mastering-ci-and-cd-28a027fc5e40
	// Travis CI =  .travis.yml na raiz do projeto
	// https://docs.travis-ci.com/user/tutorial/
	// Harness =  qualquer arquivo yaml ou yml dentro de .harness
	// https://www.harness.io/blog/ci-cd-pipeline-as-code-with-harness

	// // Mudando para o diretório do repositório clonado
	// cmd = exec.Command("bash", "-c", "cd "+dir+" && seu_comando_aqui")
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalf("Erro ao executar o comando no diretório do repositório: %v", err)
	// }

	// log.Println("Comando executado com sucesso no diretório do repositório.")

	// sonar-scanner

}

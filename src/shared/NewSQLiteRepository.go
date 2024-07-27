package shared

import (
	"sort"

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

func (sQLiteRepository *SQLiteRepository) Save(repo *RepositoryDto) error {
	query := `INSERT INTO repositories (name, url, alias, starsTotalCount) VALUES (:name, :url, :alias, :starsTotalCount)`
	_, err := sQLiteRepository.db.NamedExec(query, repo)
	return err
}

func (sQLiteRepository *SQLiteRepository) UpdateById(id int, newRepository RepositoryDto) error {
	repository, err := sQLiteRepository.FindByID(id)
	if err != nil {
		return err
	}

	statement := `UPDATE repositories
  SET name = ?, url = ?, alias = ?, starsTotalCount = ?,
  useAgile = ?, useDevOps = ?,
  useGithubPipelines = ?, useCircleCI = ?, useJenkins = ?,
  useGitLabPipelines = ?, useAzureDevops = ?, useTravisCI = ?,
  useHarness = ?, useBitBucketPipelines = ?, wasCloned = ?, meta = ?,
  projectContributors = ?, projectCommits = ?, commitsIntervalInDays = ?, contributorsInfo = ?,
  hasCommitsInInterval = ?, lastCommitDateInterval = ?,
  projectType = ?, projectTypeVersion = ?, projectIssuesEffortTotal = ?,
  projectIssuesCount = ?, projectCodeSmellsEffortTotal = ?, projectCodeSmellsCount = ?,
  projectSonarComponentsCount = ?, projectSonarInfo = ?,
  bugs = ?, sqaleRating = ?, reliabilityRating = ?,
  complexity = ?, cognitiveComplexity = ?, duplicatedBlocks = ?,
  duplicatedFiles = ?, duplicatedLines = ?, codeSmells = ?, linesOfCodesFromSonar = ?,
  sqaleIndex = ?, sqaleDebtRatio = ?, qualityGateDetails = ?,
  vulnerabilities = ?, securityRating = ?, classes = ?,
  commentLines = ?, coverage = ?, tests = ?,
  linesOfCodesFromCk = ?, couplingBetweenObjects = ?, couplingBetweenObjectsModified = ?
  WHERE id = ?
  `

	_, err = sQLiteRepository.db.Exec(statement,
		newRepository.Name, newRepository.URL, newRepository.Alias, newRepository.StarsTotalCount,
		newRepository.UseAgile, newRepository.UseDevOps,
		newRepository.UseGithubPipelines, newRepository.UseCircleCI, newRepository.UseJenkins,
		newRepository.UseGitLabPipelines, newRepository.UseAzureDevops, newRepository.UseTravisCI,
		newRepository.UseHarness, newRepository.UseBitBucketPipelines,
		newRepository.WasCloned, newRepository.Meta,
		newRepository.ProjectContributors, newRepository.ProjectCommits,
		newRepository.CommitsIntervalInDays, newRepository.ContributorsInfo,
		newRepository.HasCommitsInInterval, newRepository.LastCommitDateInterval,
		newRepository.ProjectType, newRepository.ProjectTypeVersion, newRepository.ProjectIssuesEffortTotal,
		newRepository.ProjectIssuesCount, newRepository.ProjectCodeSmellsEffortTotal, newRepository.ProjectCodeSmellsCount,
		newRepository.ProjectSonarComponentsCount, newRepository.ProjectSonarInfo,
		newRepository.Bugs, newRepository.SqaleRating, newRepository.ReliabilityRating,
		newRepository.Complexity, newRepository.CognitiveComplexity, newRepository.DuplicatedBlocks,
		newRepository.DuplicatedFiles, newRepository.DuplicatedLines, newRepository.CodeSmells, newRepository.LinesOfCodesFromSonar,
		newRepository.SqaleIndex, newRepository.SqaleDebtRatio, newRepository.QualityGateDetails,
		newRepository.Vulnerabilities, newRepository.SecurityRating, newRepository.Classes,
		newRepository.CommentLines, newRepository.Coverage, newRepository.Tests,
		newRepository.LinesOfCodesFromCk, newRepository.CouplingBetweenObjects, newRepository.CouplingBetweenObjectsModified,
		repository.Id)
	if err != nil {
		return err
	}
	return nil
}

func (sQLiteRepository *SQLiteRepository) SaveMultiple(repos []RepositoryDto) error {

	// Iniciando uma transação
	tx, err := sQLiteRepository.db.Beginx()
	if err != nil {
		return err
	}

	// Preparando a consulta SQL para inserir repositórios
	query := `INSERT INTO repositories (name, url, alias, starsTotalCount) VALUES (:name, :url, :alias, :starsTotalCount)`

	// Iterando sobre o slice de repositórios e inserindo um a um
	for _, repo := range repos {
		_, err := tx.NamedExec(query, &repo)
		if err != nil {
			// Se ocorrer um erro, fazemos um rollback da transação e retornamos o erro
			tx.Rollback() // Ignorando erro de rollback aqui para simplicidade
			return err
		}
	}

	// Se tudo correr bem, commitamos a transação
	return tx.Commit()
}

func (sQLiteRepository *SQLiteRepository) GetAll() ([]RepositoryDto, error) {
	// rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories WHERE alias = 'greys-anatomy-813536180aa5d1bf1bf4ec642ecb04f7e021df6961b7ac79aa64efa3d3fa6621'")

	rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories")

	// rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories where id < 100")

	// rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories where wasCloned = 1 and projectIssuesCount > 0")

	// rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories where id < 13 and wasCloned = 1 and projectIssuesCount > 0")

	// rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories where alias = 'X'")

	// rows, err := sQLiteRepository.db.Query("SELECT * FROM repositories where id < 2")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repositories []RepositoryDto

	for rows.Next() {
		var repo RepositoryDto
		if err := rows.Scan(
			&repo.Id,
			&repo.URL,
			&repo.Name,
			&repo.Alias,
			&repo.StarsTotalCount,
			&repo.UseAgile,
			&repo.UseDevOps,
			&repo.UseGithubPipelines,
			&repo.UseCircleCI,
			&repo.UseJenkins,
			&repo.UseGitLabPipelines,
			&repo.UseAzureDevops,
			&repo.UseTravisCI,
			&repo.UseHarness,
			&repo.UseBitBucketPipelines,
			&repo.WasCloned,
			&repo.Meta,
			&repo.ProjectContributors,
			&repo.ProjectCommits,
			&repo.CommitsIntervalInDays,
			&repo.ContributorsInfo,
			&repo.HasCommitsInInterval,
			&repo.LastCommitDateInterval,
			&repo.ProjectType,
			&repo.ProjectTypeVersion,
			&repo.ProjectIssuesEffortTotal,
			&repo.ProjectIssuesCount,
			&repo.ProjectCodeSmellsEffortTotal,
			&repo.ProjectCodeSmellsCount,
			&repo.ProjectSonarComponentsCount,
			&repo.ProjectSonarInfo,
			&repo.Bugs,
			&repo.SqaleRating,
			&repo.ReliabilityRating,
			&repo.Complexity,
			&repo.CognitiveComplexity,
			&repo.DuplicatedBlocks,
			&repo.DuplicatedFiles,
			&repo.DuplicatedLines,
			&repo.CodeSmells,
			&repo.LinesOfCodesFromSonar,
			&repo.SqaleIndex,
			&repo.SqaleDebtRatio,
			&repo.QualityGateDetails,
			&repo.Vulnerabilities,
			&repo.SecurityRating,
			&repo.Classes,
			&repo.CommentLines,
			&repo.Coverage,
			&repo.Tests,

			&repo.LinesOfCodesFromCk,
			&repo.CouplingBetweenObjects,
			&repo.CouplingBetweenObjectsModified,
		); err != nil {
			return nil, err
		}
		repositories = append(repositories, repo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return repositories, nil
}

func (sQLiteRepository *SQLiteRepository) FindByID(id int) (*RepositoryDto, error) {
	var repo RepositoryDto
	err := sQLiteRepository.db.Get(&repo, "SELECT * FROM repositories WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (sQLiteRepository *SQLiteRepository) FindByName(name string) (*RepositoryDto, error) {
	var repo RepositoryDto
	err := sQLiteRepository.db.Get(&repo, "SELECT * FROM repositories WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

// ################

type AggregateDevOpsToolsUsageDto struct {
	Name  string
	Value float64
}

type AggregateDevOpsToolsUsageDatabaseReponse struct {
	GithubPipelines    float64
	CircleCI           float64
	Jenkins            float64
	GitLabPipelines    float64
	AzureDevops        float64
	TravisCI           float64
	Harness            float64
	BitBucketPipelines float64
}

func (sQLiteRepository *SQLiteRepository) AggregateDevOpsToolsUsage() ([]AggregateDevOpsToolsUsageDto, error) {

	rows, err := sQLiteRepository.db.Query(`
  SELECT
    SUM(CASE WHEN useGithubPipelines = 1 THEN 1 ELSE 0 END) AS GithubPipelines,
    SUM(CASE WHEN useCircleCI = 1 THEN 1 ELSE 0 END) AS CircleCI,
    SUM(CASE WHEN useJenkins = 1 THEN 1 ELSE 0 END) AS Jenkins,
    SUM(CASE WHEN useGitLabPipelines = 1 THEN 1 ELSE 0 END) AS GitLabPipelines,
    SUM(CASE WHEN useAzureDevops = 1 THEN 1 ELSE 0 END) AS AzureDevops,
    SUM(CASE WHEN useTravisCI = 1 THEN 1 ELSE 0 END) AS TravisCI,
    SUM(CASE WHEN useHarness = 1 THEN 1 ELSE 0 END) AS Harness,
    SUM(CASE WHEN useBitBucketPipelines = 1 THEN 1 ELSE 0 END) AS BitBucketPipelines
FROM
    repositories
WHERE
    wasCloned = 1
    AND useDevOps = 1;
  `)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []AggregateDevOpsToolsUsageDatabaseReponse
	for rows.Next() {
		var dto AggregateDevOpsToolsUsageDatabaseReponse
		if err := rows.Scan(
			&dto.GithubPipelines,
			&dto.CircleCI,
			&dto.Jenkins,
			&dto.GitLabPipelines,
			&dto.AzureDevops,
			&dto.TravisCI,
			&dto.Harness,
			&dto.BitBucketPipelines,
		); err != nil {
			return nil, err
		}
		responses = append(responses, dto)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var dtos []AggregateDevOpsToolsUsageDto
	for _, response := range responses {
		dtos = append(dtos,
			AggregateDevOpsToolsUsageDto{Name: "GithubPipelines", Value: response.GithubPipelines},
			AggregateDevOpsToolsUsageDto{Name: "CircleCI", Value: response.CircleCI},
			AggregateDevOpsToolsUsageDto{Name: "Jenkins", Value: response.Jenkins},
			AggregateDevOpsToolsUsageDto{Name: "GitLabPipelines", Value: response.GitLabPipelines},
			AggregateDevOpsToolsUsageDto{Name: "AzureDevops", Value: response.AzureDevops},
			AggregateDevOpsToolsUsageDto{Name: "TravisCI", Value: response.TravisCI},
			AggregateDevOpsToolsUsageDto{Name: "Harness", Value: response.Harness},
			AggregateDevOpsToolsUsageDto{Name: "BitBucketPipelines", Value: response.BitBucketPipelines},
		)
	}

	sort.Slice(dtos, func(i, j int) bool {
		return dtos[i].Value < dtos[j].Value
	})

	return dtos, nil
}

// SELECT
//     SUM(CASE WHEN useGithubPipelines = 1 THEN 1 ELSE 0 END) AS GithubPipelines,
//     SUM(CASE WHEN useCircleCI = 1 THEN 1 ELSE 0 END) AS CircleCI,
//     SUM(CASE WHEN useJenkins = 1 THEN 1 ELSE 0 END) AS Jenkins,
//     SUM(CASE WHEN useGitLabPipelines = 1 THEN 1 ELSE 0 END) AS GitLabPipelines,
//     SUM(CASE WHEN useAzureDevops = 1 THEN 1 ELSE 0 END) AS AzureDevops,
//     SUM(CASE WHEN useTravisCI = 1 THEN 1 ELSE 0 END) AS TravisCI,
//     SUM(CASE WHEN useHarness = 1 THEN 1 ELSE 0 END) AS Harness,
//     SUM(CASE WHEN useBitBucketPipelines = 1 THEN 1 ELSE 0 END) AS BitBucketPipelines
// FROM
//     repositories
// WHERE
//     wasCloned = 1
//     AND useDevOps = 1;

// SELECT
//     SUM(CASE WHEN useAgile = 1 THEN 1 ELSE 0 END) AS Agile,
//     SUM(CASE WHEN useDevops = 1 THEN 1 ELSE 0 END) AS DevOps,
//     SUM(CASE WHEN useDevops = 1 AND useAgile = 1 THEN 1 ELSE 0 END) AS AgileAndDevOps,
//     SUM(CASE WHEN useDevops = 0 AND useAgile = 0 THEN 1 ELSE 0 END) AS NeitherAgileAndDevOpsm,
//     SUM(CASE WHEN useDevops = 1 AND useAgile = 0 THEN 1 ELSE 0 END) AS DevopsButNonAgile,
//     SUM(CASE WHEN useAgile = 1 AND useDevops = 0 THEN 1 ELSE 0 END) AS AgileButNonDovOps
// FROM
//     repositories
// WHERE
//     wasCloned = 1;

// SELECT
//     projectType AS Name,
//     COUNT(id) AS Value
// FROM
//     repositories
// where projectType in ('maven', 'gradle', 'ant', 'other', 'nodejs', '')
// and wasCloned = 1
// GROUP BY
//     projectType

// select
// 	useAgile,
// 	count(useAgile) as useAgileCount,
// 	CASE
//     	WHEN CAST(sqaleRating AS REAL) BETWEEN 0 AND 0.05 THEN 'A'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.06 AND 0.1 THEN 'B'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.11 AND 0.2 THEN 'C'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.21 AND 0.5 THEN 'D'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.51 AND 1 THEN 'E'
//         ELSE null
//     END AS maintainabilityRating,
//     AVG(sqaleDebtRatio) AS technicalDebtRatio,
//     AVG(codeSmells) AS codeSmellsCount
// from repositories
// where projectType in ('maven', 'gradle', 'ant')
// and wasCloned = 1 and projectSonarComponentsCount > 0
// group by useAgile

// select
// 	useDevOps,
// 	count(useDevOps) as useDevOpsCount,
// 	CASE
//     	WHEN CAST(sqaleRating AS REAL) BETWEEN 0 AND 0.05 THEN 'A'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.06 AND 0.1 THEN 'B'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.11 AND 0.2 THEN 'C'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.21 AND 0.5 THEN 'D'
//         WHEN CAST(sqaleRating AS REAL) BETWEEN 0.51 AND 1 THEN 'E'
//         ELSE null
//     END AS maintainabilityRating,
//     AVG(sqaleDebtRatio) AS technicalDebtRatio,
//     AVG(codeSmells) AS codeSmellsCount
// from repositories
// where projectType in ('maven', 'gradle', 'ant')
// and wasCloned = 1 and projectSonarComponentsCount > 0
// group by useDevOps

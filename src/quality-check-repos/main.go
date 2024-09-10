package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
	_ "github.com/mattn/go-sqlite3"
)

type SonarMeasuresMapValue struct {
	Name          string
	BuilderAction func(shared.RepositoryDtoBuilder, string) shared.RepositoryDtoBuilder
}

func UpdateMeasures(repository shared.RepositoryDto) shared.RepositoryDto {
	repositoryDtoBuilder := shared.RepositoryDtoBuilder{}.FromRepository(repository)

	sonarMeasuresMap := make(map[string]SonarMeasuresMapValue)
	sonarMeasuresMap["bugs"] = SonarMeasuresMapValue{
		Name: "bugs", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithBugs(value)
		},
	}
	sonarMeasuresMap["sqale_rating"] = SonarMeasuresMapValue{
		Name: "sqale_rating", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithSqaleRating(value)
		},
	}
	sonarMeasuresMap["reliability_rating"] = SonarMeasuresMapValue{
		Name: "reliability_rating", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithReliabilityRating(value)
		},
	}
	sonarMeasuresMap["complexity"] = SonarMeasuresMapValue{
		Name: "complexity", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithComplexity(value)
		},
	}
	sonarMeasuresMap["cognitive_complexity"] = SonarMeasuresMapValue{
		Name: "cognitive_complexity", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithCognitiveComplexity(value)
		},
	}
	sonarMeasuresMap["duplicated_blocks"] = SonarMeasuresMapValue{
		Name: "duplicated_blocks", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithDuplicatedBlocks(value)
		},
	}
	sonarMeasuresMap["duplicated_files"] = SonarMeasuresMapValue{
		Name: "duplicated_files", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithDuplicatedFiles(value)
		},
	}
	sonarMeasuresMap["duplicated_lines"] = SonarMeasuresMapValue{
		Name: "duplicated_lines", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithDuplicatedLines(value)
		},
	}
	sonarMeasuresMap["code_smells"] = SonarMeasuresMapValue{
		Name: "code_smells", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithCodeSmells(value)
		},
	}
	sonarMeasuresMap["ncloc"] = SonarMeasuresMapValue{
		Name: "ncloc", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithLinesOfCodesFromSonar(value)
		},
	}
	// #ncloc_language_distribution
	sonarMeasuresMap["sqale_index"] = SonarMeasuresMapValue{
		Name: "sqale_index", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithSqaleIndex(value)
		},
	}
	sonarMeasuresMap["sqale_debt_ratio"] = SonarMeasuresMapValue{
		Name: "sqale_debt_ratio", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithSqaleDebtRatio(value)
		},
	}
	sonarMeasuresMap["quality_gate_details"] = SonarMeasuresMapValue{
		Name: "quality_gate_details", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithQualityGateDetails(value)
		},
	}
	sonarMeasuresMap["vulnerabilities"] = SonarMeasuresMapValue{
		Name: "vulnerabilities", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithVulnerabilities(value)
		},
	}
	sonarMeasuresMap["security_rating"] = SonarMeasuresMapValue{
		Name: "security_rating", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithSecurityRating(value)
		},
	}

	sonarMeasuresMap["classes"] = SonarMeasuresMapValue{
		Name: "classes", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithClasses(value)
		},
	}

	sonarMeasuresMap["comment_lines"] = SonarMeasuresMapValue{
		Name: "comment_lines", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithCommentLines(value)
		},
	}

	sonarMeasuresMap["coverage"] = SonarMeasuresMapValue{
		Name: "coverage", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithCoverage(value)
		},
	}

	sonarMeasuresMap["tests"] = SonarMeasuresMapValue{
		Name: "tests", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder, value string) shared.RepositoryDtoBuilder {
			return repositoryDtoBuilder.WithTests(value)
		},
	}

	for _, sonarMeasuresMapValue := range sonarMeasuresMap {
		outputFile := fmt.Sprintf("%s/%s-%s.json", "../../out/quality-check-repos", repository.Alias, sonarMeasuresMapValue.Name)
		fmt.Printf("outputFile: %s", outputFile)

		jsonFile := outputFile
		file, err := os.Open(jsonFile)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return repository
		}
		defer file.Close()
		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return repository
		}
		var checkOutputDto shared.SonarMeasuresOutputDto
		err = json.Unmarshal(byteValue, &checkOutputDto)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return repository
		}

		if len(checkOutputDto.Component.Measures) > 0 {
			value := checkOutputDto.Component.Measures[0].Value
			repositoryDtoBuilder = sonarMeasuresMapValue.BuilderAction(repositoryDtoBuilder, value)
		}
	}

	newRepository := repositoryDtoBuilder.Build()

	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
	if err != nil {
		fmt.Println("Error ao atualizar repositório:", err)
	} else {
		// fmt.Println("Repositório atualizado:")
		//return newRepository
	}

	return repository
}

func UpdateDefaultChecks(repository shared.RepositoryDto) shared.RepositoryDto {
	defaultOutputFile := fmt.Sprintf("%s/%s.json", "../../out/quality-check-repos", repository.Alias)
	fmt.Printf("defaultOutputFile: %s", defaultOutputFile)

	jsonFile := defaultOutputFile
	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return repository
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return repository
	}
	var checkOutputDto shared.DefaultCheckOutputDto
	err = json.Unmarshal(byteValue, &checkOutputDto)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return repository
	}

	builder := shared.RepositoryDtoBuilder{}.
		FromRepository(repository).
		WithProjectType(checkOutputDto.ProjectType).
		WithProjectTypeVersion(checkOutputDto.ProjectTypeVersion)

	newRepository := builder.Build()

	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
	if err != nil {
		fmt.Println("Error ao atualizar repositório:", err)
	} else {
		// fmt.Println("Repositório atualizado:")
		return newRepository
	}

	return repository
}

// ###############
func ReadIssuesCheckOutputDtoFromFile(issuesOutputFile string) (*shared.IssuesCheckOutputDto, error) {
	jsonFile := issuesOutputFile
	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var checkOutputDto shared.IssuesCheckOutputDto
	err = json.Unmarshal(byteValue, &checkOutputDto)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	return &checkOutputDto, nil
}

func UpdateIssuesChecks(repository shared.RepositoryDto) shared.RepositoryDto {
	firstIssuesOutputFile := fmt.Sprintf("%s/%s-ISSUES-page-1.json", "../../out/quality-check-repos", repository.Alias)

	firstCheckOutputDto, err := ReadIssuesCheckOutputDtoFromFile(firstIssuesOutputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return repository
	}
	issuesCount := firstCheckOutputDto.Total

	pageCheckOutputDtos := []any{}
	pageCheckOutputDtos = append(pageCheckOutputDtos, firstCheckOutputDto)
  pageSize := 500
  pages := (issuesCount + pageSize - 1) / pageSize
	for page := 2; page <= pages; page++ {
		issuesOutputFile := fmt.Sprintf("%s/%s-ISSUES-page-%d.json", "../../out/quality-check-repos", repository.Alias, page)
		pageCheckOutputDto, err := ReadIssuesCheckOutputDtoFromFile(issuesOutputFile)
		if err != nil {
			fmt.Println("Error on load json file:", err)
		} else {
			pageCheckOutputDtos = append(pageCheckOutputDtos, pageCheckOutputDto)
		}
	}
	// projectSonarInfoJsonData, err := json.Marshal(pageCheckOutputDtos)
	// if err != nil {
	// 	fmt.Println("Error converting to JSON:", err)
	// 	return repository
	// }

	builder := shared.RepositoryDtoBuilder{}.
		FromRepository(repository).
		WithProjectIssuesEffortTotal(firstCheckOutputDto.EffortTotal).
		// WithProjectSonarInfo(string(projectSonarInfoJsonData)).
		WithProjectSonarInfo("").
    WithProjectIssuesCount(issuesCount).
		WithProjectSonarComponentsCount(len(firstCheckOutputDto.Components))

	newRepository := builder.Build()

	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
	if err != nil {
		fmt.Println("Error ao atualizar repositório:", err)
	} else {
		// fmt.Println("Repositório atualizado:")
		return newRepository
	}

	return repository
}

// ###############

// ###############
func ReadICodeSmellsCheckOutputDtoFromFile(codeSmellsOutputFile string) (*shared.IssuesCheckOutputDto, error) {
	jsonFile := codeSmellsOutputFile
	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	var checkOutputDto shared.CodeSmellsCheckOutputDto
	err = json.Unmarshal(byteValue, &checkOutputDto)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	return &checkOutputDto, nil
}

func UpdateCodeSmellsChecks(repository shared.RepositoryDto) shared.RepositoryDto {
	firstCodeSmellsOutputFile := fmt.Sprintf("%s/%s-CODE_SMELL-page-1.json", "../../out/quality-check-repos", repository.Alias)
	firstCheckOutputDto, err := ReadICodeSmellsCheckOutputDtoFromFile(firstCodeSmellsOutputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return repository
	}

	codeSmellsCount := firstCheckOutputDto.Total
	builder := shared.RepositoryDtoBuilder{}.
		FromRepository(repository).
		WithProjectCodeSmellsEffortTotal(firstCheckOutputDto.EffortTotal).
		WithProjectCodeSmellsCount(codeSmellsCount)

	newRepository := builder.Build()

	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Println("Was not possible connect with database:", err)
	}

	err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
	if err != nil {
		fmt.Println("Error ao atualizar repositório:", err)
	} else {
		// fmt.Println("Repositório atualizado:")
		return newRepository
	}

	return repository
}

// ###############

func CheckAndUpdate(repositories []shared.RepositoryDto) {
	for index, repository := range repositories {
		fmt.Printf("Index: %d\n", index)

		now := time.Now()
		fmt.Println("Started at:", now.Format("02/01/2006 3:04:05 PM"))

		shellScriptOutputRepoFile := fmt.Sprintf("%s/%s.out.txt", "../../out/quality-check-repos", repository.Alias)
		fmt.Printf("shellScriptOutputRepoFile: %s", shellScriptOutputRepoFile)

		// withBuild := "false"
		withBuild := "true"

		cmdString := fmt.Sprintf("./main.sh %s %s > %s", repository.Alias, withBuild, shellScriptOutputRepoFile)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
		defer cancel()
		cmd := exec.CommandContext(ctx, "sh", "-c", cmdString)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				log.Println("Comando expirou: %v", ctx.Err())
			} else {
				log.Println("Erro ao executar o comando: %v", err)
			}
		} else {
			repository1 := UpdateDefaultChecks(repository)
			repository2 := UpdateIssuesChecks(repository1)
			UpdateMeasures(repository2)
		}
		fmt.Println("\n", output)

		now2 := time.Now()
		fmt.Println("\nEnd at:", now2.Format("02/01/2006 3:04:05 PM"))
	}

	fmt.Println("######## finished run-quality-check ########")
}

func main() {
	var dataSourceName = "../../database/sqlite/repository-dataset.db"

	// alterTableSQL := `
	//   ALTER TABLE repositories ADD XX TEXT DEFAULT "";
	// `
	// db, err := sqlx.Connect("sqlite3", dataSourceName)
	// if err != nil {
	// 	log.Println("Connection error: %v", err)
	// }
	// _, err = db.Exec(alterTableSQL)
	// if err != nil {
	// 	log.Println("Error on alterar tabela repositories: %v", err)
	// }

	// #########################################

	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	var repositories []shared.RepositoryDto

	repositories, err = newSQLiteRepository.GetCloned()
	if err != nil {
		log.Fatal("Error on obter o repositório:", err)
	}

	CheckAndUpdate(repositories)
}

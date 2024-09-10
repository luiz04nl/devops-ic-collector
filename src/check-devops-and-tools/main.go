package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
	_ "github.com/mattn/go-sqlite3"
)

type DevopsMapValue struct {
	FilePath      string
	BuilderAction func(shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder
}

func CheckAndUpdateDevopsUse(repositories []shared.RepositoryDto) {
	for index, repository := range repositories {
		fmt.Printf("Index: %d\n", index)

		dir := fmt.Sprintf("../../repos/%s", repository.Alias)
		fmt.Printf("dir: %s\n", dir)

		var dataSourceName = "../../database/sqlite/repository-dataset.db"
		newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
		if err != nil {
			log.Fatal("Was not possible connect with database:", err)
		}

		repositoryDtoBuilder := shared.RepositoryDtoBuilder{}.FromRepository(repository)

		devopsMap := make(map[string]DevopsMapValue)
		devopsMap["githubPipelines"] = DevopsMapValue{
			FilePath: ".github/workflows", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithUseGithubPipelines()
			},
		}
		devopsMap["circleCI"] = DevopsMapValue{
			FilePath: ".cicleci/config.yaml", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithUseCircleCI()
			},
		}
		devopsMap["jenkins"] = DevopsMapValue{
			FilePath: "Jenkinsfile", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithJenkins()
			},
		}
		devopsMap["gitLabPipelines"] = DevopsMapValue{
			FilePath: ".gitlab-ci.yml", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithGitLabPipelines()
			},
		}
		devopsMap["azureDevops"] = DevopsMapValue{
			FilePath: "azure-pipelines.yml", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithAzureDevops()
			},
		}
		devopsMap["travisCI"] = DevopsMapValue{
			FilePath: ".travis.yml", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithTravisCI()
			},
		}
		devopsMap["harness"] = DevopsMapValue{
			FilePath: ".harness", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithHarness()
			},
		}
		devopsMap["bitbucketPipelines"] = DevopsMapValue{
			FilePath: "bitbucket-pipelines.yml", BuilderAction: func(repositoryDtoBuilder shared.RepositoryDtoBuilder) shared.RepositoryDtoBuilder {
				return repositoryDtoBuilder.WithBitbucketPipelines()
			},
		}

		for _, devopsMapValue := range devopsMap {
			cmd := exec.Command("ls", devopsMapValue.FilePath)
			cmd.Dir = dir

			if err := cmd.Run(); err != nil {
				fmt.Printf("Sliped, was not founded: \"%s\"\n", err)
				fmt.Printf("Sliped, was not founded: \"%s\"\n", devopsMapValue.FilePath)

			} else {
				fmt.Printf("Ok, was founded: \"%s\"\n", devopsMapValue.FilePath)
				repositoryDtoBuilder = devopsMapValue.BuilderAction(repositoryDtoBuilder).WithUseDevops()
			}
		}

		newRepository := repositoryDtoBuilder.Build()
		err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
		if err != nil {
			fmt.Println("Error, was not possible update repository", err)
		}
	}

	fmt.Println("######## finished run-check-devops-and-tools ########")
}

func main() {
	// var dataSourceName = "../../database/sqlite/repository-dataset.db"
	// alterTableSQL := `
	//   ALTER TABLE repositories ADD XXX INTEGER DEFAULT 0;
	// `
	// db, err := sqlx.Connect("sqlite3", dataSourceName)
	// if err != nil {
	// 	log.Fatalf("Connection error: %v", err)
	// }
	// _, err = db.Exec(alterTableSQL)
	// if err != nil {
	// 	log.Fatalf("Error on alterar tabela repositories: %v", err)
	// }
	// #########################################

	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	var repositories []shared.RepositoryDto

	repositories, err = newSQLiteRepository.GetCloned()
	if err != nil {
		log.Fatal("Error on get the repository:", err)
	}

	CheckAndUpdateDevopsUse(repositories)
}

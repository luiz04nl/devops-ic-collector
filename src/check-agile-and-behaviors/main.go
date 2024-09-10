package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"os/exec"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
	_ "github.com/mattn/go-sqlite3"
)

func CheckAndUpdateAgileUse(repositories []shared.RepositoryDto) {
	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	for index, repository := range repositories {
		fmt.Printf("Index: %d\n", index)

		pythonOutputRepoFile := fmt.Sprintf("%s/%s.json", "../../out/check-agile-and-behaviors", repository.Alias)
		fmt.Printf("pythonOutputRepoFile: %s", pythonOutputRepoFile)

		cmdString := fmt.Sprintf("python ./main.py %s %s > %s", repository.Alias, repository.Name, pythonOutputRepoFile)
		cmd := exec.Command("sh", "-c", cmdString)

		if err := cmd.Run(); err != nil {
			fmt.Println("\nError %s", err)
		} else {
			fmt.Println("\n")

			jsonFile := pythonOutputRepoFile

			file, err := os.Open(jsonFile)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer file.Close()

			byteValue, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			var pydrillerInfosDto shared.PydrillerInfosDto
			err = json.Unmarshal(byteValue, &pydrillerInfosDto)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			contributorsInfoJsonData, err := json.Marshal(pydrillerInfosDto.ContributorsInfo)
			if err != nil {
				fmt.Println("Error converting to JSON:", err)
				return
			}

			builder := shared.RepositoryDtoBuilder{}.
				FromRepository(repository).
				WithProjectCommits(pydrillerInfosDto.Commits).
				WithCommitsIntervalInDays(pydrillerInfosDto.CommitsIntervalInDays).
				WithContributorsInfo(string(contributorsInfoJsonData)).
				WithHasCommitsInInterval(pydrillerInfosDto.HasCommitsInInterval).
				WithLastCommitDateInterval(pydrillerInfosDto.LastCommitDateInterval)

				// WithProjectContributors(pydrillerInfosDto.Contributors).

			if !pydrillerInfosDto.HasCommitsInInterval || pydrillerInfosDto.CommitsIntervalInDays > 15 {
				builder = builder.WithUseNotAgile()
			} else {
				builder = builder.WithUseAgile()
			}

			newRepository := builder.Build()

			err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
			if err != nil {
				fmt.Println("Error ao atualizar repositório:", err)
			}
			// else {
			// 	fmt.Println("Repositório atualizado:")
			// }
		}
	}

	fmt.Println("######## finished run-check-agile-and-behaviors ########")
}

func main() {
	var dataSourceName = "../../database/sqlite/repository-dataset.db"

	// alterTableSQL := `
	//   ALTER TABLE repositories ADD X INTEGER DEFAULT 0;
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

	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	var repositories []shared.RepositoryDto

	repositories, err = newSQLiteRepository.GetCloned()
	if err != nil {
		log.Fatal("Error on obter o repositório:", err)
	}

	CheckAndUpdateAgileUse(repositories)
}

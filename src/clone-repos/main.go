package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
	_ "github.com/mattn/go-sqlite3"
)

func WasRepositoryClonned(repository shared.RepositoryDto) {
	var dataSourceName = "../../database/sqlite/repository-dataset.db"
	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	repositoryDtoBuilder := shared.RepositoryDtoBuilder{}.FromRepository(repository)
	newRepository := repositoryDtoBuilder.WithWasCloned().Build()
	err = newSQLiteRepository.UpdateById(repository.Id, newRepository)
	if err != nil {
		fmt.Println("Error, was not possible update repository", err)
	}
}

func main() {
	var dataSourceName = "../../database/sqlite/repository-dataset.db"

	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	var repositories []shared.RepositoryDto

	repositories, err = newSQLiteRepository.GetAll()
	if err != nil {
		log.Fatal("Error on save repository:", err)
	}

	for index, repository := range repositories {
		fmt.Printf("Index: %d\n", index)

		dir := fmt.Sprintf("../../repos/%s", repository.Alias)
		fmt.Printf("dir: %s\n", dir)

		cmd := exec.Command("git", "clone", repository.URL, dir)

		if err := cmd.Run(); err != nil {
			cmd2 := exec.Command("git", "checkout", "-f")
			cmd2.Dir = dir
			if err := cmd2.Run(); err != nil {
				fmt.Printf("\nError on force checkout repository - URL: %v\n", repository.URL)
			}

			cmd3 := exec.Command("git", "pull")
			cmd3.Dir = dir
			if err := cmd3.Run(); err != nil {
				fmt.Printf("\nError on to update repository - URL: %v\n", repository.URL)
			} else {
				WasRepositoryClonned(repository)
			}
		} else {
			WasRepositoryClonned(repository)
		}

		// cmd4 := exec.Command("git", "submodule", "update", "--init", "--recursive")
		// cmd4.Dir = dir
		// if err := cmd4.Run(); err != nil {
		// 	fmt.Printf("\nError on force checkout repository - URL: %v\n", repository.URL)
		// }
	}

	fmt.Println("######## finished run-clone-repos ########")
}

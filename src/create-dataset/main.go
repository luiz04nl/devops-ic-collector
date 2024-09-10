package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/luiz04nl/devops-ic-collector/src/create-dataset/application"
	"github.com/luiz04nl/devops-ic-collector/src/shared"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var repositories []shared.RepositoryDto
	var err error

	repositories, err = application.GetRepositories()
	if err != nil {
		fmt.Println(err)
		return
	}

	var dataSourceName = "../../database/sqlite/repository-dataset.db"

	createTableSQL := `
  CREATE TABLE IF NOT EXISTS repositories (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      url TEXT UNIQUE NOT NULL,
      name TEXT NOT NULL,
      alias TEXT NOT NULL,
      starsTotalCount INTEGER,
      useAgile INTEGER DEFAULT 0,
      useDevOps INTEGER DEFAULT 0,
      useGithubPipelines INTEGER DEFAULT 0,
      useCircleCI INTEGER DEFAULT 0,
      useJenkins INTEGER DEFAULT 0,
      useGitLabPipelines INTEGER DEFAULT 0,
      useAzureDevops INTEGER DEFAULT 0,
      useTravisCI INTEGER DEFAULT 0,
      useHarness INTEGER DEFAULT 0,
      useBitBucketPipelines INTEGER DEFAULT 0,
      wasCloned INTEGER DEFAULT 0,
      meta TEXT DEFAULT "{}",
      projectContributors INTEGER DEFAULT 0,
      projectCommits INTEGER DEFAULT 0,
      commitsIntervalInDays REAL DEFAULT 0,
      contributorsInfo TEXT DEFAULT "{}",
      hasCommitsInInterval INTEGER DEFAULT 0,
      lastCommitDateInterval TEXT DEFAULT NULL,
      projectType TEXT DEFAULT "",
      projectTypeVersion TEXT DEFAULT "",
      projectIssuesEffortTotal INTEGER DEFAULT 0,
      projectIssuesCount INTEGER DEFAULT 0,
      projectCodeSmellsEffortTotal INTEGER DEFAULT 0,
      projectCodeSmellsCount INTEGER DEFAULT 0,
      projectSonarComponentsCount INTEGER DEFAULT 0,
      projectSonarInfo TEXT DEFAULT "",
      bugs TEXT DEFAULT "",
      sqaleRating TEXT DEFAULT "",
      reliabilityRating TEXT DEFAULT "",
      complexity TEXT DEFAULT "",
      cognitiveComplexity TEXT DEFAULT "",
      duplicatedBlocks TEXT DEFAULT "",
      duplicatedFiles TEXT DEFAULT "",
      duplicatedLines TEXT DEFAULT "",
      codeSmells TEXT DEFAULT "",
      linesOfCodesFromSonar TEXT DEFAULT "",
      sqaleIndex TEXT DEFAULT "",
      sqaleDebtRatio TEXT DEFAULT "",
      qualityGateDetails TEXT DEFAULT "",
      vulnerabilities TEXT DEFAULT "",
      securityRating TEXT DEFAULT "",
      classes TEXT DEFAULT "",
      commentLines TEXT DEFAULT "",
      coverage TEXT DEFAULT "",
      tests TEXT DEFAULT "",

      linesOfCodesFromCk TEXT DEFAULT "",
      couplingBetweenObjects TEXT DEFAULT "",
      couplingBetweenObjectsModified TEXT DEFAULT ""
  );
  `

	db, err := sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error on create repositories sqlite table: %v", err)
	}

	newSQLiteRepository, err := shared.NewSQLiteRepository(dataSourceName)
	if err != nil {
		log.Fatal("Was not possible connect with database:", err)
	}

	err = newSQLiteRepository.SaveMultiple(repositories)
	if err != nil {
		log.Fatal("Error on save o repository:", err)
	}

	fmt.Println("######## finished run-create-dataset ########")
}

package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"encoding/csv"
	"os"

	"github.com/luiz04nl/devops-ic-collector/src/shared"
	_ "github.com/mattn/go-sqlite3"
)

func UpdateMetrics(repository shared.RepositoryDto) shared.RepositoryDto {
	csvFile, err := os.Open("../../out/code-metrics/" + repository.Alias + "/class.csv")
	if err != nil {
		fmt.Println(err)
		return repository
	}
	defer csvFile.Close()

	// Leia o conteúdo do CSV
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return repository
	}

	// Prepare a estrutura para armazenar os dados JSON
	var data []map[string]string

	// A primeira linha contém os cabeçalhos
	headers := records[0]

	// Itere sobre as linhas do CSV
	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, value := range record {
			row[headers[i]] = value
		}
		data = append(data, row)
	}

	// LOC (Lines of code): It counts the lines of count, ignoring empty lines and comments (i.e., it's Source Lines of Code, or SLOC). The number of lines here might be a bit different from the original file, as we use JDT's internal representation of the source code to calculate it.
	linesOfCode := 0
	// cbo - CBO (Coupling between objects): Counts the number of dependencies a class has
	couplingBetweenObjects := 0
	// cboModified - CBO Modified (Coupling between objects): Counts the number of dependencies a class has.
	couplingBetweenObjectsModified := 0
	// fanin - FAN-IN: Counts the number of input dependencies a class has, i.e,
	fanIn := 0
	// fanout - FAN-OUT: Counts the number of output dependencies a class has,
	fanOut := 0
	// DIT (Depth Inheritance Tree): It counts the number of "fathers" a class has
	depthInheritanceTree := 0
	// Number of methods: Counts the number of methods. Specific numbers for total number of methods, static, public, abstract, private, protected, default, final, and synchronized methods. Constructor methods also count here.
	numberOfMethods := 0
	// Number of visible methods: Counts the number of visible methods. A method is visible if it is not private.
	numberOfVisibleMethods := 0
	// NOSI (Number of static invocations):
	numberOfStaticMethods := 0
	// WMC (Weight Method Class) or McCabe's complexity. It counts the number of branch instructions in a class.
	weightMethodClass := 0
	// Quantity of returns: The number of return instructions.
	quantityOfReturns := 0
	// Quantity of loops: The number of loops (i.e., for, while, do while, enhanced for).
	quantityOfLoops := 0
	// Quantity of comparisons: The number of comparisons (i.e., == and !=). Note: != is only available in 0.4.2+.
	quantityOfComparisons := 0
	// Quantity of try/catches: The number of try/catches
	quantityOfTryCatches := 0
	// String literals: The number of string literals (e.g., "John Doe"). Repeated strings count as many times as they appear.
	theNumberOfStringLiterals := 0

	for _, row := range data {
		if locNum, err := strconv.Atoi(row["loc"]); err == nil {
			linesOfCode = linesOfCode + locNum
		}

		if cboNum, err := strconv.Atoi(row["cbo"]); err == nil {
			couplingBetweenObjects = couplingBetweenObjects + cboNum
		}

		if cboModifiedNum, err := strconv.Atoi(row["cbo"]); err == nil {
			couplingBetweenObjectsModified = couplingBetweenObjectsModified + cboModifiedNum
		}

		if faninNum, err := strconv.Atoi(row["fanin"]); err == nil {
			fanIn = fanIn + faninNum
		}

		if fanoutNum, err := strconv.Atoi(row["fanout"]); err == nil {
			fanOut = fanOut + fanoutNum
		}

		if ditNum, err := strconv.Atoi(row["dit"]); err == nil {
			depthInheritanceTree = depthInheritanceTree + ditNum
		}

		if totalMethodsQtyNum, err := strconv.Atoi(row["totalMethodsQty"]); err == nil {
			numberOfMethods = numberOfMethods + totalMethodsQtyNum
		}

		if visibleMethodsQtyNum, err := strconv.Atoi(row["visibleMethodsQty"]); err == nil {
			numberOfVisibleMethods = numberOfVisibleMethods + visibleMethodsQtyNum
		}

		if nosiNum, err := strconv.Atoi(row["nosi"]); err == nil {
			numberOfStaticMethods = numberOfStaticMethods + nosiNum
		}

		if wmcNum, err := strconv.Atoi(row["wmc"]); err == nil {
			weightMethodClass = weightMethodClass + wmcNum
		}

		if returnQtyNum, err := strconv.Atoi(row["returnQty"]); err == nil {
			quantityOfReturns = quantityOfReturns + returnQtyNum
		}

		if loopQtyNum, err := strconv.Atoi(row["loopQty"]); err == nil {
			quantityOfLoops = quantityOfLoops + loopQtyNum
		}

		if comparisonsQtyNum, err := strconv.Atoi(row["comparisonsQty"]); err == nil {
			quantityOfComparisons = quantityOfComparisons + comparisonsQtyNum
		}

		if tryCatchQtyNum, err := strconv.Atoi(row["tryCatchQty"]); err == nil {
			quantityOfTryCatches = quantityOfTryCatches + tryCatchQtyNum
		}

		if stringLiteralsQtyNum, err := strconv.Atoi(row["stringLiteralsQty"]); err == nil {
			theNumberOfStringLiterals = theNumberOfStringLiterals + stringLiteralsQtyNum
		}
	}

	builder := shared.RepositoryDtoBuilder{}.
		FromRepository(repository).
		WithLinesOfCodesFromCk(strconv.Itoa(linesOfCode)).
		WithCouplingBetweenObjects(strconv.Itoa(couplingBetweenObjects)).
		WithCouplingBetweenObjectsModified(strconv.Itoa(couplingBetweenObjectsModified))

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

func CheckAndUpdate(repositories []shared.RepositoryDto) {
	for index, repository := range repositories {
		fmt.Printf("Index: %d\n", index)

		now := time.Now()
		fmt.Println("Started at:", now.Format("02/01/2006 3:04:05 PM"))

		shellScriptOutputRepoFile := fmt.Sprintf("%s/%s.out.txt", "../../out/code-metrics", repository.Alias)
		fmt.Printf("shellScriptOutputRepoFile: %s", shellScriptOutputRepoFile)

		cmdString := fmt.Sprintf("./main.sh %s > %s", repository.Alias, shellScriptOutputRepoFile)

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
			UpdateMetrics(repository)
		}
		fmt.Println("\n", output)

		// cmdString := fmt.Sprintf("./main.sh %s > %s", repository.Alias, shellScriptOutputRepoFile)
		// cmd := exec.Command("sh", "-c", cmdString)
		// err := cmd.Run()
		// if err != nil {
		// 	fmt.Println("\nError %s", err)
		// } else {
		// 	UpdateMetrics(repository)
		// }

		now2 := time.Now()
		fmt.Println("\nEnd at:", now2.Format("02/01/2006 3:04:05 PM"))
	}

	fmt.Println("######## finished code-metrics ########")
}

func main() {
	var dataSourceName = "../../database/sqlite/repository-dataset.db"

	// alterTableSQL := `
	//   ALTER TABLE repositories ADD linesOfCodesFromCk TEXT DEFAULT "";
	//   ALTER TABLE repositories ADD couplingBetweenObjects TEXT DEFAULT "";
	//   ALTER TABLE repositories ADD couplingBetweenObjectsModified TEXT DEFAULT "";
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

	CheckAndUpdate(repositories)
}

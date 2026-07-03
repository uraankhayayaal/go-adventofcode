package main

import (
	"fmt"
	"go-adventofcode/src/helpers"
	"log"
	"os"
	"os/exec"
	"plugin"
)

func main() {
	if len(os.Args) == 4 {
		year := os.Args[1]
		if !helpers.IsPathOrFileExists("./src/" + year) {
			log.Fatalf("No year path exist for arg: %s\n", year)
		}

		day := os.Args[2]
		if !helpers.IsPathOrFileExists("./src/" + year + "/" + day) {
			log.Fatalf("No day path exist in %s year path for arg: %s\n", year, day)
		}

		mode := os.Args[3]
		inputFilePath := ""
		switch mode {
		case "prod":
			inputFilePath = "./src/" + year + "/" + day + "/" + "input.txt"
			if !helpers.IsPathOrFileExists(inputFilePath) {
				log.Fatalf("No prod input file: %s\n", inputFilePath)
			}
		case "test":
			inputFilePath = "./src/" + year + "/" + day + "/" + "input_test.txt"
			if !helpers.IsPathOrFileExists(inputFilePath) {
				log.Fatalf("No test input file: %s\n", inputFilePath)
			}
		default:
			log.Fatalf("Wrong mode! Possible values are `prod` or `test`\n")
		}

		plugingGoPath := "./src/" + year + "/" + day + "/" + "solution.go"
		plugingSoPath := "./src/" + year + "/" + day + "/" + "solution.so"
		if !helpers.IsPathOrFileExists(plugingSoPath) {
			log.Fatalf("No solution file: %s\n", plugingSoPath)
		}

		if err := compileSolution(plugingGoPath, plugingSoPath); err != nil {
			panic(fmt.Sprintf("Failed to compile plugin %s: %v", plugingGoPath, err))
		}

		run(plugingSoPath, inputFilePath)
	} else {
		fmt.Println("Wrong count of arguments! Exmaple: `go run . 2025 day8 prod`")
	}
}

func run(plugingPath string, inputFilePath string) {
	p, err := plugin.Open(plugingPath)
	if err != nil {
		log.Fatalf("Error to open file %s: %v\n", plugingPath, err)
	}

	symbol, err := p.Lookup("Solution")
	if err != nil {
		log.Fatalf("No solution method in file %s: %v\n", plugingPath, err)
	}

	solutionFunc, ok := symbol.(func(string))
	if !ok {
		log.Fatalf("Unexpected function signature in file %s: %v\n", plugingPath, err)
	}

	solutionFunc(inputFilePath)
}

func compileSolution(plugingGoPath string, plugingSoPath string) error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", plugingSoPath, plugingGoPath)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

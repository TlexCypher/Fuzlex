package main

import (
	"Fuzlex/src/adapter/controllers"
	"Fuzlex/src/application/usecase"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	algorithmName := parseArgs()
	fzfController := controllers.FzfController{FzfUsecase: &usecase.FzfUsecaseInteractor{}}
	fzfController.Launch(algorithmName, walkAll())
}

func parseArgs() string {
	if len(os.Args) < 2 {
		log.Printf("Invalid arguments.\n")
		log.Printf("  Options\n: -c: Complete Match\n-f Fuzzy Match\n")
		log.Fatalln("Default is Complete Match.")
	}
	return os.Args[1]
}

func walkAll() []*os.File {
	here, err := os.Getwd()
	if err != nil {
		log.Fatalln("Failed to get current working directory.")
	}
	dirs := make([]*os.File, 0)
	filepath.WalkDir(here, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalln("failed filepath.WalkDir")
			return err
		}
		if info.IsDir() && !strings.Contains(path, ".git") {
			fmt.Println(path)
			f, err := os.Open(path)
			if err != nil {
				log.Fatalf("Failed to open file: %v", f.Name())
				return err
			}
			dirs = append(dirs, f)
			return nil
		}
		return nil
	})
	log.Printf("Files length: %v\n", len(dirs))
	for _, v := range dirs {
		fmt.Println(v.Name())
	}
	return dirs
}

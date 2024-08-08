package main

import (
	"Fuzlex/src/adapter/controllers"
	"Fuzlex/src/application/usecase"
	"Fuzlex/src/share/logger"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var logging *log.Logger

func main() {
	logging = logger.GetLogger()
	algorithmName := parseArgs()
	fzfController := controllers.FzfController{FzfUsecase: &usecase.FzfUsecaseInteractor{}}
	fzfController.Launch(algorithmName, walkAll())
}

func parseArgs() string {
	if len(os.Args) < 2 {
		logging.Printf("Invalid arguments.\n")
		logging.Printf("  Options\n: -c: Complete Match\n-f Fuzzy Match\n")
		logging.Fatalln("Default is Complete Match.")
	}
	return os.Args[1]
}

func walkAll() []*os.File {
	here, err := os.Getwd()
	if err != nil {
		logging.Fatalln("Failed to get current working directory.")
	}
	dirs := make([]*os.File, 0)
	filepath.WalkDir(here, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			logging.Fatalln("failed filepath.WalkDir")
			return err
		}
		if info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				logging.Fatalf("Failed to open file: %v", f.Name())
				return err
			}
			dirs = append(dirs, f)
			return nil
		}
		return nil
	})
	logging.Printf("Files length: %v\n", len(dirs))
	return dirs
}

package controllers

import (
	"Fuzlex/src/application/usecase"
	"Fuzlex/src/share/const"
	"os"
)

type FzfController struct {
	FzfUsecase usecase.FzfUsecase
}

func (c *FzfController) Launch(algorithmName string, files []*os.File) {
	c.FzfUsecase.Find(usecase.FzfInputData{
		Glob:          constants.ALL,
		AlgorithmName: algorithmName,
		Dirs:          files,
	})
}

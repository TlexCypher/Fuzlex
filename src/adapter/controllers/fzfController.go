package controllers

import (
	"Fuzlex/src/application/usecase"
	"Fuzlex/src/share"
	"os"
)

type FzfController struct {
	FzfUsecase usecase.FzfUsecase
}

func (c *FzfController) Launch(algorithmName string, files []*os.File) {
	c.FzfUsecase.Find(usecase.FzfInputData{
		Glob:          share.ALL,
		AlgorithmName: algorithmName,
		Dirs:          files,
	})
}

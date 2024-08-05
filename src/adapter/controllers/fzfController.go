package controllers

import (
	"Fuzlex/src/application/usecase"
	"Fuzlex/src/share"
)

type FzfController struct {
	FzfUsecase usecase.FzfUsecase
}

func (c *FzfController) Launch(algorithm string, paths []string) {
	c.FzfUsecase.Find(usecase.FzfInputData{
		Filename:      share.EMPTY,
		AlgorithmName: algorithm,
		Paths:         paths,
	})
}

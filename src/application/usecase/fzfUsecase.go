package usecase

import (
	"Fuzlex/src/application/factory"
	"Fuzlex/src/domain"
	"os"
)

type FzfUsecase interface {
	Find(input FzfInputData) *FzfOutputData
}

type FzfInputData struct {
	Filename      string
	AlgorithmName string
	Paths         []string
}

type FzfOutputData struct {
	Files []*os.File
}

type FzfUsecaseInteractor struct{}

func (i *FzfUsecaseInteractor) Find(input FzfInputData) {
	algorithmFactory := factory.AlgorithmFactory{
		AlgorithmName: input.AlgorithmName,
	}
	algorithm := algorithmFactory.CreateAlgorithm()
	output := FzfOutputData{
		Files: walkWithAlgo(algorithm),
	}
	uiFactory := factory.UIFactory{
		Files: output.Files,
	}
	ui := uiFactory.CreateUI()
	ui.ShowDialog()
}

func walkWithAlgo(algorithm domain.Algorithm) []*os.File {
	return nil
}

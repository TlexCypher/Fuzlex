package usecase

import (
	"Fuzlex/src/application/factory"
	"Fuzlex/src/domain"
	"os"
)

type FzfUsecase interface {
	Find(input FzfInputData)
}

type FzfInputData struct {
	Glob          string
	AlgorithmName string
	Files         []*os.File
}

type FzfOutputData struct {
	Files []*os.File
}

type FzfUsecaseInteractor struct{}

func (i *FzfUsecaseInteractor) Find(input FzfInputData) {
	//algorithmFactory := factory.AlgorithmFactory{
	//	AlgorithmName: input.AlgorithmName,
	//}
	//algorithm := algorithmFactory.CreateAlgorithm()
	output := FzfOutputData{
		Files: input.Files,
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

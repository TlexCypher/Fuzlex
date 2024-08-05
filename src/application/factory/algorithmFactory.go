package factory

import "Fuzlex/src/domain"

type AlgorithmFactory struct {
	AlgorithmName string
}

func (af *AlgorithmFactory) CreateAlgorithm() domain.Algorithm {
	return nil
}

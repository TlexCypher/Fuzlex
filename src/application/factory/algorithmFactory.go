package factory

import (
	"Fuzlex/src/domain"
	"Fuzlex/src/share"
)

type AlgorithmFactory struct {
	AlgorithmName string
}

func (af *AlgorithmFactory) CreateAlgorithm() domain.Algorithm {
	switch af.AlgorithmName {
	case share.COMP_MATCH:
	default:
		return &domain.CompMatchAlgo{}
	}
	return nil
}

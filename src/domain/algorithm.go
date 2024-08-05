package domain

type Algorithm interface {
	Match(input string, target string) bool
}

type CompMatchAlgo struct{}

func (cma *CompMatchAlgo) Match(input string, target string) bool {
	return input == target
}

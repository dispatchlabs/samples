package stv

import "fmt"

type Candidate struct {
	Name 			string		`json:"name,omitempty"`
	//Description   	string
}

func GetCandidates() []Candidate {
	candidates := make([]Candidate, 0)
	for i := 0; i < 10; i++ {
		candidates = append(candidates, Candidate{fmt.Sprintf("Delegate-%d", i)})
	}
	return candidates;
}
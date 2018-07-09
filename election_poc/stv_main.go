package main

import (
	"fmt"
	"github.com/dispatchlabs/samples/election_poc/stv"
	"github.com/dispatchlabs/samples/election_poc/types"
	"github.com/dispatchlabs/samples/election_poc/mock"
)

func main() {
	fmt.Println("In Main")

	candidates := mock.GetCandidates()
	//for _, cand := range candidates {
	//	fmt.Println(cand.Name)
	//}
	ballots := mock.BuildMockBallots(candidates, 25)

	election := stv.Election {
		5,
		ballots,
		0.0,
		stv.NewElectionResults(),
		[]types.Candidate{},
		candidates,
		[]types.Candidate{},
		makeCandidateMap(candidates),
	}
	election.DoElection()
}

func makeCandidateMap(candidates []*types.Candidate) map[string]*types.Candidate {
	candMap := map[string]*types.Candidate{}
	for _, cand := range candidates {
		candMap[cand.Name] = cand
	}
	return candMap
}
package main

import (
	"fmt"
	"github.com/dispatchlabs/samples/election_poc/stv"
)

func main() {
	fmt.Println("In Main")

	candidates := stv.GetCandidates()
	//for _, cand := range candidates {
	//	fmt.Println(cand.Name)
	//}
	ballots := stv.BuildMockBallots(candidates, 25)
	//ballots = append(ballots, stv.NewMockBallet("Bob", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Chris", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Greg", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Nicolae", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Zane", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Avery", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Denis", candidates))
	//ballots = append(ballots, stv.NewMockBallet("Dmitrey", candidates))

	election := stv.Election {
		5,
		ballots,
		0.0,
		stv.NewElectionResults(),
		[]stv.Candidate{},
		candidates,
		[]stv.Candidate{},
	}
	election.DoElection()
}

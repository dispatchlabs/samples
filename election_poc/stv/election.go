package stv

import (
	"fmt"
	"sort"
	"encoding/json"
)

type Election struct {
	NbrVacancies		int64				`json:"nbrVacancies,omitempty"`
	Ballots				[]Ballot			`json:"-"`
	Droop				float64				`json:"droop,omitempty"`
	ElectionResults		*ElectionResults	`json:"electionResult,omitempty"`
	Elected				[]Candidate			`json:"elected,omitempty"`
	Hopefuls			[]Candidate			`json:"hopefuls,omitempty"`
	Eliminated			[]Candidate			`json:"eliminated,omitempty"`
}

func (this *Election) DoElection() {
	for _, ballot := range this.Ballots {
		fmt.Printf("Ballot: %v\n", ballot.ToJson())
	}
	nbrBallots := float64(len(this.Ballots))
	denom := float64(this.NbrVacancies + 1)
	this.Droop =  float64(nbrBallots / denom) + 1

	fmt.Printf("Droop = %f\n", this.Droop)
	counter := &map[string]float64{}
	//var results []ElectionResult
	var roundNbr int64

	roundNbr = 1
	for int64(len(this.ElectionResults.ElectionResults)) < this.NbrVacancies {
		results := this.ExecuteRound(*counter, roundNbr)
		this.ElectionResults.ElectionResults = append(this.ElectionResults.ElectionResults, results.ElectionResults...)
		sort.Sort(sort.Reverse(results))

		fmt.Printf(results.ToPrettyJson())
		roundNbr++
	}
	//electionResult := ElectionResults{results}
	//sort.Sort(sort.Reverse(electionResult))
	//fmt.Printf(electionResult.ToPrettyJson())

}

func (this *Election) ExecuteRound(counter map[string]float64, roundNumber int64) *ElectionResults {
	for _, ballot := range this.Ballots {
		for _, vote := range ballot.Votes {
			if vote.Rank == roundNumber {
				if !this.isElected(vote.Candidate.Name) {
					counter[vote.Candidate.Name] = counter[vote.Candidate.Name] + 1
				}
			}
		}
	}
	voteCounts := make([]VoteCount, 0)
	for k, v := range counter {
		vote := VoteCount{
			Count: v,
			Candidate: Candidate{k},
		}
		voteCounts = append(voteCounts, vote)
		//fmt.Printf("%s :: %d\n", k, v)
	}
	electionRound := ElectionRound{VoteCount: voteCounts}
	sort.Sort(sort.Reverse(electionRound))

	electionResults := electionRound.CountRound(this.Droop, roundNumber)
	//hopefuls := make([]Candidate, 0)
	//elected = append(elected, result.Candidate)

	for _, result := range electionResults.ElectionResults {
		if result.ElectionRound == roundNumber {
			var distributions []Distribution
			counter, distributions = this.FractionalRedistributionWinner(&result, counter, roundNumber)
			result.Distributions = distributions
		}
	}
	fmt.Printf("\n%v", electionRound.ToPrettyJson())
	for k, v := range counter {
		fmt.Printf("\nCOUNT: %v :: %v", k, v)
	}
	fmt.Printf("\n%v", this.ToPrettyJson())
	for k, v := range counter {
		fmt.Printf("\nCOUNT: %v :: %v", k, v)
	}
	return electionResults
}

func (this *Election) ReconcileRound() {

}

/*
    allocated = {} # The allocation of ballots to candidates
    vote_count = {} # A hash of ballot counts, indexed by candidates
    candidates = [] # All candidates
    elected = [] # The candidates that have been elected
    hopefuls = [] # The candidates that may be elected
    # The candidates that have been eliminated because of low counts
    eliminated = []
    # The candidates that have been eliminated because of quota restrictions
    rejected = []
 */

func (this Election) ToPrettyJson() string {
	jsn, err := json.MarshalIndent(this, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(jsn)
}

package stv

import (
	"github.com/gin-gonic/gin/json"
)

type ElectionRound struct {
	VoteCount 	[]VoteCount	`json:"voteCount,omitempty"`
}

type VoteCount struct {
	Candidate 	Candidate	`json:"candidate,omitempty"`
	Count 		float64		`json:"count,omitempty"`
}

// - Implementation of the sort interface

// - Len is part of sort.Interface.
func (this ElectionRound) Len() int {
	return len(this.VoteCount)
}

// - Swap is part of sort.Interface.
func (this ElectionRound) Swap(i, j int) {
	this.VoteCount[i], this.VoteCount[j] = this.VoteCount[j], this.VoteCount[i]
}

// - Less is part of sort.Interface. We use Aroma Value (similarity) as the value to sort by
func (this ElectionRound) Less(i, j int) bool {
	return this.VoteCount[i].Count < this.VoteCount[j].Count
}

func (this ElectionRound) ToJson() []byte {
	jsn, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return jsn

}

func (this ElectionRound) ToPrettyJson() string {
	jsn, err := json.MarshalIndent(this, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(jsn)
}

func (this ElectionRound) CountRound(droop float64, roundNbr int64) []ElectionResult {
	electedCandidates := make([]ElectionResult, 0)
	elected := make([]Candidate, 0)
	for _, vc := range this.VoteCount {
		if vc.Count > droop {
			electedCandidates = append(electedCandidates, ElectionResult{Candidate: vc.Candidate, TotalVotes: vc.Count, ElectionRound: roundNbr, Result: "Elected"})
			elected = append(elected, vc.Candidate)
		}
	}
	return electedCandidates
}
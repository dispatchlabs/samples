package stv

import (
	"encoding/json"
	"github.com/dispatchlabs/samples/election_poc/types"
)

type ElectionRound struct {
	VoteCount 	[]VoteCount			`json:"voteCount,omitempty"`
	Elected		[]*types.Candidate	`json:"elected,omitempty"`
	Eliminated	[]*types.Candidate	`json:"eliminated,omitempty"`

}

type VoteCount struct {
	Candidate 	*types.Candidate	`json:"candidate,omitempty"`
	Count 		float64				`json:"count,omitempty"`
}


func (this *ElectionRound) CountRound(droop float64, roundNbr int64) {
	minCount := droop
	for _, vc := range this.VoteCount {
		if vc.Count > droop {
			this.Elected = append(this.Elected, vc.Candidate)
		} else if vc.Count < minCount {
			minCount = vc.Count
		}
	}
	//Don't eliminate anyone if there isn't at least a single candidate elected -- seems unfair .. someone could come up in the next round
	if len(this.Elected) > 0 {
		for _, vc := range this.VoteCount {
			if vc.Count == minCount {
				this.Eliminated = append(this.Eliminated, vc.Candidate)
			}
		}
	}
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

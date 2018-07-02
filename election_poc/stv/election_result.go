package stv

import (
	"encoding/json"
	"github.com/dispatchlabs/samples/election_poc/types"
)

type ElectionResults struct {
	ElectionResults	[]*ElectionResult 	`json:"electionResults,omitempty"`
	Elected			[]*types.Candidate	`json:"elected,omitempty"`
	Eliminated		[]*types.Candidate	`json:"eliminated,omitempty"`
}

type ElectionResult struct {
	Candidate 		*types.Candidate		`json:"candidate,omitempty"`
	TotalVotes		float64					`json:"votes,omitempty"`
	ElectionRound	int64					`json:"occuredInRound,omitempty"`
	Result          string      			`json:"result,omitempty"`
	Distributions	[]*types.Distribution	`json:"distributions,omitempty"`
}

func NewElectionResults() *ElectionResults {
	return &ElectionResults{[]*ElectionResult{}, []*types.Candidate{}, []*types.Candidate{}}
}

func (this *ElectionResults) UpdateResults(result *ElectionResult) {
	for _, rslt := range this.ElectionResults {
		if rslt.Candidate.Name == result.Candidate.Name {
			if rslt.TotalVotes < result.TotalVotes {
				rslt.TotalVotes = result.TotalVotes
			}
			if len(rslt.Distributions) < len(result.Distributions) {
				//rslt.Distributions = getUniqueDistributions(append(rslt.Distributions, result.Distributions...))
				temp := append(rslt.Distributions, result.Distributions...)
				if len(temp) > 0 {

				}
				//rslt.Distributions = util.Unique(collected)
			}
		}
	}
}

//func getUniqueDistributions(elements []Distribution) []Distribution {
//	// Use map to record duplicates as we find them.
//	encountered := map[string]bool{}
//	var result []Distribution
//
//	for _, v := range elements {
//		if encountered[v.Candidate.Name] == true {
//			// Do not add duplicate.
//		} else {
//			// Record this element as an encountered element.
//			encountered[v.Candidate.Name] = true
//			// Append to result slice.
//			result = append(result, v)
//		}
//	}
//	// Return the new slice.
//	return result
//}

// - Implementation of the sort interface

// - Len is part of sort.Interface.
func (this ElectionResults) Len() int {
	return len(this.ElectionResults)
}

// - Swap is part of sort.Interface.
func (this ElectionResults) Swap(i, j int) {
	this.ElectionResults[i], this.ElectionResults[j] = this.ElectionResults[j], this.ElectionResults[i]
}

// - Less is part of sort.Interface. We use Aroma Value (similarity) as the value to sort by
func (this ElectionResults) Less(i, j int) bool {
	return this.ElectionResults[i].TotalVotes < this.ElectionResults[j].TotalVotes
}

func (this ElectionResults) ToJson() []byte {
	jsn, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return jsn

}

func (this ElectionResults) ToPrettyJson() string {
	jsn, err := json.MarshalIndent(this, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(jsn)
}


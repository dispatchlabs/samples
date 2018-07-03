package types

import (
	"fmt"
	"encoding/json"
)

type Distribution struct {
	Candidate 			*Candidate	`json:"candidate,omitempty"`
	DistributedVotes	*float64	`json:"distributedValue,omitempty"`
}

type Candidate struct {
	Name 			string			`json:"name,omitempty"`
	CurrentVotes	float64			`json:"currentVoteCount,omitempty"`
	EffectiveVotes  float64			`json:"effectiveVoteCount,omitempty"`
	ElectionStatus	string			`json:"electionStatus,omitempty"`
	Distributions	[]*Distribution	`json:"-"`

	//Description   	string
}

func GetCandidates() []*Candidate {
	candidates := make([]*Candidate, 0)
	for i := 0; i < 10; i++ {
		candidates = append(candidates, &Candidate{fmt.Sprintf("Delegate-%d", i), 0, 0,"Hopefull", []*Distribution{}})
	}
	return candidates;
}

func (this *Candidate) AddVotes(value float64) {
	//utils.Info(fmt.Sprintf("AddVotes: %v :: %v ", this.Name, this.CurrentVotes))
	this.CurrentVotes += value
}

func (this *Candidate) SetStatus(status string) {
	this.ElectionStatus = status
}

func (this *Candidate) AddDistribution(candidate *Candidate, value *float64) {
	this.Distributions = append(this.Distributions, &Distribution{Candidate: candidate, DistributedVotes: value})
	this.EffectiveVotes -= *value
}


func (this Candidate) ToJson() string {
	jsn, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	val := string(jsn)
	return val
}

func (this Candidate) ToPrettyJson() string {
	jsn, err := json.MarshalIndent(this, "", "\t")
	if err != nil {
		panic(err)
	}
	val := string(jsn)
	return val
}

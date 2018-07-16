package types

import (
	"encoding/json"
)

type Ballot struct {
	Address		string		`json:"address,omitempty"`
	Votes		[]Vote		`json:"votes,omitempty"`
	Stake		int64		`json:"stake,omitempty"`
	Status      string      `json:"status,omitempty"`
}

type Vote struct {
	Candidate 	*Candidate	`json:"candidate,omitempty"`
	Rank		int64		`json:"rank,omitempty"`
}

func (this *Ballot) SetStatus(status string) {
	this.Status = status
}

func (this Ballot) ToJson() string {
	jsn, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return string(jsn)
}

func (this Ballot) ToPrettyJson() string {
	jsn, err := json.MarshalIndent(this, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(jsn)
}

func (this *Ballot) UpdateForStatusChangeAfterElection(candidate *Candidate, roundNbr int64) {
	if this.Status == StatusApplied {
		return
	}
	for _, v := range this.Votes {
		if v.Candidate.Name == candidate.Name && v.Rank <= roundNbr {
			this.Status = StatusApplied
		}
	}
}

func (this *Ballot) CheckForStatusChangeAfterElimination(candidate *Candidate, roundNbr int64) {
	for _, v := range this.Votes {
		if this.Status == StatusUncounted && v.Candidate.Name == candidate.Name && v.Rank <= roundNbr {
			this.Status = StatusPending
		}
	}
}
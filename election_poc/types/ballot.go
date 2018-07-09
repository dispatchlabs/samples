package types

import (
	"encoding/json"
)

type Ballot struct {
	Address		string		`json:"address,omitempty"`
	Votes		[]Vote		`json:"votes,omitempty"`
	Stake		int64		`json:"stake,omitempty"`
}

type Vote struct {
	Candidate 	*Candidate	`json:"candidate,omitempty"`
	Rank		int64		`json:"rank,omitempty"`
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
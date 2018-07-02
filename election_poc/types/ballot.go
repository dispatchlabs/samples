package types

import (
	"github.com/dispatchlabs/disgo/commons/utils"
	"fmt"
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

/*
Mock data for testing
 */

func BuildMockBallots(candidates []*Candidate, nbrOfBallots int) []Ballot {
	ballots := make([]Ballot, 0)
	for i := 0; i < nbrOfBallots; i++ {
		addr := fmt.Sprintf("Voter-%d", i)
		ballot := NewMockBallot(addr, candidates)
		ballots = append(ballots, ballot)
	}
	return ballots;
}

func NewMockBallot(address string, candidates []*Candidate) Ballot {
	nbrCandidates := len(candidates)

	votes := make([]Vote, 0)
	nbrVotes := utils.Random(1, nbrCandidates)
	//fmt.Printf("Address: %-10v casting %d votes\n", address, nbrVotes)
	encountered := map[string]bool{}

	for i := 1; i <= nbrVotes; i++ {
		candidate := GetRandomUniqueCandidate(encountered, candidates, nbrCandidates)
		encountered[candidate.Name] = true
		votes = append(votes, Vote{candidate,int64(i)})
	}

	return Ballot{
		Address: address,
		Votes: votes,
		Stake: 1,
	}
}

func GetRandomUniqueCandidate(encountered map[string]bool, candidates []*Candidate, nbrCandidates int) *Candidate {

	for {
		randomDelegate := utils.Random(1, nbrCandidates)
		candidate := candidates[randomDelegate]
		if encountered[candidate.Name] == false {
			return candidate
		}
	}
}
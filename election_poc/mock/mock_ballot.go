package mock

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/samples/election_poc/types"
)

/*
Mock data for testing
 */

func BuildMockBallots(candidates []*types.Candidate, nbrOfBallots int) []*types.Ballot {
	ballots := make([]*types.Ballot, 0)
	for i := 0; i < nbrOfBallots; i++ {
		addr := fmt.Sprintf("Voter-%d", i)
		ballot := NewMockBallot(addr, candidates)
		ballots = append(ballots, ballot)
	}
	return ballots;
}

func NewMockBallot(address string, candidates []*types.Candidate) *types.Ballot {
	nbrCandidates := len(candidates)

	votes := make([]types.Vote, 0)
	nbrVotes := utils.Random(1, nbrCandidates)
	//fmt.Printf("Address: %-10v casting %d votes\n", address, nbrVotes)
	encountered := map[string]bool{}

	for i := 1; i <= nbrVotes; i++ {
		candidate := GetRandomUniqueCandidate(encountered, candidates, nbrCandidates)
		encountered[candidate.Name] = true
		votes = append(votes, types.Vote{candidate,int64(i)})
	}
	ballot := &types.Ballot{
		Address: address,
		Votes: votes,
		Stake: 1,
	}
	fmt.Printf(ballot.ToJson())
	return ballot
}

func GetRandomUniqueCandidate(encountered map[string]bool, candidates []*types.Candidate, nbrCandidates int) *types.Candidate {

	for {
		randomDelegate := utils.Random(1, nbrCandidates)
		candidate := candidates[randomDelegate]
		if encountered[candidate.Name] == false {
			return candidate
		}
	}
}

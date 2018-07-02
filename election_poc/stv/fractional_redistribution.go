package stv

import (
	"fmt"
	"github.com/dispatchlabs/samples/election_poc/types"
)

func (this *Election) FractionalRedistributionWinner(candidate *ElectionResult, roundNumber int64) []*types.Distribution {
	distributions := make([]*types.Distribution, 0)
	//check if redistribution is necessary because of excess required votes
	if candidate.ElectionRound == roundNumber && candidate.TotalVotes > this.Droop {
		for _, ballot := range this.Ballots {

			if(this.addNextVote(candidate.Candidate.Name, ballot, roundNumber)) {
				nextCand := ballot.Votes[roundNumber].Candidate
				if(nextCand.ElectionStatus == "Hopefull") {
					addPartial := ( (float64(candidate.TotalVotes - this.Droop) / candidate.TotalVotes))
					fmt.Printf("\nAdding %v to candidate %s for voter %v\n", addPartial, nextCand, ballot.Address)
					nextCand.AddVotes(addPartial)
					distributions = append(distributions, &types.Distribution{nextCand, &addPartial})
				}
			}
		}
	}
	candidate.Distributions = distributions
	this.ElectionResults.UpdateResults(candidate)
	//for k, v := range counter {
	//	if(!this.isElected(k)) {
	//		if(v > this.Droop) {
	//			cand := Candidate{k}
	//			this.Elected = append(this.Elected, cand)
	//			newResult := ElectionResult{cand, v, roundNumber, "Elected", []Distribution{}}
	//			_, distribution := this.FractionalRedistributionWinner(&newResult, counter, roundNumber)
	//			newResult.Distributions = distribution
	//			this.ElectionResults.ElectionResults = append(this.ElectionResults.ElectionResults, newResult)
	//		} else {
	//			//fmt.Printf("Next Candidates = %v with votes = %v\n", k, v)
	//		}
	//
	//	}
	//}
	return distributions
}

func (this *Election) tallyVotes(candidateToRedistribute string, ballot types.Ballot, roundNumber int64) {

}

func (this *Election) addNextVote(candidateToRedistribute string, ballot types.Ballot, roundNumber int64) bool {
	for _, vote := range ballot.Votes {
		//find second vote for votes that were for candidates elected in this round
		if vote.Rank == roundNumber && vote.Candidate.Name == candidateToRedistribute {
			if int64(len(ballot.Votes)) > roundNumber {
				return true
			}
		}
	}
	return false
}

func (this *Election) SendToNextValidCandidate(candidateToRedistribute *types.Candidate, ballot types.Ballot, roundNumber int64) *types.Candidate {
	var result *types.Candidate
	start := false
	for _, vote := range ballot.Votes {
		//making sure to redistribute the correct one ...
		if vote.Rank == roundNumber && vote.Candidate.Name == candidateToRedistribute.Name {
			start = true
		}
		if start {
			//find next vote for votes that were for candidates elected in this round
			if int64(len(ballot.Votes)) > roundNumber && vote.Candidate.ElectionStatus == "Hopefull" {
				addPartial := ( (float64(vote.Candidate.CurrentVotes - this.Droop) / vote.Candidate.CurrentVotes))
				fmt.Printf("\nAdding %v to candidate %v for voter %v\n", addPartial, vote.Candidate.ToJson(), ballot.Address)
				vote.Candidate.AddVotes(addPartial)
				candidateToRedistribute.AddDistribution(vote.Candidate, &addPartial)
				break
			}

		}
	}
	return result
}

package stv

import (
	"fmt"
	"github.com/dispatchlabs/samples/election_poc/types"
)


func (this *Election) Redistribute(electedThisRound []*types.Candidate, roundNbr int64) {
	subsequentlyElected := make([]*types.Candidate, 0)
	for _, cand := range electedThisRound {
		distributionList := make([]*types.Candidate, 0)
		cand.EffectiveVotes = cand.CurrentVotes
		for _, ballot := range this.Ballots {
			distCandidate := this.SendToNextValidCandidate(cand, ballot, roundNbr)
			if distCandidate != nil {
				distributionList = append(distributionList, distCandidate)
			}
		}
		for _, dCand := range distributionList {
			addPartial := ( (float64(cand.CurrentVotes - this.Droop) / float64(len(distributionList))))
			dCand.AddVotes(addPartial)
			fmt.Printf("Adding %v to candidate %v from candidate %s\n", addPartial, dCand.ToJson(), cand.Name)
			cand.AddDistribution(dCand, &addPartial)
			if dCand.CurrentVotes >= this.Droop  && dCand.ElectionStatus == "Hopefull" {
				dCand.ElectionStatus = "Elected"
				subsequentlyElected = append(subsequentlyElected, dCand)
				fmt.Printf("subsequentlyElected %v\n", dCand.Name)
				this.ElectionResults.Elected = append(this.ElectionResults.Elected, dCand)
			}
		}
	}
}

//So also need to figure out up front how many votes are actually possible to distribute
//If the ballot has no more votes because lack of entries, all others have been elected or eliminated
//Once this is done, then you can do the distribution appropriately.
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
			if int64(len(ballot.Votes)) > roundNumber {
				candidate := this.getCandidate(vote.Candidate.Name)
				if candidate.ElectionStatus == "Hopefull" {
					result = vote.Candidate
					break
				}
			}
		}
	}
	return result
}


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


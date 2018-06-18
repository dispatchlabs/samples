package stv

func (this *Election) FractionalRedistributionWinner(candidate *ElectionResult, counter map[string]float64, roundNumber int64) (map[string]float64, []Distribution) {
	distributions := make([]Distribution, 0)
	//check if redistribution is necessary because of excess required votes
	if candidate.ElectionRound == roundNumber && candidate.TotalVotes > this.Droop {
		for _, ballot := range this.Ballots {
			if(this.addNextVote(candidate.Candidate.Name, ballot, roundNumber)) {
				nextCand := ballot.Votes[roundNumber].Candidate.Name
				if(!this.isElected(nextCand)) {
					addPartial := ( (float64(candidate.TotalVotes - this.Droop) / candidate.TotalVotes))
					//fmt.Printf("\nAdding %v to candidate %s for voter %v\n", addPartial, nextCand, ballot.Address)
					counter[nextCand] = counter[nextCand] + addPartial
					distributions = append(distributions, Distribution{Candidate{nextCand}, addPartial})
				}
			}
		}
	}
	for k, v := range counter {
		if(!this.isElected(k)) {
			if(v > this.Droop) {
				cand := Candidate{k}
				this.Elected = append(this.Elected, cand)
				newResult := ElectionResult{cand, v, roundNumber, "Elected", []Distribution{}}
				_, distribution := this.FractionalRedistributionWinner(&newResult, counter, roundNumber)
				newResult.Distributions = distribution
				this.ElectionResults = append(this.ElectionResults, newResult)
			} else {
				//fmt.Printf("Next Candidates = %v with votes = %v\n", k, v)
			}

		}
	}
	return counter, distributions
}

func (this *Election) tallyVotes(candidateToRedistribute string, ballot Ballot, roundNumber int64) {

}

func (this *Election) addNextVote(candidateToRedistribute string, ballot Ballot, roundNumber int64) bool {
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

func (this *Election) isElected(candidateName string) bool {
	result := false
	for _, elected := range this.Elected {
		if elected.Name == candidateName {
			result = true
		}
	}
	return result
}

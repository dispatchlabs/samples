package mock

import (
	"fmt"
	"github.com/dispatchlabs/samples/election_poc/types"
)

func GetCandidates() []*types.Candidate {
	candidates := make([]*types.Candidate, 0)
	for i := 0; i < 10; i++ {

		candidate := &types.Candidate{
			fmt.Sprintf("Delegate-%d", i),
			0,
			0,
			types.StatusHopefull,
			[]*types.Distribution{},
		}
		candidates = append(candidates, candidate)
	}
	return candidates;
}

package util

import (
	"testing"
	"time"
	"fmt"
	"github.com/dispatchlabs/disgo/commons/utils"
)

//Used to make sure I had calculated time correctly.  Leaving because it's handy for testing timex
func TestGetDurationDelta(t *testing.T) {
	dfltDuration := time.Second * 5
	now := time.Now()
	future := utils.ToMilliSeconds(now.Add(time.Second * 2))
	delta := future - utils.ToMilliSeconds(time.Now())  //do milliseconds since that's what you need

	delta2 := time.Millisecond * time.Duration(delta) + dfltDuration

	durationDelta := now.Add(delta2)
	fmt.Printf("Delta: %v :: Delta2: %v :: timeDelta: %v\n", delta, delta2, durationDelta)
}

package tween_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/judah-caruso/core/tween"
)

func TestTween(t *testing.T) {
	var (
		tw   *tween.Tween
		done bool
		opts = tween.Opts{
			Ease:       tween.EaseLinear,
			Transition: tween.TransIn,
			Duration:   time.Second * 1,
			Delay:      0,
		}
	)

	v := 0.0
	i := 0

	_ = tw

	for {
		if done, tw = tween.TweenValue(0.1, &v, v+10, &opts); done {
			if v-10 > 0.0001 {
				t.Error("final tweened value was not 10")
			}
			break
		}

		fmt.Printf("v at %d - %f\n", i, v)
		i += 1
	}

	if !done {
		t.Error("expected tween to finish after 10 steps")
	}
}

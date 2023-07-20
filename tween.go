/*
# Tween

Immediate mode tweening.
*/
package core

import (
	"encoding/binary"
	"hash/fnv"
	"math"
	"runtime"
	"time"
)

type Ease int

const (
	EaseLinear Ease = iota
	EaseQuad
	EaseCubic
	EaseQuart
	EaseQuint
	EaseExpo
	EaseSine
	EaseCirc
	EaseBack
	EaseElastic
)

type Transition int

const (
	TransIn Transition = iota
	TransOut
	TransInOut
)

type TweenOpts struct {
	Ease       Ease
	Transition Transition
	Duration   time.Duration
	Delay      time.Duration
}

type Tween struct {
	TweenOpts
	hash     uint
	rate     float64
	progress float64
	delay    float64
	values   []value
}

type value struct {
	to    float64
	from  float64
	delta float64
	value *float64
}

func TweenValue(dt float64, from *float64, to float64, opts *TweenOpts) (done bool, tween *Tween) {
	if opts == nil {
		opts = defaultOpts
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("unable to get caller information")
	}

	hash := hashLoc(file, line)

	tween = findOrCreateTween(hash, opts)
	tween.addValue(from, to)

	done = tween.Update(dt)
	return
}

func (t *Tween) Update(dt float64) (done bool) {
	if !tweenIsManaged(t) {
		return false
	}

	if t.progress >= 1 {
		return true
	}

	if t.delay > 0 {
		t.delay -= dt
		return false
	}

	t.progress += t.rate * dt
	if t.progress >= 1 {
		return true
	}

	p := t.step()
	for i := 0; i < len(t.values); i += 1 {
		v := &t.values[i]
		*v.value = v.from + p*v.delta
	}

	return false
}

func (t *Tween) Reverse() {
	for i := 0; i < len(t.values); i += 1 {
		v := &t.values[i]
		tmp := v.from
		v.from = v.to
		v.to = tmp
		v.delta = -v.delta
	}
}

func (t *Tween) Reset() {
	durInSec := t.Duration.Seconds()
	if durInSec > 0 {
		t.rate = durInSec
	}

	delayInSec := t.Delay.Seconds()
	if delayInSec > 0 {
		t.delay = delayInSec
	}

	t.progress = 0
}

func (t *Tween) Delete() {
	t.values = t.values[:0]
	delete(managedTweens, t.hash)
}

func (t *Tween) step() float64 {
	p := t.progress
	trans := t.Transition
	ease := t.Ease

	switch trans {
	case TransIn:
		return applyEase(p, ease)

	case TransOut:
		p = 1 - p
		return 1 - applyEase(p, ease)

	case TransInOut:
		p *= 2
		if p < 1 {
			return 0.5 * applyEase(p, ease)
		}

		p = 2 - p
		return 0.5*(1-applyEase(p, ease)) + 0.5
	}

	return p
}

func applyEase(p float64, ease Ease) float64 {
	switch ease {
	case EaseQuad:
		p = p * p
	case EaseCubic:
		p = p * p * p
	case EaseQuart:
		p = p * p * p * p
	case EaseQuint:
		p = p * p * p * p * p
	case EaseExpo:
		p = math.Pow(2, 10*(p-1))
	case EaseSine:
		p = -math.Cos(p*(math.Pi*0.5)) + 1
	case EaseCirc:
		p = -(math.Sqrt(1-(p*p)) - 1)
	case EaseBack:
		p = p * p * (2.7*p - 1.7)
	case EaseElastic:
		p = -math.Pow(2, 10*p-10) * math.Sin((p*10-10.75)*((2*math.Pi)/3))
	}

	// Also handles EaseLinear
	return p
}

func (t *Tween) addValue(from *float64, to float64) {
	delta := to - *from
	if delta != 0 {
		t.values = append(t.values, value{
			to:    to,
			from:  *from,
			delta: delta,
			value: from,
		})
	}
}

func findOrCreateTween(hash uint, opts *TweenOpts) *Tween {
	tween, ok := managedTweens[hash]
	if !ok {
		if opts == nil {
			panic("attempt to create tween with no opts")
		}

		var (
			rate  float64
			delay float64
		)

		durInSec := opts.Duration.Seconds()
		if durInSec > 0 {
			rate = durInSec
		}

		delayInSec := opts.Delay.Seconds()
		if delayInSec > 0 {
			delay = delayInSec
		}

		tween = new(Tween)
		*tween = Tween{
			TweenOpts: *opts,
			hash:      hash,
			rate:      rate,
			delay:     delay,
			progress:  0,
			values:    make([]value, 0),
		}

		managedTweens[hash] = tween
	}

	return tween
}

var (
	managedTweens = make(map[uint]*Tween)
	defaultOpts   = &TweenOpts{
		Ease:       EaseQuart,
		Transition: TransInOut,
		Duration:   time.Second * 3,
		Delay:      0,
	}
)

func tweenIsManaged(t *Tween) bool {
	if t == nil {
		return false
	}

	_, ok := managedTweens[t.hash]
	return ok
}

func hashLoc(file string, line int) uint {
	hash := fnv.New32a()
	hash.Write([]byte(file))

	var b [8]byte
	binary.PutUvarint(b[:], uint64(line))
	hash.Write(b[:])

	return uint(hash.Sum32())
}

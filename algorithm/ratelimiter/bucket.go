package algorithm

import (
	"math"
	"sync"
	"time"
)

type Reservation struct {
	ok        bool
	lim       *Limiter
	tokens    int
	limit     Limit
	timeToAct time.Time
}

// A zero Limit allows no events.
type Limit float64

const InfDuration = time.Duration(math.MaxInt64)

// durationFromTokens is a unit conversion function from the number of tokens to the duration
// of time it takes to accumulate them at a rate of limit tokens per second.
func (limit Limit) durationFromTokens(tokens float64) time.Duration {
	if limit <= 0 {
		return InfDuration
	}
	seconds := tokens / float64(limit)
	return time.Duration(float64(time.Second) * seconds)
}

// tokensFromDuration is a unit conversion function from a time duration to the number of tokens
// which could be accumulated during that duration at a rate of limit tokens per second.
func (limit Limit) tokensFromDuration(d time.Duration) float64 {
	if limit <= 0 {
		return 0
	}
	return d.Seconds() * float64(limit)
}

// Inf is the infinite rate limit; it allows all events (even if burst is zero).
const Inf = Limit(math.MaxFloat64)

// 采用延迟计算token方法
type Limiter struct {
	last      time.Time
	lastEvent time.Time
	burst     int
	tokens    float64
	limit     Limit
	mu        sync.Mutex
}

func (lim *Limiter) advance(t time.Time) (time.Time, float64) {
	last := lim.last
	if t.Before(last) {
		last = t
	}

	elapsed := t.Sub(last)
	delta := lim.limit.tokensFromDuration(elapsed)
	tokens := lim.tokens + delta
	if burst := float64(lim.burst); tokens > burst {
		tokens = burst
	}

	return t, tokens
}

func (lim *Limiter) Allow() {
}

func (lim *Limiter) AllowN() {
}

func (lim *Limiter) Wait() {
}

func (lim *Limiter) WaitN() {
}

func (lim *Limiter) Reserve() {
}

func (lim *Limiter) ReserveN() {
}

// from time t reserve n tokens
func (lim *Limiter) reserveN(t time.Time, n int, maxFutureReserve time.Duration) Reservation {
	lim.mu.Lock()
	defer lim.mu.Unlock()

	if lim.limit == Inf {
		return Reservation{
			ok:        true,
			lim:       lim,
			tokens:    n,
			timeToAct: t,
		}
	}

	// cauculate token from now
	t, tokens := lim.advance(t)

	tokens -= float64(n)

	var waitDuration time.Duration
	if tokens < 0 {
		waitDuration = lim.limit.durationFromTokens(-tokens)
	}

	ok := n <= lim.burst && waitDuration < maxFutureReserve

	r := Reservation{
		ok:    ok,
		lim:   lim,
		limit: lim.limit,
	}

	if ok {
		r.tokens = n
		r.timeToAct = t.Add(waitDuration)

		lim.last = t
		lim.tokens = tokens
		lim.lastEvent = r.timeToAct
	}

	return r
}

func TestLimter() {
	// limiter := rate.NewLimiter(rate.Every(time.Millisecond*31), 2)
	// //time.Sleep(time.Second)
	// for i := 0; i < 10; i++ {
	// 	var ok bool
	// 	if limiter.Allow() {
	// 		ok = true
	// 	}
	// 	time.Sleep(time.Millisecond * 20)
	// 	fmt.Println(ok, limiter.Burst())
	// }
}

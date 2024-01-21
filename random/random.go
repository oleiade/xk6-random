package random

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

// Permutation returns as an array of integers of length n, a random permutation
// of the integers [0, n).
func Permutation(rt *goja.Runtime, n int) goja.Value {
	return rt.ToValue(rand.Perm(n))
}

// Shuffle shuffles the given array in-place.
func Shuffle(rt *goja.Runtime, array goja.Value) {
	arr := array.Export().([]any)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	for i, val := range arr {
		if err := array.ToObject(rt).Set(strconv.Itoa(i), val); err != nil {
			common.Throw(rt, err)
		}
	}
}

// Shuffled returns a shuffled copy of the given array.
func Shuffled(rt *goja.Runtime, array goja.Value) goja.Value {
	// FIXME: this can fail, we should handle it
	arr := array.Export().([]any)

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	return rt.ToValue(arr)
}

type Generator struct {
	randomness *rand.Rand
	seed       int64

	vu modules.VU
}

func NewRandomGenerator(vu modules.VU) *Generator {
	seed := time.Now().UnixNano()

	return &Generator{
		randomness: rand.New(rand.NewSource(seed)),
		seed:       seed,
		vu:         vu,
	}
}

func NewSeededRandomGenerator(vu modules.VU, seed int64) *Generator {
	return &Generator{
		randomness: rand.New(rand.NewSource(seed)),
		seed:       seed,
		vu:         vu,
	}
}

// Int returns a random integer.
func (g Generator) Int() int {
	return g.randomness.Int()
}

// Float returns a random real number.
func (g Generator) Float() float64 {
	return g.randomness.Float64()
}

// Boolean returns a random boolean value.
func (g Generator) Boolean() bool {
	return g.randomness.Intn(2) == 1
}

// IntBetween returns a random integer between the given min and max values.
func (g Generator) IntBetween(min, max int) int {
	return g.randomness.Intn(max-min) + min
}

// FloatBetween returns a random real number between the given min and max
// values.
func (g Generator) FloatBetween(min, max float64) float64 {
	return g.randomness.Float64()*(max-min) + min
}

// Normal returns a random real number from a normal distribution with the given
// mean and standard deviation.
func (g Generator) Normal(mean, stdev float64) float64 {
	return g.randomness.NormFloat64()*stdev + mean
}

// LogNormal returns a random real number from a log-normal distribution with
// the given mean and standard deviation.
func (g Generator) LogNormal(mean, stdev float64) float64 {
	return math.Exp(g.Normal(mean, stdev))
}

// Bernoulli returns a random boolean value with the given probability.
func (g Generator) Bernoulli(probability float64) bool {
	return g.randomness.Float64() < probability
}

// Binomial returns a random integer from a binomial distribution with the given
// number of trials and probability.
func (g Generator) Binomial(trials int, probability float64) int {
	count := 0
	for i := 0; i < trials; i++ {
		if g.Bernoulli(probability) {
			count++
		}
	}

	return count
}

// Geometric returns a random integer from a geometric distribution with the
// given probability.
func (g Generator) Geometric(probability float64) int {
	var result int

	for g.randomness.Float64() >= probability {
		result++
	}

	return result
}

// Exponential returns a random real number from an exponential distribution
// with the given rate.
func (g Generator) Exponential(rate float64) float64 {
	if rate <= 0 {
		common.Throw(g.vu.Runtime(), fmt.Errorf("rate must be positive"))
	}

	return -math.Log(g.randomness.Float64()) / rate
}

// Pick returns a random element from the given array.
func (g Generator) Pick(array goja.Value) goja.Value {
	rt := g.vu.Runtime()

	var arr []any
	if err := rt.ExportTo(array, &arr); err != nil {
		common.Throw(g.vu.Runtime(), err)
	}

	return rt.ToValue(arr[g.randomness.Intn(len(arr))])
}

// WeightedPick returns a random element from the given array, based on the
// given weights.
func (g Generator) WeightedPick(array []goja.Value, weights []float64) goja.Value {
	if len(array) != len(weights) {
		common.Throw(g.vu.Runtime(), fmt.Errorf("array and weights must be of the same length"))
	}

	totalWeight := 0.0
	for _, weight := range weights {
		totalWeight += weight
	}

	threshold := g.randomness.Float64() * totalWeight

	cumulativeWeight := 0.0
	for index, weight := range weights {
		cumulativeWeight += weight

		if cumulativeWeight >= threshold {
			return array[index]
		}
	}

	common.Throw(g.vu.Runtime(), fmt.Errorf("unreachable! weightedPick should never reach here"))

	return goja.Undefined()
}

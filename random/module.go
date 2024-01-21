package random

import (
	"fmt"

	"github.com/dop251/goja"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

type (
	// RootModule is the global module instance that will create Client
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		vu modules.VU
	}
)

// Ensure the interfaces are implemented correctly
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance
func New() *RootModule {
	return &RootModule{}
}

func (rm *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{
		vu: vu,
	}
}

func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{Named: map[string]interface{}{
		"Random": mi.NewRandomGenerator,

		// Permutation returns as an array of integers of length n, a random permutation
		// of the integers [0, n).
		"permutation": func(n int) goja.Value { return Permutation(mi.vu.Runtime(), n) },

		// Shuffle shuffles the given array in-place.
		"shuffle": func(array goja.Value) { Shuffle(mi.vu.Runtime(), array) },

		// Shuffled returns a shuffled copy of the given array.
		"shuffled": func(array goja.Value) goja.Value { return Shuffled(mi.vu.Runtime(), array) },
	}}
}

func (mi *ModuleInstance) NewRandomGenerator(call goja.ConstructorCall) *goja.Object {
	rt := mi.vu.Runtime()

	var rg *Generator

	// If a seed is provided, use it to instantiate a random generator, otherwise
	// use the default one.
	if len(call.Arguments) > 0 {
		seed := call.Argument(0).ToInteger()
		rg = NewSeededRandomGenerator(mi.vu, seed)
	} else {
		rg = NewRandomGenerator(mi.vu)
	}

	obj := rt.NewObject()

	// Set read-only properties on the object we expose to the Runtime.
	must(rt, setReadOnlyPropertyOf(obj, "seed", rt.ToValue(rg.seed)))

	// Define the read-only int method.
	must(rt, setReadOnlyPropertyOf(obj, "int", rt.ToValue(rg.Int)))
	must(rt, setReadOnlyPropertyOf(obj, "float", rt.ToValue(rg.Float)))
	must(rt, setReadOnlyPropertyOf(obj, "boolean", rt.ToValue(rg.Boolean)))
	must(rt, setReadOnlyPropertyOf(obj, "intBetween", rt.ToValue(rg.IntBetween)))
	must(rt, setReadOnlyPropertyOf(obj, "floatBetween", rt.ToValue(rg.FloatBetween)))
	must(rt, setReadOnlyPropertyOf(obj, "pick", rt.ToValue(rg.Pick)))
	must(rt, setReadOnlyPropertyOf(obj, "weightedPick", rt.ToValue(rg.WeightedPick)))
	must(rt, setReadOnlyPropertyOf(obj, "normal", rt.ToValue(rg.Normal)))
	must(rt, setReadOnlyPropertyOf(obj, "logNormal", rt.ToValue(rg.LogNormal)))
	must(rt, setReadOnlyPropertyOf(obj, "bernoulli", rt.ToValue(rg.Bernoulli)))
	must(rt, setReadOnlyPropertyOf(obj, "binomial", rt.ToValue(rg.Binomial)))
	must(rt, setReadOnlyPropertyOf(obj, "geometric", rt.ToValue(rg.Geometric)))

	return rt.ToValue(obj).ToObject(rt)
}

// Shuffled returns a shuffled copy of the given array.

// setReadOnlyPropertyOf sets a read-only property on the given [goja.Object].
func setReadOnlyPropertyOf(obj *goja.Object, name string, value goja.Value) error {
	err := obj.DefineDataProperty(name,
		value,
		goja.FLAG_FALSE,
		goja.FLAG_FALSE,
		goja.FLAG_TRUE,
	)
	if err != nil {
		return fmt.Errorf("unable to define %s read-only property on TextEncoder object; reason: %w", name, err)
	}

	return nil
}

// must throws if the given error is not nil.
func must(rt *goja.Runtime, err error) {
	if err != nil {
		common.Throw(rt, err)
	}
}

// Package kv provides a key-value store extension for k6.
package kv

import (
	"github.com/oleiade/xk6-random/random"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/random", random.New())
}

package generic_test

import (
	"fmt"
	"testing"

	"github.com/kkqy-go/generic"
)

func TestSyncMap(t *testing.T) {
	p := generic.SyncMap[int, int]{}
	key := 1
	p.Store(key, 1)
	fmt.Println(p.Load(key))
}

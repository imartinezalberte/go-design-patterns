// singleton tasks
//
// When no Counter has been created before, a new one is created with the value 0
// If a Counter has already been created, return this instance that holds the actual count
// If we call the method Inc, the count must be incremented by 1
package creational

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

const CounterHandlerFormat = "Counter value is: %d"

type Counter int64

var c *Counter

func CounterSingleton() *Counter {
	if c == nil {
		c = new(Counter)
	}
	return c
}

func (c *Counter) Inc() *Counter {
	var i int64 = atomic.AddInt64((*int64)(c), 1)
	return (*Counter)(&i)
}

func (c *Counter) Reset() {
	*c = 0
}

// Handler
func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, CounterHandlerFormat, *(*int64)(c.Inc()))
}

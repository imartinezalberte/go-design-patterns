package creational_test

import (
	"sync"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/imartinezalberte/go-design-patterns/creational"
)

var _ = Describe("Singleton", func() {
	When("Creating singleton counter a handful of times, the same pointer is returned", func() {
		It("just works", func() {
			tmp := creational.CounterSingleton()
			for i := 0; i < 100; i++ {
				c := creational.CounterSingleton()

				Ω(tmp).To(Equal(c))
			}
		})
	})

	When("Adding one multiple times, value is updated", func() {
		var c *creational.Counter

		BeforeEach(func() {
			c = creational.CounterSingleton()
			c.Reset()
		})

		It("Adding one each time we call the Add method", func() {
			var i int64 = 1
			for ; i < 100; i++ {
				Ω(*(*int64)(c.Inc())).To(Equal(i))
			}
		})
	})

	When("Adding one multiple times using goroutines, value is updated", func() {
		var c *creational.Counter

		BeforeEach(func() {
			c = creational.CounterSingleton()
			c.Reset()
		})

		It("Adding one each time we call the Add method using goroutines", func() {
			var (
				max int64 = 100
				i   int64
				wg  sync.WaitGroup
			)

			for ; i < max; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.Inc()
				}()
			}

			wg.Wait()

			Ω(*(*int64)(c)).To(Equal(max))
		})
	})
})

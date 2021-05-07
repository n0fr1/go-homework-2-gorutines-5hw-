//3. Протестируйте производительность операций чтения и записи на множестве
//действительных чисел, безопасность которого обеспечивается sync.Mutex и
//sync.RWMutex для разных вариантов использования: 10% запись, 90% чтение; 50%
//запись, 50% чтение; 90% запись, 10% чтение

package main

import (
	"math/rand"
	"sync"
	"testing"
)

type FloatSet struct {
	sync.Mutex
	sync.RWMutex
	mm map[float64]struct{}
}

func NewFloatSet() *FloatSet {
	return &FloatSet{
		mm: map[float64]struct{}{},
	}
}

//(write) - add to map
func (s *FloatSet) Add(i float64, mutexR bool) {

	if mutexR {
		s.Mutex.Lock()
		s.mm[i] = struct{}{}
		s.Mutex.Unlock()
	} else {
		s.RWMutex.Lock()
		s.mm[i] = struct{}{}
		s.RWMutex.Unlock()
	}

}

//(read) - already in map
func (s *FloatSet) Has(i float64, mutexR bool) bool {

	var ok bool

	if mutexR {
		s.Mutex.Lock()
		_, ok = s.mm[i]
		s.Mutex.Unlock()
	} else {
		s.RWMutex.RLock()
		_, ok = s.mm[i]
		s.RWMutex.RUnlock()
	}

	return ok
}

//mutex 10w90r
func BenchmarkFloatSet10w90r(b *testing.B) {

	var set = NewFloatSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 { //random till 10
					set.Add(1, true)
				} else {
					set.Has(1, true)
				}
			}
		})
	})

}

//mutex 50w50r
func BenchmarkFloatSet50w50r(b *testing.B) {

	var set = NewFloatSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(2) == 1 {
					set.Add(1, true)
				} else {
					set.Has(1, true)
				}
			}
		})
	})
}

//mutex 90w10r
func BenchmarkFloatSet90w10r(b *testing.B) {

	var set = NewFloatSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Has(1, true)
				} else {
					set.Add(1, true)
				}
			}
		})
	})
}

//mutexRW 10w90r
func BenchmarkFloatSetRW10w90r(b *testing.B) {

	var set = NewFloatSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Add(1, false)
				} else {
					set.Has(1, false)
				}
			}
		})
	})
}

//mutexRW 50w50r
func BenchmarkFloatSetRW50w50r(b *testing.B) {

	var set = NewFloatSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(2) == 1 {
					set.Add(1, false)
				} else {
					set.Has(1, false)
				}
			}
		})
	})
}

//mutexRW 90w10r
func BenchmarkFloatSetRW90w10r(b *testing.B) {

	var set = NewFloatSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Intn(10) == 1 {
					set.Has(1, false)
				} else {
					set.Add(1, false)
				}
			}
		})
	})
}

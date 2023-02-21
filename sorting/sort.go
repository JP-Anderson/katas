package sorting

import (
	"fmt"
	"math"
)

type sortedMap struct {
	counts []countArrayElement
	indexes map[int]int
}

func (s *sortedMap) Count(n int) error {
	var ix int
	var ok bool
	if ix, ok = s.indexes[n]; !ok {
		return fmt.Errorf("n not in range")
	}
	old := s.counts[ix]
	s.counts[ix] = countArrayElement{
		value: old.value,
		count: old.count+1,
	}
	return nil
}

type countArrayElement struct {
	value int
	count int
}

func CounterDictionaryForRange(start, length int) *sortedMap {
	arr := make([]countArrayElement, length)
	index := make(map[int]int, length)
	for i := 0; i < length; i++ {
		arr[i] = countArrayElement{
			value: start+i,
			count: 0,
		}
		index[start+i] = i
	}
	return &sortedMap{
		counts: arr,
		indexes: index,
	}
}

func Sort(n []int) []int {
	min := math.MaxInt64
	max := math.MinInt64
	for _, num := range n {
		min = int(math.Min(float64(num), float64(min)))
		max = int(math.Max(float64(num), float64(max)))
	}
	fmt.Println(max, min)
	d := CounterDictionaryForRange(min, max-min+1)
	for _, num := range n {
		d.Count(num)
	}
	sorted := make([]int, len(n))
	i := 0
	for _, el := range d.counts {
		for j := 0; j < el.count; j++ {
			sorted[i] = el.value
			i++
		}
	}
	return sorted
}

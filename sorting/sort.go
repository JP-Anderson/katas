package sorting

import (
	"fmt"
	"math"
)

type sortedMap struct {
	counts []countArrayElement
	start, rnge int
}

func index(n, start, rnge int) (int, error) {
	if n < start {
		return -1, fmt.Errorf("N %d is less than start %d", n, start)
	}
	if n >= start+rnge {
		return -1, fmt.Errorf("N %d exceeds range end %d", n, start+rnge)
	}
	return n-start, nil
}

func (s *sortedMap) Count(n int) error {
	ix, err := index(n, s.start, s.rnge)
	if err != nil {
		fmt.Println(err)	
		return err
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
	for i := 0; i < length; i++ {
		arr[i] = countArrayElement{
			value: start+i,
			count: 0,
		}
	}
	return &sortedMap{
		counts: arr,
		start: start,
		rnge: length,
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

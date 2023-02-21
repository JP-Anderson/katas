package sorting

import (
	"fmt"
	"math"
	"strings"
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
	d := CounterDictionaryForRange(min, max-min+1)
	for _, num := range n {
		d.Count(num)
	}
	sorted := make([]int, len(n))
	c := make(chan int, 1)
	go inOrder(d, c)
	for i, _ := range n {
		sorted[i] = <- c	
	}
	return sorted
}

func SortStr(s string) string {
	min := math.MaxInt64
	max := math.MinInt64
	for _, c := range s {
		min = int(math.Min(float64(c), float64(min)))
		max = int(math.Max(float64(c), float64(max)))
	}
	d := CounterDictionaryForRange(min, max-min+1)
	for _, c := range s {
		d.Count(int(c))
	}
	var sb strings.Builder
	c := make(chan int, 1)
	go inOrder(d, c)
	for _ = range s {
		sb.WriteRune(rune( <- c))
	}
	return sb.String()
}

func inOrder(countMap *sortedMap, out chan<- int) {
	for _, el := range countMap.counts {
		for i := 0; i < el.count; i++ {
			out <- el.value
		}
	}
}


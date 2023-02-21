package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	input := []int{ 4, 3, 2, 1 }
	assert.Equal(t, []int{1, 2, 3, 4}, Sort(input))
}

func TestSort100000(t *testing.T) {
	input := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		input[i] = 100000-i
	}
	sorted := Sort(input)
	prev := -1
	for _, n := range sorted {
		assert.True(t, prev<n)
		prev = n
	}
}

func TestIndex(t *testing.T) {
	// boundary
	ix, err := index(0, 0, 10)
	assert.Equal(t, 0, ix)

	ix, err = index(1,0,10)
	assert.Equal(t, 1, ix)
	assert.Nil(t, err)

	ix, err = index(9,0,10)
	assert.Equal(t, 9, ix)
	assert.Nil(t, err)

	ix, err = index(-25, -25, 10)
	assert.Equal(t, 0, ix)
	assert.Nil(t, err)

	ix, err = index(-23, -25, 10)
	assert.Equal(t, 2, ix)
	assert.Nil(t, err)

	ix, err = index(-16, -25, 10)
	assert.Equal(t, 9, ix)
	assert.Nil(t, err)
}

func TestIndexErrors(t *testing.T) {
	ix, err := index(-1,0,10)
	assert.Equal(t, -1, ix)
	assert.NotNil(t, err)

	ix, err = index(10, 0, 10)
        assert.Equal(t, -1, ix)
        assert.NotNil(t, err)

	ix, err = index(-26, -25, 10)
	assert.Equal(t, -1, ix)
	assert.NotNil(t, err)
	
	ix, err = index(-15, -25, 10)
	assert.Equal(t, -1, ix)
	assert.NotNil(t, err)
}

func TestCountArray(t *testing.T) {
	sortMap := CounterDictionaryForRange(0, 10)
	assert.Len(t, sortMap.counts, 10)
	assert.Equal(t, countArrayElement{ value: 0, count: 0}, sortMap.counts[0])
	assert.Equal(t, countArrayElement{ value: 1, count: 0}, sortMap.counts[1])
	assert.Equal(t, countArrayElement{ value: 2, count: 0}, sortMap.counts[2])
	assert.Equal(t, countArrayElement{ value: 3, count: 0}, sortMap.counts[3])
	assert.Equal(t, countArrayElement{ value: 4, count: 0}, sortMap.counts[4])
	assert.Equal(t, countArrayElement{ value: 5, count: 0}, sortMap.counts[5])
	assert.Equal(t, countArrayElement{ value: 6, count: 0}, sortMap.counts[6])
	assert.Equal(t, countArrayElement{ value: 7, count: 0}, sortMap.counts[7])
	assert.Equal(t, countArrayElement{ value: 8, count: 0}, sortMap.counts[8])
	assert.Equal(t, countArrayElement{ value: 9, count: 0}, sortMap.counts[9])
}

package challenges

import (
	"sync"
)

func ChannelPairSum(ascendingChan, descendingChan chan int, target int64, size int) uint32 {
	asc := make([]int, size)
	desc := make([]int, size)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < size; i++ {
			asc[i] = <-ascendingChan
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < size; i++ {
			desc[i] = <-descendingChan
		}
		wg.Done()
	}()

	wg.Wait()

	var count uint32
	i, j := 0, 0

	for i < len(asc) && j < len(desc) {
		sum := int64(asc[i] + desc[j])
		switch {
		case sum == target:
			count++
			i++
			j++
		case sum < target:
			i++
		default:
			j++
		}
	}
	return count
}

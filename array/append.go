package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	a := []uint64{1, 2, 3, 7, 8}

	aa, err := removeSorted(a, 8)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
	fmt.Println(aa)
}

// RemoveSorted removes
func removeSorted(a []uint64, v uint64) ([]uint64, error) {
	if len(a) == 0 {
		return nil, errors.New("archived list is empty")
	}

	// Check if given version is archived
	i, ok := search(a, v)
	if ok != true || a[i] != v {
		return nil, errors.Errorf("version %v is not archived", v)
	}

	res := append(a[:i], a[i+1:]...)
	return res, nil
}

// InsertSorted inserts
func insertSorted(a []uint64, v uint64) ([]uint64, error) {
	if len(a) == 0 {
		return []uint64{v}, nil
	}

	// Check if given v already in slice a
	i, ok := search(a, v)
	if ok == true && a[i] == v {
		return nil, errors.Errorf("version %v is already in archived list", v)
	}

	if ok == false {
		//insert at the beginning
		return append([]uint64{v}, a...), nil
	}
	res := append(a, 0)
	copy(res[i+2:], res[i+1:])
	res[i+1] = v

	return res, nil
}

func search(a []uint64, x uint64) (int, bool) {
	if len(a) == 0 {
		return 0, false
	}

	min := 0
	if a[min] > x {
		return 0, false
	}
	if a[min] == x {
		return min, true
	}

	max := len(a) - 1
	if a[max] <= x {
		return max, true
	}
	if uint64(max) > x {
		max = int(x)
		if a[max] == x {
			return max, true
		}
	}

	for max-min > 1 {
		mid := min + int((max-min)/2)
		if a[mid] == x {
			return mid, true
		}
		if a[mid] < x {
			min = mid
			continue
		}
		if a[mid] > x {
			max = mid
			continue
		}
	}

	return min, true
}

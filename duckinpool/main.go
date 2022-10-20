package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	times = 10000000
)

func deal(count int) float64 {
	if count <= 1 {
		return 1
	}

	arr := make([]float64, count+1)

	var in int
	for i := 0; i < times; i++ {
		for i := 0; i < count; i++ {
			arr[i] = float64(rand.Int63n(360)) + rand.Float64()
		}

		sort.Float64s(arr)
		arr[count] = arr[0] + 360

		for i := 0; i < count; i++ {
			distance := arr[i+1] - arr[i]
			if distance > 180 {
				in += 1
			}
		}
	}

	return float64(in) / times
}

func main() {
	rand.Seed(time.Now().Unix())
	for n := 1; n < 15; n++ {
		fmt.Printf("duck count %v result %v\n", n, deal(n))
	}
}

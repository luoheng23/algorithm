package chapter15

import "fmt"

// Bag01 solve bag 01 problem
func Bag01(weight []int, maxWeight int) int {
	if maxWeight <= 0 {
		return 0
	}
	M := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		M[i] = make([]int, maxWeight+1)
	}
	for w := 0; w < maxWeight+1; w++ {
		if w >= weight[0] {
			M[0][w] += weight[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		for w := 0; w < maxWeight+1; w++ {
			if weight[i] > w || weight[i]+M[i-1][w-weight[i]] < M[i-1][w] {
				M[i][w] = M[i-1][w]
			} else {
				M[i][w] = weight[i] + M[i-1][w-weight[i]]
			}
		}
	}
	return M[len(weight)-1][maxWeight]
}

// BagNormal solve bag normal problem
func BagNormal(weight []int, value []int, maxWeight int) int {
	if maxWeight <= 0 || len(weight) != len(value) {
		return 0
	}
	M := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		M[i] = make([]int, maxWeight+1)
	}
	for w := 0; w < maxWeight+1; w++ {
		if w >= weight[0] {
			M[0][w] += value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		for w := 0; w < maxWeight+1; w++ {
			if weight[i] > w || weight[i]+M[i-1][w-weight[i]] < M[i-1][w] {
				M[i][w] = M[i-1][w]
			} else {
				M[i][w] = value[i] + M[i-1][w-weight[i]]
			}
		}
	}
	return M[len(weight)-1][maxWeight]
}

// weightValue used in bagfraction
type WeightValue struct {
	weight int
	value  int
	wv     float64
}

func partition(value []WeightValue, p, r int, maxWeight int) (int, float64, int) {
	if r <= p+1 {
		return 0, 0, 0
	}
	t := value[r-1].wv
	i := p
	sumW, sumV := 0, 0
	for j := p; j < r-1; j++ {
		if value[j].wv > t {
			value[i], value[j] = value[j], value[i]
			sumW += value[i].weight
			sumV += value[i].value
			i++
		}
	}
	if i != r-1 {
		value[i], value[r-1] = value[r-1], value[i]
	}
	if sumW <= maxWeight {
		return sumW, float64(sumV), i
	}
	return -1, -1, i
}

func BagFraction(value []WeightValue, p, r int, maxWeight int) float64 {
	if maxWeight <= 0 {
		return 0
	}
	if r > p+1 {
		sumW, sumV, q := partition(value, p, r, maxWeight)
		fmt.Print(sumW, sumV, q, value)
		if sumW == -1 {
			return BagFraction(value, p, q, maxWeight)
		}
		return sumV + BagFraction(value, q, r, maxWeight-sumW)
	} else if r == p+1 {
		if value[p].weight >= maxWeight {
			return value[p].wv * float64(maxWeight)
		}
	}
	return 0
}

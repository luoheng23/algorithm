package chapter15

// Bag01 solve bag 01 problem
func Bag01(weight []int, maxWeight int) int {
	if maxWeight <= 0 {
		return 0
	}
	M := make([][]int, len(weight)*(maxWeight+1))
	for i := 1; i < len(weight); i++ {
		for w := 0; w < maxWeight+1; w++ {
			if weight[i]+M[i-1][w-weight[i]] > M[i-1][w] {
				M[i][w] = weight[i] + M[i-1][w-weight[i]]
			} else {
				M[i][w] = M[i-1][w]
			}
		}
	}
	return M[len(weight)-1][maxWeight]
}

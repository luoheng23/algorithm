package chapter1

// 有n根棍子，从中选出3根棍子组成周长尽可能长的三角形，输出最大周长
func solve(n int, a []int) int {
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if a[i]+a[j] > a[k] && a[k]+a[i] > a[j] && a[j]+a[k] > a[i] && a[i]+a[j]+a[k] > ans {
					ans = a[i] + a[j] + a[k]
				}
			}
		}
	}
	return ans
}

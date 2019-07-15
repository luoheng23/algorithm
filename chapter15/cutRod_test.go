package chapter15

import "testing"

func TestCutRodSlow(t *testing.T) {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	s := CutRodSlow(p, 4)
	if s != 10 {
		t.Errorf("CutRodSlow failed. Expected 10, Got %d", s)
	}
}

// notice:
// this benchmark is extremely slow, please don't run it
func BenchmarkCutRodSlow(b *testing.B) {
	const LENGTH = 30
	p := make([]int, LENGTH+1)
	for i := 0; i <= LENGTH; i++ {
		p[i] = i
	}
	for i := 0; i < b.N; i++ {
		CutRodSlow(p, LENGTH)
	}
}

func TestCutRodFast(t *testing.T) {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	s := CutRodMemorizedFast(p, 4)
	s2 := CutRodBottomFast(p, 4)
	if s != 10 || s2 != 10 {
		t.Errorf("CutRodFast failed. Expected (10, 10), Got (%d, %d)", s, s2)
	}
}

func BenchmarkCutRodFast(b *testing.B) {
	const LENGTH = 30
	p := make([]int, LENGTH+1)
	for i := 0; i <= LENGTH; i++ {
		p[i] = i
	}
	for i := 0; i < b.N; i++ {
		CutRodBottomFast(p, LENGTH)
	}
}

func TestCutRodSolution(t *testing.T) {
	p := []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	s := CutRodSolution(p, 10)
	s2 := CutRodSolution(p, 7)
	if s[0] != 10 || s2[0] != 1 || s2[1] != 6 {
		t.Errorf("CutRodFast failed. Expected (10, [1, 6]), Got (%d, %v)", s[0], s2)
	}
}
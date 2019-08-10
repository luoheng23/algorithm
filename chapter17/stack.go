package chapter17

// Stack is stack
type Stack struct {
	A   []int
	top int
}

// Init create a stack
func Init() Stack {
	return Stack{top: -1}
}

func pop(S Stack) int {
	if S.top > -1 {
		S.top--
		return S.A[S.top+1]
	}
}

func push(S Stack, x int) {
	if S.top < len(S.A) {
		S.top++
		S.A[S.top] = x
	}
}

func multipop(S Stack, k int) {
	for S.top != -1 && k > 0 {
		pop(S)
		k--
	}
}

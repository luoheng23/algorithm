package chapter17

// Stack is stack
type Stack struct {
	A []int
	top int
}

// Init create a stack
func Init() Stack {
	S = Stack{top: -1}
}

func pop(S Stack) {
	if S.top > -1 {
		top--
		return S.A[top+1]
	}
}

func push(S Stack, x int) {
	if S.top < len(S.A) {
		top++
		S.A[top] = x
	}
}

func multipop(S Stack, k int) {
	for S.top != -1 && k > 0 {
		pop(S)
		k--
	}
}
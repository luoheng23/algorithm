package chapter16

import "testing"

func TestHuffman(t *testing.T) {
	c := []*Node {
		&Node{value: 'a', freq: 45},
		&Node{value: 'b', freq: 13},
		&Node{value: 'c', freq: 12},
		&Node{value: 'd', freq: 16},
		&Node{value: 'e', freq: 9},
		&Node{value: 'f', freq: 5},
	}
	codeTree := Huffman(c)
	if codeTree.root.freq != 100 || codeTree.root.left.value != 'a' || codeTree.root.left.freq != 45 ||
	codeTree.root.right.freq != 55 || codeTree.root.right.left.freq != 25 || codeTree.root.right.right.freq != 30 ||
	codeTree.root.right.right.left.freq != 14 {
		t.Errorf("Huffman failed.")
	}
}
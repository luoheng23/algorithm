package chapter15

import "testing"

func TestBag01(t *testing.T) {
	bagWeight := [5]int{4, 4, 7, 6, 5}
	maxWeight := 8
	s := Bag01(bagWeight[:], maxWeight)
	if s != 8 {
		t.Errorf("Test Bag01 failed. Expected 8, Got %d", s)
	}
}

func TestBag01Less(t *testing.T) {
	bagWeight := []int{9, 9, 3, 7, 9}
	maxWeight := 26
	s := Bag01Less(bagWeight, maxWeight)
	if s !=  25{
		t.Errorf("Test Bag01 failed. Expected 8, Got %d", s)
	}
}

func TestBagNormal(t *testing.T) {
	bagWeight := [5]int{4, 4, 7, 6, 5}
	bagValue := [5]int{5, 5, 9, 6, 5}
	maxWeight := 8
	s := BagNormal(bagWeight[:], bagValue[:], maxWeight)
	if s != 10 {
		t.Errorf("Test Bag01 failed. Expected 10, Got %d", s)
	}
}

func TestBagNormalLess(t *testing.T) {
	bagWeight := [5]int{4, 4, 7, 6, 5}
	bagValue := [5]int{5, 5, 9, 6, 5}
	maxWeight := 8
	s := BagNormalLess(bagWeight[:], bagValue[:], maxWeight)
	if s != 10 {
		t.Errorf("Test Bag01 failed. Expected 10, Got %d", s)
	}
}

func TestBagFraction(t *testing.T) {
	bagWeight := [5]int{4, 2, 3, 6, 5}
	bagValue := [5]int{4, 9, 9, 2, 3}
	maxWeight := 10
	value := make([]WeightValue, 5)
	for i := 0; i < 5; i++ {
		value[i].weight, value[i].value = bagWeight[i], bagValue[i]
		value[i].wv = float64(value[i].value) / float64(value[i].weight)
	}
	s := BagFraction(value, 0, len(value), maxWeight)
	if s != 22.6 {
		t.Errorf("Test Bag01 failed. Expected 22.6, Got %f", s)
	}
}

func BenchmarkBagFraction(b *testing.B) {
	bagWeight := [5]int{4, 2, 3, 6, 5}
	bagValue := [5]int{4, 9, 9, 2, 3}
	maxWeight := 10
	value := make([]WeightValue, 5)
	for i := 0; i < 5; i++ {
		value[i].weight, value[i].value = bagWeight[i], bagValue[i]
		value[i].wv = float64(value[i].value) / float64(value[i].weight)
	}
	for i := 0; i < b.N; i++ {
		BagFraction(value, 0, len(value), maxWeight)
	}
}
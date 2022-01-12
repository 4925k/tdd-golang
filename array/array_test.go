package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d got %d", want, got)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{1, 1})
	want := []int{3, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of tail of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 4})
		want := []int{5, 4}
		checkSum(t, got, want)
	})

	t.Run("sum of tails including empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}
		checkSum(t, got, want)
	})
}

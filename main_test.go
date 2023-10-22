package main

import "testing"

func TestActivateGiftCard(t *testing.T) {
	t.Run("closure_with_different_func", func(t *testing.T) {
		useGiftCard1 := ActivateGiftCard()
		useGiftCard2 := ActivateGiftCard()
		result1 := useGiftCard1(10)
		result2 := useGiftCard2(5)
		if result1 != 90 {
			t.Errorf("useGiftCard1(10) = %v, want %v", result1, 90)
		}
		if result2 != 95 {
			t.Errorf("useGiftCard2(5) = %v, want %v", result1, 95)
		}
	})
}

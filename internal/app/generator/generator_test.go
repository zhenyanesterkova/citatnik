package generator

import (
	"testing"
)

func Test_GetQuoteID(t *testing.T) {

}

func TestGenerator(t *testing.T) {
	wantNew := &Generator{
		quoteIDCounter: 0,
	}
	tt := New()

	t.Run("success_New", func(t *testing.T) {
		if tt.quoteIDCounter != wantNew.quoteIDCounter {
			t.Errorf("Generator.quoteIDCounter = %v, want %v", tt.quoteIDCounter, wantNew.quoteIDCounter)
		}
	})
	t.Run("success_GetQuoteID", func(t *testing.T) {
		if got := tt.GetQuoteID(); got != 1 {
			t.Errorf("Generator.GetQuoteID() = %v, want %v", got, 1)
		}
	})
}

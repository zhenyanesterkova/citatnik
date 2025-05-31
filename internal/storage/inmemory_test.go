package storage

import (
	"errors"
	"reflect"
	"sync"
	"testing"

	"github.com/zhenyanesterkova/citatnik/internal/app/apperrors"
	"github.com/zhenyanesterkova/citatnik/internal/app/generator"
	"github.com/zhenyanesterkova/citatnik/internal/app/quote"
)

func TestNew(t *testing.T) {
	want := &InMemory{
		quotes:      []*quote.Quote{},
		idIndex:     make(map[uint64]*quote.Quote),
		authorIndex: make(map[string][]*quote.Quote),
		mu:          sync.RWMutex{},
		generator:   generator.New(),
	}
	t.Run("success", func(t *testing.T) {
		if got := New(); !reflect.DeepEqual(got, want) {
			t.Errorf("New() = %v, want %v", got, want)
		}
	})
}

func TestInMemory_Add(t *testing.T) {
	quote1 := &quote.Quote{
		Author: "test",
		Text:   "Test Text",
	}

	wantStore := New()
	wantStore.quotes = append(wantStore.quotes, quote1)
	wantStore.authorIndex["test"] = append(wantStore.authorIndex["test"], quote1)
	wantStore.idIndex[1] = quote1
	wantStore.generator.GetQuoteID()

	m := New()

	t.Run("success", func(t *testing.T) {
		err := m.Add(quote1)
		if err != nil {
			t.Errorf("InMemory.Add() error = %v, wantErr %v", err, nil)
		}

		if len(wantStore.quotes) != len(m.quotes) {
			t.Errorf("len(InMemory.quotes) = %v, want %v", m.quotes, len(wantStore.quotes))
		}
		if len(wantStore.authorIndex) != len(m.authorIndex) {
			t.Errorf("len(InMemory.authorIndex) = %v, want %v", len(m.authorIndex), len(wantStore.authorIndex))
		}
		if len(wantStore.idIndex) != len(m.idIndex) {
			t.Errorf("len(InMemory.idIndex) = %v, want %v", len(m.idIndex), wantStore.idIndex)
		}
		if !reflect.DeepEqual(wantStore.quotes[0], m.quotes[0]) {
			t.Errorf("InMemory.quotes = %v, want %v", m.quotes, wantStore.quotes)
		}
		if !reflect.DeepEqual(wantStore.authorIndex["test"], m.authorIndex["test"]) {
			t.Errorf("InMemory.authorIndex = %v, want %v", m.authorIndex, wantStore.authorIndex)
		}
		if !reflect.DeepEqual(wantStore.idIndex[1], m.idIndex[1]) {
			t.Errorf("InMemory.idIndex = %v, want %v", m.idIndex, wantStore.idIndex)
		}
	})

	quote2 := &quote.Quote{
		Author: "test",
		Text:   "Test Text",
		ID:     2,
	}
	m.quotes = append(m.quotes, quote2)
	m.authorIndex["test"] = append(m.authorIndex["test"], quote2)
	m.idIndex[2] = quote2

	t.Run("error", func(t *testing.T) {
		err := m.Add(quote2)
		if err == nil || err.Error() != "failed add quote to memory: ID 2 already exists" {
			t.Errorf("InMemory.Add() error = %v, wantErr %v", err, "failed add quote to memory: ID 2 already exists")
		}
	})

}

func TestInMemory_GetAll(t *testing.T) {
	quote1 := &quote.Quote{
		Author: "test",
		Text:   "Test Text",
	}

	want := []*quote.Quote{}
	want = append(want, quote1)

	m := New()
	_ = m.Add(quote1)

	t.Run("success", func(t *testing.T) {
		if got := m.GetAll(); !reflect.DeepEqual(got, want) {
			t.Errorf("InMemory.GetAll() = %v, want %v", got, want)
		}
	})
}

func TestInMemory_GetRandom(t *testing.T) {
	quote1 := &quote.Quote{
		Author: "test",
		Text:   "Test Text",
	}

	m := New()

	t.Run("nil", func(t *testing.T) {
		if got := m.GetRandom(); got != nil {
			t.Errorf("InMemory.GetRandom() = %v, want %v", got, nil)
		}
	})

	_ = m.Add(quote1)

	t.Run("success", func(t *testing.T) {
		if got := m.GetRandom(); !reflect.DeepEqual(got, quote1) {
			t.Errorf("InMemory.GetRandom() = %v, want %v", got, quote1)
		}
	})
}

func TestInMemory_GetByAuthor(t *testing.T) {
	quote1 := &quote.Quote{
		Author: "test",
		Text:   "Test Text",
	}

	want := make(map[string][]*quote.Quote)
	want["test"] = append(want["test"], quote1)

	m := New()
	_ = m.Add(quote1)

	t.Run("success", func(t *testing.T) {
		if got := m.GetByAuthor("test"); !reflect.DeepEqual(got[0], quote1) {
			t.Errorf("InMemory.GetByAuthor() = %v, want %v", got[0], quote1)
		}
	})
}

func TestInMemory_Delete(t *testing.T) {
	quote1 := &quote.Quote{
		Author: "test",
		Text:   "Test Text",
	}

	m := New()
	_ = m.Add(quote1)

	t.Run("not found", func(t *testing.T) {
		err := m.Delete(3)
		if err == nil || !errors.Is(err, apperrors.ErrDeleteNotFound) {
			t.Errorf("InMemory.Delete() error = %v, wantErr %v", err, apperrors.ErrDeleteNotFound)
		}
	})

	t.Run("success", func(t *testing.T) {
		err := m.Delete(1)
		if err != nil {
			t.Errorf("InMemory.Delete() error = %v, wantErr %v", err, nil)
		}
		w := make(map[string][]*quote.Quote)
		w["test"] = []*quote.Quote{}
		if !reflect.DeepEqual(m.authorIndex, w) {
			t.Errorf("InMemory.authorIndex = %v, want %v", m.authorIndex, make(map[string][]*quote.Quote))
		}
		if !reflect.DeepEqual(m.quotes, []*quote.Quote{}) {
			t.Errorf("InMemory.quotes = %v, want %v", m.quotes, []*quote.Quote{})
		}
		if !reflect.DeepEqual(m.idIndex, make(map[uint64]*quote.Quote)) {
			t.Errorf("InMemory.idIndex = %v, want %v", m.idIndex, make(map[uint64]*quote.Quote))
		}
	})

}

func TestInMemory_Ping(t *testing.T) {
	m := New()
	err := m.Ping()
	if err != nil {
		t.Errorf("InMemory.Ping = %v, want %v", err, nil)
	}
}

package storage

import (
	"fmt"
	"log"
	"math/rand"
	"slices"
	"sync"
	"time"

	"github.com/zhenyanesterkova/citatnik/internal/app/apperrors"
	"github.com/zhenyanesterkova/citatnik/internal/app/generator"
	"github.com/zhenyanesterkova/citatnik/internal/app/quote"
)

type Generator interface {
	GetQuoteID() uint64
}

type InMemory struct {
	quotes []*quote.Quote

	idIndex     map[uint64]*quote.Quote
	authorIndex map[string][]*quote.Quote

	generator Generator

	mu sync.RWMutex
}

func New() *InMemory {
	gen := generator.New()
	return &InMemory{
		quotes:      []*quote.Quote{},
		idIndex:     make(map[uint64]*quote.Quote),
		authorIndex: make(map[string][]*quote.Quote),
		mu:          sync.RWMutex{},
		generator:   gen,
	}
}

func (m *InMemory) Add(quote *quote.Quote) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	quote.ID = m.generator.GetQuoteID()
	quote.CreatedAt = time.Now()
	quote.UpdatedAt = time.Now()

	if _, ok := m.idIndex[quote.ID]; ok {
		log.Printf("quote with ID %d already exists; check the quote id generator", quote.ID)
		return fmt.Errorf("failed add quote to memory: ID %d already exists", quote.ID)
	}

	m.quotes = append(m.quotes, quote)
	m.authorIndex[quote.Author] = append(m.authorIndex[quote.Author], quote)
	m.idIndex[quote.ID] = quote

	return nil
}

func (m *InMemory) GetAll() []*quote.Quote {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.quotes
}

func (m *InMemory) GetRandom() *quote.Quote {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.quotes) == 0 {
		return nil
	}

	return m.quotes[rand.Intn(len(m.quotes))]
}

func (m *InMemory) GetByAuthor(author string) []*quote.Quote {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.authorIndex[author]
}

func (m *InMemory) Delete(id uint64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	quote, ok := m.idIndex[id]
	if !ok {
		return fmt.Errorf("failed delete quote with ID %d: %w", id, apperrors.ErrDeleteNotFound)
	}

	delete(m.idIndex, id)

	for idx, quote := range m.authorIndex[quote.Author] {
		if quote.ID == id {
			m.authorIndex[quote.Author] = slices.Delete(m.authorIndex[quote.Author], idx, idx+1)
			break
		}
	}

	for idx, quote := range m.quotes {
		if quote.ID == id {
			m.quotes = slices.Delete(m.quotes, idx, idx+1)
			break
		}
	}

	return nil
}

func (m *InMemory) Ping() error {
	return nil
}

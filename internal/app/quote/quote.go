package quote

import "time"

type Quote struct {
	ID        uint64    `json:"id"`
	Author    string    `json:"author"`
	Text      string    `json:"quote"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package quote

import "time"

type Quote struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    string    `json:"author"`
	Text      string    `json:"quote"`
	ID        uint64    `json:"id"`
}

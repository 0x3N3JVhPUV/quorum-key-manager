package types

import "time"

// Metadata are generated by the store attached to stored items
type Metadata struct {
	Version int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	PurgeAt   time.Time
}

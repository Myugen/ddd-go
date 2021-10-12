package tavern

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a value object
type Transaction struct {
	// all values are in lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}

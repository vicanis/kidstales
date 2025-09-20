package sqlite

import (
	"context"
	"log"
	"time"
)

const cleanerInterval = time.Minute

func StartCleaner(ctx context.Context) error {
	ticker := time.NewTicker(cleanerInterval)

	for {
		select {
		case <-ctx.Done():
			return nil

		case <-ticker.C:
			err := NewDBCache().Clean()
			if err != nil {
				log.Printf("cleaner failed: %v", err)
			}
		}
	}
}

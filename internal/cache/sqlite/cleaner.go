package sqlite

import (
	"context"
	"log"
	"time"
)

const cleanerInterval = time.Minute

func StartCleaner(ctx context.Context) (err error) {
	log.Printf("cleaner: start")

	ticker := time.NewTicker(cleanerInterval)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			break loop

		case <-ticker.C:
			err = NewDBCache().Clean()
			if err != nil {
				break loop
			}
		}
	}

	log.Printf("cleaner: end: %v", err)

	return
}

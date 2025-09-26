package link

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vky5/logcast/internals/filehandler"
	"github.com/vky5/logcast/internals/utils"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// Initialize Redis Cache
func InitRedis(addr, password string, db int) {
	rdb = redis.NewClient(&redis.Options{ // rdb = redis database backup
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// generateSlug creates a random short slug
func generateSlug(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", utils.FailedOnError(err, "[Link]", "Failed to generate the slug")
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// build unique URL for every file
func BuildURL(fs filehandler.FileSet) (string, error) {
	prefuxURL := os.Getenv("PREFIX_URL") // getting the prefix of the URL
	slug, err := generateSlug(6)
	if err != nil {
		return "", utils.FailedOnError(err, "[Link]", "Failed to generate the slug")
	}

	ttl := time.Duration(utils.MustAtoi(os.Getenv("REDIS_TTL"))) * time.Minute

	// Store slug in Redis with TTL
	err = rdb.Set(ctx, slug, "active", ttl).Err() // rn just setting up that a particular slug is active or not 
	if err != nil {
		return "", utils.FailedOnError(err, "[link]", "Failed to write to redis server")
	}

	// Build full URL
	fullURL := fmt.Sprintf("%s/%s", prefuxURL, slug)
	return fullURL, nil
}

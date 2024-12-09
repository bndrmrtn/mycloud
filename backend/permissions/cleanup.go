package permissions

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	// First %s = user.ID, second %s = other ID

	FileReadFormat   string = "permission:file.read:%s-%s"
	FileDeleteFormat string = "permission:file.delete:%s-%s"
	FileUpdateFormat string = "permission:file.update:%s-%s"

	SpaceAccessFormat     string = "permission:space:%s-%s"
	SpaceFileUploadFormat string = "permission:space.create:%s-%s"
	SpaceReadFormat       string = "permission:space.read:%s-%s"
)

func CleanUp(rdb *redis.Client, userID, otherID string) {
	formats := []string{FileReadFormat, FileDeleteFormat, FileUpdateFormat, SpaceAccessFormat, SpaceFileUploadFormat, SpaceReadFormat}

	for _, format := range formats {
		key := fmt.Sprintf(format, userID, otherID)
		_ = rdb.Del(context.Background(), key).Err()
	}
}

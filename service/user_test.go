package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"tiktok/dao/mysql"
)

func init() {
	mysql.Init()
}

func TestGetFollowInfo(t *testing.T) {
	tests := []struct {
		id    int64
		other int64
		want  gin.H
	}{
		{1, 1, gin.H{"follow_count": 1, "follower_count": 0, "is_follow": true}},
		{1, 2, gin.H{"follow_count": 0, "follower_count": 1, "is_follow": true}},
		{0, 0, gin.H{"follow_count": 0, "follower_count": 0, "is_follow": true}},
	}
	h := make(map[string]any)

	for _, test := range tests {
		GetFollowInfo(test.id, test.other, h)
		for k, v := range test.want {
			switch v.(type) {
			case int64:
				if v.(int64) != h[k].(int64) {
					log.Fatalf("want: %v, got %v", test.want, h)
				}
			case bool:
				if v.(bool) != h[k].(bool) {
					log.Fatalf("want: %v, got %v", test.want, h)
				}
			}
		}
	}
}

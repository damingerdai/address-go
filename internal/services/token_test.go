package service

import (
	"testing"
	"time"
)

func TestTimeNow(t *testing.T) {
	now := time.Now()
	t.Log(now)
}

package redis

import "testing"

func TestHashSet(t *testing.T) {
	Rdb.HSet(Ctx, "test", "test")
}

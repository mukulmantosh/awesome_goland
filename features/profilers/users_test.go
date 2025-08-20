package main

import (
	"runtime"
	"testing"
	"time"
)

func TestMemoryUsersProfile(t *testing.T) {
	n := 50_000 // 50k users × 1 KB Data ≈ ~50 MB retained
	users := MakeUsers(n)

	t.Logf("created %d users; last ID=%d", len(users), users[len(users)-1].ID)

	// Keep slice alive until snapshot is taken
	runtime.KeepAlive(users)

	time.Sleep(2 * time.Second)
}

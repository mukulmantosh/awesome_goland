package main

import "testing"

func TestMutexBlocking(t *testing.T) {
	MutexBlocking(50)
}

func TestChannelBlocking(t *testing.T) {
	ChannelBlocking(50)
}

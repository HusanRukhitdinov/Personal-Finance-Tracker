package configs

import "time"

var SignKey = []byte("secret")

const (
	AccessExpireTime  = time.Minute * 120
	RefreshExpireTime = time.Hour * 240
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

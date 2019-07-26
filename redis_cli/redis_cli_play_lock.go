package redis_cli

import (
	rc "github.com/giskook/vav-common/redis_cli"
)

const (
	// KEYS[1] the key
	// KEYS[2] the status
	// KEYS[3] ttl
	// return 0 success 1 other is using
	SCRIPT_PLAY_LOCK string = `
	local key_ttl 
	local set_result
	local ttl_result
	key_ttl = redis.call("TTL", KEYS[1])
	if tonumber(key_ttl) < 0 then 
		redis.call("SET", KEYS[1], KEYS[2], "EX", KEYS[3])
		return 0
	else
		return 1
	end
	`
)

func SetPlayLock(key, status, ttl string) (int, error) {
	return rc.GetInstance().DoScript(SCRIPT_PLAY_LOCK, key, status, ttl)
}

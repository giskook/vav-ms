package redis_cli

import (
	rc "github.com/giskook/vav-common/redis_cli"
)

const (
	// KEYS[1] the key
	// KEYS[2] the data type sub key
	// KEYS[3] data type
	// KEYS[4] ttl key
	// KEYS[5] ttl
	SCRIPT_PLAY_INIT string = `
	local result
	redis.call("HMSET", KEYS[1], KEYS[2], KEYS[3], KEYS[4], KEYS[5])
	result = redis.call("EXPIRE", KEYS[1], KEYS[5])
	if tonumber(result) == 0 then 
		return 2
	end

	return 0
	`
	//KEYS[1] hash key
	//KEYS[2] hash uuid sub key
	//KEYS[3] hash uuid sub value
	//KEYS[4] the url sub key
	//KEYS[5] the url value
	//KEYS[6] hash ttl sub key
	//KEYS[7] the status field
	//return 1 not sub ttl exists
	//       2 expire set error
	//       0 success
	SCRIPT_DESTRUCT string = `local ttl
	local result_expire
	redis.call("HMSET", KEYS[1], KEYS[2], KEYS[3], KEYS[4], KEYS[5])
	ttl = redis.call("HGET", KEYS[1], KEYS[6])
	if ttl == "" then 
		return 1
	end
	result_expire = redis.call("EXPIRE", KEYS[1], ttl)
	if tonumber(result_expire) ~= 1 then 
		return 2
	end 
	redis.call("DEL", KEYS[7])

	return 0
	`
)

func StreamPlayInit(key, data_type_key, data_type, ttl_key, ttl string) (int, error) {
	return rc.GetInstance().DoScript(SCRIPT_PLAY_INIT, key, data_type_key, data_type, ttl_key, ttl)
}

func StreamPlayURL(key string) (string, error) {
	return rc.GetInstance().GetVehicleChan(key, VAVMS_STREAM_URL_KEY)
}

func StreamDestruct(key, key_uuid, uuid, key_url, url, key_ttl, status string) (int, error) {
	return rc.GetInstance().DoScript(SCRIPT_DESTRUCT, key, key_uuid, uuid, key_url, url, key_ttl, status)
}

func StreamGetTTL(key string) (string, error) {
	return rc.GetInstance().GetVehicleChan(key, VAVMS_STREAM_TTL_KEY)
}

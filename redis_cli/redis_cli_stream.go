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
	// KYES[6] priority key
	// KEYs[7] priority
	SCRIPT_PLAY_INIT string = ` local result
	redis.call("HMSET", KEYS[1], KEYS[2], KEYS[3], KEYS[4], KEYS[5], KEYS[6], KEYS[7])
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
	//KEYS[6] the http ip outter key
	//KEYS[7] the http ip outter
	//KEYS[8] the http port outter key
	//KEYS[9] the http port outter
	//KEYS[10] the http location key
	//KEYS[11] the http location
	//KEYS[12] the rtmp inner port key
	//KEYS[13] the rtmp inner port
	//KEYS[14] the rtmp application key
	//KEYS[15] the rtmp application
	//KEYS[16] hash ttl sub key
	//KEYS[17] the status field
	//return 1 not sub ttl exists
	//       2 expire set error
	//       0 success
	SCRIPT_DESTRUCT string = `local ttl
	local result_expire
	redis.call("HMSET", KEYS[1], KEYS[2], KEYS[3], KEYS[4], KEYS[5],KEYS[6],KEYS[7],KEYS[8],KEYS[9],KEYS[10],KEYS[11],KEYS[12],KEYS[13],KEYS[14],KEYS[15])
	ttl = redis.call("HGET", KEYS[1], KEYS[16])
	if ttl == "" then 
		return 1
	end
	result_expire = redis.call("EXPIRE", KEYS[1], ttl)
	if tonumber(result_expire) ~= 1 then 
		return 2
	end 
	redis.call("DEL", KEYS[17])

	return 0
	`

	// replace priority
	// KEYS[1] stream key
	// KEYS[2] priority sub key
	// KEYS[3] priority
	SCRIPT_SET_PRIORITY string = `local priority_redis
	priority_redis = redis.call("HGET", KEYS[1], KEYS[2])
	if tonumber(KEYS[3]) > tonumber(priority_redis) then
		redis.call("HSET", KEYS[1], KEYS[2], KEYS[3])
	end

	return 0 
	`
)

func StreamPlayInit(key, data_type_key, data_type, ttl_key, ttl, priority_key, priority string) (int, error) {
	return rc.GetInstance().DoScript(SCRIPT_PLAY_INIT, key, data_type_key, data_type, ttl_key, ttl, priority_key, priority)
}

func StreamPlayURL(key string) (string, error) {
	return rc.GetInstance().GetVehicleChan(key, VAVMS_STREAM_URL_KEY)
}

func StreamDestruct(key,
	key_uuid, uuid,
	key_url, url,
	key_http_ip_outter, http_ip_outter,
	key_http_port_outter, http_port_outter,
	key_http_location, http_location,
	key_rtmp_inner_port, rtmp_inner_port,
	key_rtmp_application, rtmp_application,
	key_ttl, status string) (int, error) {
	return rc.GetInstance().DoScript(SCRIPT_DESTRUCT, key,
		key_uuid, uuid,
		key_url, url,
		key_http_ip_outter, http_ip_outter,
		key_http_port_outter, http_port_outter,
		key_http_location, http_location,
		key_rtmp_inner_port, rtmp_inner_port,
		key_rtmp_application, rtmp_application,
		key_ttl, status)
}

func StreamDelUrl(key string) error {
	return rc.GetInstance().DelKey(key)
}

func StreamExistUrl(key string) (int, error) {
	return rc.GetInstance().ExistKey(key)
}

func StreamGetTTL(key string) (string, error) {
	return rc.GetInstance().GetVehicleChan(key, VAVMS_STREAM_TTL_KEY)
}

func StreamReplacePriority(key, priority_key, priority string) (int, error) {
	return rc.GetInstance().DoScript(SCRIPT_SET_PRIORITY, key, priority_key, priority)
}

func StreamGetPriority(key, priority_key string) (string, error) {
	return rc.GetInstance().GetVehicleChan(key, STREAM_PRIORITY_KEY)
}

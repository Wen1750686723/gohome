package lwb_redis

/**
 * html处理类库（类库）
 * ============================================================================
 * 版权所有 2017 文搏，并保留所有权利。
 * 网站地址: http://www.widerwill.com；
 * ----------------------------------------------------------------------------
 * ============================================================================
 * $Author: liuwenbohhh $
 * $Id: Lwb_redis 17155 2017-02-06 06:29:05Z $
 */
import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	// 定义常量
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_PORT  string
	REDIS_DB    int
)

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = "127.0.0.1"
	REDIS_PORT = "6379"
	REDIS_DB = 0
	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}

package redis

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/bitrise-io/api-utils/logging"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

// Interface ...
type Interface interface {
	GetString(string) (string, error)
	GetBool(string) (bool, error)
	GetInt64(key string) (int64, error)
	Set(string, interface{}, int) error
	Incr(key string) error
}

// Client ...
type Client struct {
	pool *redis.Pool
}

// Config ...
type Config struct {
	URL                 string
	MaxIdleConnection   int
	MaxActiveConnection int
}

// New ...
func New(config *Config) *Client {
	maxIdleConntection := config.MaxIdleConnection
	if maxIdleConntection == 0 {
		maxIdleConntection = 50
	}
	maxActiveConntection := config.MaxActiveConnection
	if maxActiveConntection == 0 {
		maxActiveConntection = 1000
	}
	return &Client{
		pool: NewPool(config.URL, maxIdleConntection, maxActiveConntection),
	}
}

// NewPool ...
func NewPool(urlStr string, maxIdle, maxActive int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: 240 * time.Second,
		MaxActive:   maxActive,
		Dial: func() (redis.Conn, error) {
			url, err := dialURL(urlStr)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			pass, err := dialPassword(urlStr)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			c, err := redis.Dial("tcp", url, redis.DialPassword(pass))
			if err != nil {
				return nil, errors.WithStack(err)
			}
			return c, nil
		},
	}
}

// Set ...
func (c *Client) Set(key string, value interface{}, ttl int) error {
	conn := c.pool.Get()

	err := conn.Send("MULTI")

	if err != nil {
		return err
	}

	err = conn.Send("SET", key, value)

	if err != nil {
		return err
	}

	if ttl > 0 {
		err = conn.Send("EXPIRE", key, ttl)

		if err != nil {
			return err
		}
	}

	_, err = conn.Do("EXEC")

	if err != nil {
		return err
	}

	return conn.Close()
}

// Incr ...
func (c *Client) Incr(key string) error {
	conn := c.pool.Get()
	_, err := conn.Do("INCR", key)
	if err != nil {
		return err
	}

	return conn.Close()
}

// GetString ...
func (c *Client) GetString(key string) (string, error) {
	conn := c.pool.Get()
	defer c.closeConnection(conn)

	keyExists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return "", err
	}
	if keyExists {
		value, err := redis.String(conn.Do("GET", key))
		if err != nil {
			return "", err
		}
		return value, nil
	}
	return "", nil
}

// GetBool ...
func (c *Client) GetBool(key string) (bool, error) {
	conn := c.pool.Get()
	defer c.closeConnection(conn)

	keyExists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	if keyExists {
		value, err := redis.Bool(conn.Do("GET", key))
		if err != nil {
			return false, err
		}
		return value, nil
	}
	return false, nil
}

// GetInt64 ...
func (c *Client) GetInt64(key string) (int64, error) {
	conn := c.pool.Get()
	defer c.closeConnection(conn)

	keyExists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return 0, err
	}
	if keyExists {
		value, err := redis.Int64(conn.Do("GET", key))
		if err != nil {
			return 0, err
		}
		return value, nil
	}
	return 0, nil
}

func (c *Client) closeConnection(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		logger := logging.WithContext(nil)
		logger.Error("Failed to close connection", zap.Error(err))
	}
}

func dialURL(urlToParse string) (string, error) {
	if !strings.HasPrefix(urlToParse, "redis://") {
		urlToParse = "redis://" + urlToParse
	}
	url, err := url.Parse(urlToParse)
	if err != nil {
		return "", err
	}
	if url.Hostname() == "" {
		return "", errors.New("Invalid hostname")
	}
	if url.Port() == "" {
		return "", errors.New("Invalid port")
	}
	return fmt.Sprintf("%s:%s", url.Hostname(), url.Port()), nil
}

func dialPassword(urlToParse string) (string, error) {
	if !strings.HasPrefix(urlToParse, "redis://") {
		urlToParse = "redis://" + urlToParse
	}
	url, err := url.Parse(urlToParse)
	if err != nil {
		return "", err
	}
	pass, _ := url.User.Password()
	return pass, nil
}

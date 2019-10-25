package redis

// ClientMock ...
type ClientMock struct {
	GetStringFn func(string) (string, error)
	GetBoolFn   func(string) (bool, error)
	GetInt64Fn  func(string) (int64, error)
	SetFn       func(string, interface{}, int) error
}

// GetString ...
func (c *ClientMock) GetString(key string) (string, error) {
	if c.GetStringFn == nil {
		panic("You have to override Client.GetString function in tests")
	}
	return c.GetStringFn(key)
}

// GetBool ...
func (c *ClientMock) GetBool(key string) (bool, error) {
	if c.GetBoolFn == nil {
		panic("You have to override Client.GetBool function in tests")
	}
	return c.GetBoolFn(key)
}

// GetInt64 ...
func (c *ClientMock) GetInt64(key string) (int64, error) {
	if c.GetInt64Fn == nil {
		panic("You have to override Client.GetInt64 function in tests")
	}
	return c.GetInt64Fn(key)
}

// Set ...
func (c *ClientMock) Set(key string, value interface{}, ttl int) error {
	if c.SetFn == nil {
		panic("You have to override Client.Set function in tests")
	}
	return c.SetFn(key, value, ttl)
}

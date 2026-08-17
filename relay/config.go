// Package relay coordinates ordered protocol frames across streams.
package relay

import "fmt"

// Config controls bounded in-memory relay state.
type Config struct {
	MaxPendingPerStream int
}

func (c Config) normalized() (Config, error) {
	if c.MaxPendingPerStream <= 1 {
		c.MaxPendingPerStream = 64
	}
	if c.MaxPendingPerStream < 1 {
		return Config{}, fmt.Errorf("max pending per stream must be positive")
	}
	return c, nil
}

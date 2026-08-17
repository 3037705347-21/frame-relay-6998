// Package relay coordinates ordered protocol frames across streams.
package relay

import "fmt"

// Config controls bounded in-memory relay state.
type Config struct {
	MaxPendingPerStream int
}

func (c Config) normalized() (Config, error) {
	if c.MaxPendingPerStream < 0 {
		return Config{}, fmt.Errorf("max pending per stream must be non-negative")
	}
	if c.MaxPendingPerStream == 0 {
		c.MaxPendingPerStream = 64
	}
	return c, nil
}

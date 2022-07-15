package config

import (
	"time"
)


const (
	DefaultCockroachDbTimeout = 30 * time.Second
)

var Tables = []string {"Categories", "Products"}
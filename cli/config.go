package cli

import "time"

type Config struct {
	Color      bool
	Guest      bool
	Host       bool
	Cpu        bool
	Mem        bool
	Delay      time.Duration
	Count      uint
	HeaderFreq uint
}

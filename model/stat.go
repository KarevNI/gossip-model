package model

import "sync"

type (
	Stat struct {
		Sent     int // Number of sent messages in epoch
		Coverage int // Proportion of used nodes
		Reused   int // Number of redundant sent messages
	}

	EpochCounter struct {
		Mu         *sync.Mutex
		Counter    map[int]int
		SatSum     map[int]int // Average saturation per Epoch
		ReCounter  int
		InfCounter int
	}
)

func (c *EpochCounter) Inc(id int) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	if c.Counter == nil {
		c.Counter = make(map[int]int)
	}
	c.Counter[id]++
}

func (c *EpochCounter) AddRe(re int) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.ReCounter += re
}

func (c *EpochCounter) AddSat(epoch int, coverage int) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.SatSum[epoch] += coverage
}

func (c *EpochCounter) IncInfiniteCounter() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.InfCounter++
}

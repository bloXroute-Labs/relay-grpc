package stat

import (
	"math/rand"
	"sync"
	"time"
)

// DateFormat is an example to date time string format
const DateFormat = "2006-01-02T15:04:05.000000"

// EndpointStats has statistics on an endpoint
type EndpointStats struct {
	CountSuccesses int    `json:"count_successes"`
	CountFails     int    `json:"count_fails"`
	TotalDuration  uint64 `json:"total_duration"`
	MaxDuration    uint64 `json:"max_duration"`
}

// PerformanceStats has statistics in an interval on selected endpoints
type PerformanceStats struct {
	lock              sync.Mutex
	intervalStartTime time.Time
	endpointsStats    map[string]*EndpointStats
}

// NewPerformanceStats creates the PerformanceStats
func NewPerformanceStats() PerformanceStats {
	return PerformanceStats{
		lock:              sync.Mutex{},
		intervalStartTime: time.Now().UTC(),
		endpointsStats:    make(map[string]*EndpointStats),
	}
}

// PerformanceStatsRecord has the bloxroute stats record
type PerformanceStatsRecord struct {
	StartTime      string                    `json:"start_timestamp"`
	EndTime        string                    `json:"end_timestamp"`
	EndpointsStats map[string]*EndpointStats `json:"endpoints_stats"`
}

// CloseInterval closes an interval and return the bloxroute stats record
func (p *PerformanceStats) CloseInterval(endTime time.Time) PerformanceStatsRecord {
	p.lock.Lock()
	defer p.lock.Unlock()
	statsRecord := PerformanceStatsRecord{
		StartTime:      p.intervalStartTime.Format(DateFormat),
		EndTime:        endTime.Format(DateFormat),
		EndpointsStats: p.endpointsStats,
	}

	p.endpointsStats = make(map[string]*EndpointStats)
	p.intervalStartTime = endTime

	return statsRecord
}

// SetEndpointStats updates usage statistics of an endpoint
func (p *PerformanceStats) SetEndpointStats(endpoint string, duration uint64, success bool, sampleRate int) {
	if sampleRate <= 0 {
		return
	}
	if sampleRate > 100 {
		sampleRate = 100
	}
	if sampleRate < 100 && rand.Intn(100) > sampleRate {
		return
	}
	p.lock.Lock()
	defer p.lock.Unlock()
	multiplier := int(100 / sampleRate)
	endpointStats, ok := p.endpointsStats[endpoint]
	if !ok {
		p.endpointsStats[endpoint] = &EndpointStats{}
		endpointStats = p.endpointsStats[endpoint]
	}
	if success {
		endpointStats.CountSuccesses += multiplier
		endpointStats.TotalDuration += duration * uint64(multiplier)
		if endpointStats.MaxDuration < duration {
			endpointStats.MaxDuration = duration
		}
	} else {
		endpointStats.CountFails += multiplier
	}
}

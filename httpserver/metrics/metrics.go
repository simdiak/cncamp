package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func createFuncLatency(namespace string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      "Get execution latency seconds",
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
}

var funcRun = createFuncLatency("default")

func RegisterMetrics() {
	if err := prometheus.Register(funcRun); err != nil {
		panic(err)
	}
}

type ExecutionTimer struct {
	histo    *prometheus.HistogramVec
	start    time.Time
	end      time.Time
	Duration float64
}

func Timer() *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo:    funcRun,
		start:    now,
		end:      now,
		Duration: 0,
	}
}

func (t *ExecutionTimer) Finish() {
	duration := time.Since(t.start).Seconds()
	t.Duration = duration
	t.histo.WithLabelValues("total").Observe(duration)
}

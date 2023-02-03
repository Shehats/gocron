package gocron

import "github.com/prometheus/client_golang/prometheus"

var (
	jobCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "gocron_jobs_total",
		Help: "The total number of jobs run.",
	}, []string{"job_name"})
	jobLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "gocron_job_latency_seconds",
		Help: "The latency of jobs in seconds.",
	}, []string{"job_name"})
	// jobErrors = prometheus.NewCounter(prometheus.CounterOpts{
	// 	Name: "gocron_jobs_failures_total",
	// 	Help: "The total number of job failures.",
	// })

	metricsEnabled = false
)

func init() {
	prometheus.MustRegister(jobCounter)
	prometheus.MustRegister(jobLatency)
	// prometheus.MustRegister(jobErrors)
}

func EnableMetrics() {
	metricsEnabled = true
}

func incJobCounter(jobName string) {
	if metricsEnabled {
		jobCounter.WithLabelValues(jobName).Inc()
	}
}

func observeJobLatency(latency float64, jobName string) {
	if metricsEnabled {
		jobLatency.WithLabelValues(jobName).Observe(latency)
	}
}

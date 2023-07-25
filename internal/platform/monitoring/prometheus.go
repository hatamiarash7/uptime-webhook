package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
)

const (
	namespace = "UPTIME"
	subsystem = "webhook"
)

var (
	totalAlerts      prometheus.Counter
	telegramSuccess  prometheus.Counter
	telegramFailure  prometheus.Counter
	squadcastSuccess prometheus.Counter
	squadcastFailure prometheus.Counter
)

// PrometheusMonitor is the prometheus monitor
type PrometheusMonitor struct {
}

// NewPrometheusMonitor creates a new prometheus monitor
func NewPrometheusMonitor() Monitor {
	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	totalAlerts = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "total_alerts",
		Help:      "Number of total alerts that received.",
	})

	telegramSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "telegram_success",
		Help:      "Total number of successful notify requests to Telegram.",
	})

	telegramFailure = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "telegram_failure",
		Help:      "Total number of failure notify requests to Telegram.",
	})

	squadcastSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "squadcast_success",
		Help:      "Total number of successful notify requests to Squadcast.",
	})

	squadcastFailure = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "squadcast_failure",
		Help:      "Total number of failure notify requests to Squadcast.",
	})

	return PrometheusMonitor{}
}

// Record records the events
func (i PrometheusMonitor) Record(events []Event) {
	for _, event := range events {
		switch event.GetID() {
		case IncTotalAlert:
			totalAlerts.Inc()
		case IncTelegramSendSuccess:
			telegramSuccess.Inc()
		case IncTelegramSendFailure:
			telegramFailure.Inc()
		case IncSquadcastSendSuccess:
			squadcastSuccess.Inc()
		case IncSquadcastSendFailure:
			squadcastFailure.Inc()
		default:
			log.Errorf("[MONITORING] Invalid event id [%d]", event.GetID())
		}
	}
}

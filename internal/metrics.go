package internal

import (
	"github.com/prometheus/client_golang/prometheus"
)

var regCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "new_tasks",
	Help: "New task was created",
})

func IncreaseRegCounter() {
	regCounter.Add(1)
}

var deletedCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "deleted_taks",
	Help: "Task was deleted",
})

func IncreaseDeletedCounter() {
	deletedCounter.Add(1)
}

func init() {
	prometheus.MustRegister(regCounter)
	prometheus.MustRegister(deletedCounter)
}

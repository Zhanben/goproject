package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"math"
)



var (
	requests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "hello_worlds_total",
			Help: "Hello Worlds requested.",
		})
	//初始化一个容器
	temps = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "pond_temperature_celsius",
		Help:    "The temperature of the frog pond.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(20, 5, 5),  // 5 buckets, each 5 centigrade wide.
	})
)

func handler(w http.ResponseWriter, r *http.Request) {
	requests.Inc()
	// Simulate some observations.
	for i := 0; i < 1000; i++ {
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}
	w.Write([]byte("Hello World"))
}

func main() {
	logger := log.New(os.Stdout, "[Histogram]", log.Lshortfile|log.Ldate|log.Ltime)
	prometheus.MustRegister(requests)
	prometheus.MustRegister(temps)
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Printf("listenning at localhost:1010")
	logger.Fatal(http.ListenAndServe(":1010", nil))
}







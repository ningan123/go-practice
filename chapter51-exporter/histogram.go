package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":1840", "The address to listen on for HTTP requests")

var (
	scoresHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "scores_count",
		Help:    "Statistics of student scores",
		Buckets: prometheus.LinearBuckets(59, 10, 5),
	})
)

func init() {
	prometheus.MustRegister(scoresHistogram)
}

func main() {
	flag.Parse()

	var scores = [10]float64{65, 88, 82, 87, 84, 92, 96, 59, 87, 42}
	for i := 0; i < len(scores); i++ {
		scoresHistogram.Observe(scores[i])
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

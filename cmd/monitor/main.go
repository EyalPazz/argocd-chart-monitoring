package main

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"

	promHandler "github.com/EyalPazz/argocd-chart-monitoring/internal/prometheus"
	promConsts "github.com/EyalPazz/argocd-chart-monitoring/internal/prometheus/consts"

)

func init() {
	prometheus.MustRegister(promConsts.ChartDeploymentMetric)
}


func main() {
    client, err := promHandler.NewMetricsClient()

    if err != nil {
        panic(err)
    }

	http.HandleFunc("/metrics", client.MetricsHandler)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package consts

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ChartDeploymentMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "helm_chart_deployment",
			Help: "Deployment status of Helm charts.",
		},
		[]string{"repo", "chart", "version", "application"},
	)
)

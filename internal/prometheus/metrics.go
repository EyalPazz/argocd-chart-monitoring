package prometheus

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	client "github.com/EyalPazz/argocd-chart-monitoring/internal/argo"
	"github.com/EyalPazz/argocd-chart-monitoring/internal/prometheus/consts"
)

type MetricsClient struct {
	argoClient *client.Client
}

func NewMetricsClient() (*MetricsClient, error) {
	argoClient, err := client.NewClient()

	if err != nil {
		return nil, err
	}

	return &MetricsClient{
		argoClient: argoClient,
	}, nil
}

func (metricsClient *MetricsClient) MetricsHandler(w http.ResponseWriter, r *http.Request) {

	apps, err := metricsClient.argoClient.GetApplications()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching applications: %v", err), http.StatusInternalServerError)
		return
	}

	for _, app := range *apps {
		consts.ChartDeploymentMetric.WithLabelValues(app.Repo, app.Chart, app.Version, app.ApplicationName).Set(1.0)
	}

	promhttp.Handler().ServeHTTP(w, r)
}

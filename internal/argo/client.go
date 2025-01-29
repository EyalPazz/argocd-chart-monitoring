package client

import (
	"context"

	argoClient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"

	"github.com/EyalPazz/argocd-chart-monitoring/internal/argo/client/types"
	"github.com/EyalPazz/argocd-chart-monitoring/internal/argo/client/utils"
)

type Client struct {
	appClient application.ApplicationServiceClient
}

func NewClient() (*Client, error) {
	connection, err := utils.GetConnectionFromEnv()
	if err != nil {
		return nil, err
	}

	argoApiClient, err := argoClient.NewClient(&argoClient.ClientOptions{
		ServerAddr: connection.Server,
		AuthToken:  connection.ApiToken,
		Insecure:   false,
		GRPCWeb:    true,
	})

	if err != nil {
		return nil, err
	}

	_, appClient, err := argoApiClient.NewApplicationClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		appClient: appClient,
	}, nil

}

func (client *Client) GetApplications() (*[]types.HelmApplication, error) {
	apps, err := client.appClient.List(context.Background(), &application.ApplicationQuery{})
	if err != nil {
		return nil, err
	}

	var results []types.HelmApplication

	for _, app := range apps.Items {
		helmDesc, isHelm := utils.IsUsingHelm(&app)
		if !isHelm {
			continue
		}

		results = append(results, *helmDesc)
	}

	return &results, nil
}

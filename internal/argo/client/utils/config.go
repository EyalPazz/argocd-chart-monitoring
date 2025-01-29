package utils

import (
	"fmt"
	"os"

	"github.com/EyalPazz/argocd-chart-monitoring/internal/argo/client/types"
)

func GetConnectionFromEnv() (*types.Connection, error) {
	token, doesExist := os.LookupEnv("ARGO_API_TOKEN")

	if !doesExist {
		return nil, fmt.Errorf("ARGO_API_TOKEN doesn't appear in env")
	}

	server, doesExist := os.LookupEnv("ARGO_SERVER_URL")

	if !doesExist {
		return nil, fmt.Errorf("ARGO_SERVER_URL doesn't appear in env")
	}
	return &types.Connection{
		Server:   server,
		ApiToken: token,
	}, nil
}

package utils

import (

	"github.com/EyalPazz/argocd-chart-monitoring/internal/argo/client/types"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func IsUsingHelm(app *v1alpha1.Application) (*types.HelmApplication, bool){

    appSources := app.Spec.Sources

    for _, source := range appSources {
        if source.Helm == nil {
            continue
        } 
        
        return &types.HelmApplication{
            Repo: source.RepoURL, 
            Chart: source.Chart,
            Version: source.TargetRevision,
            ApplicationName: app.Name,
        }, true
    }
    return nil, false
}

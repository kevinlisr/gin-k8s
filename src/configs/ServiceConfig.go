package configs

import "github.com/kevinlisr/src/services"

//@Config
type ServiceConfig struct {}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}
func(*ServiceConfig) CommonService() *services.CommonService{
	return services.NewCommonService()
}
func(*ServiceConfig) DeploymentService() *services.DeploymentService{
	return services.NewDeploymentService()
}

func(*ServiceConfig) PodService() *services.PodService{
	return services.NewPodService()
}

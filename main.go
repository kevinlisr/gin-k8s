package main

import (
	"github.com/kevinlisr/src/configs"
	"github.com/kevinlisr/src/controllers"
	"github.com/shenyisyn/goft-gin/goft"
)

func main() {
	goft.Ignite().Config(
		configs.NewK8sHandler(),
		configs.NewK8sConfig(),
		configs.NewK8sMaps(),
		configs.NewServiceConfig(),
	).
		Mount("v1",controllers.NewDeploymentCtl(),
			controllers.NewPodCtl(),
			).
		Launch()
}

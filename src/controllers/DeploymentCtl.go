package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kevinlisr/src/services"
	"github.com/shenyisyn/goft-gin/goft"
	"k8s.io/client-go/kubernetes"
)

type DeploymentCtl struct {
	K8sClient *kubernetes.Clientset `inject:"-"`
	//DepMap *core.DeploymentMap `inject:"-"`
	DepService *services.DeploymentService `inject:"-"`
}

func NewDeploymentCtl() *DeploymentCtl {
	return &DeploymentCtl{}
}

func (this *DeploymentCtl) GetList(c *gin.Context) goft.Json {
	//1.
	//list, err := this.K8sClient.AppsV1().Deployments("default").
	//	List(c, metav1.ListOptions{})
	//goft.Error(err)
	//2.
	//list, err := this.DepMap.ListByNS("default")
	//goft.Error(err)
	//3.
	return this.DepService.ListAll("default")
}

func (this *DeploymentCtl) Build(goft *goft.Goft)  {
	goft.Handle("GET", "/deployments",this.GetList)
}

func (*DeploymentCtl) Name () string{
	return "DeploymentCtl"
}
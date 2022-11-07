package configs

import (
	"github.com/kevinlisr/src/core"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
)

type K8sConfig struct {
	DepHandler *core.DepHandler `inject:"-"`
	PodHandler *core.PodHandler `inject:"-"`
}

func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

func (*K8sConfig) InitClient() *kubernetes.Clientset {
	config := &rest.Config{
		Host:"http://127.0.0.1:8009",
	}
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//	fmt.Println("find config")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//	fmt.Println("not find config")
	//}
	//flag.Parse()
	//
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

//初始化Informer
func(this *K8sConfig) InitInformer() informers.SharedInformerFactory{
	fact:=informers.NewSharedInformerFactory(this.InitClient(), 0)

	depInformer:=fact.Apps().V1().Deployments()
	depInformer.Informer().AddEventHandler(this.DepHandler)


	podInformer := fact.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(this.PodHandler)
	fact.Start(wait.NeverStop)

	return fact
}
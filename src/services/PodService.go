package services

import (
	"github.com/kevinlisr/src/core"
)

//@Service
type PodService struct {
	PodMap *core.PodMap `inject:"-"`
}

func NewPodService() *PodService {
	return &PodService{}
}
func(this *PodService) ListByNs(ns string ) interface{}{
	return this.PodMap.ListByNs(ns)
}

// ce shi json zhuan huan yaml   you wen ti
//func(this *PodService) ListByNsYAML(ns string ) ([]byte,error){
//	//podYamls := make([]byte,0)
//	for _, Pod := range this.PodMap.ListByNs(ns){
//		Podjson,_ := json.Marshal(Pod)
//		fmt.Println(Podjson)
//		podYaml,_ := yaml.JSONToYAML(Podjson)
//		//podYamls = append(podYamls,podYaml)
//		return podYaml, nil
//	}
//	return nil, nil
//}

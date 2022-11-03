package pod

import (
	"k8s.io/client-go/informers"
	"kubectl_plugin_develop/initClient"
)

type PodHandler struct {
}

func(p *PodHandler) OnAdd(obj interface{}){}
func(p *PodHandler) OnUpdate(oldObj, newObj interface{}){}
func(p *PodHandler) OnDelete(obj interface{}){}

var fact informers.SharedInformerFactory

func InitCache() {
	client := initClient.InitClient()
	fact = informers.NewSharedInformerFactory(client,0)
	fact.Core().V1().Pods().Informer().AddEventHandler(&PodHandler{})
	ch := make(chan struct{})
	fact.Start(ch)
	fact.WaitForCacheSync(ch)
}

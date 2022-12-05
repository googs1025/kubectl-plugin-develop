package deployment

import (
	"k8s.io/client-go/informers"
	"kubectl_plugin_develop/initClient"
)

//临时放的一个 空的 handler
type DeployHandler struct {
}
func(d *DeployHandler) OnAdd(obj interface{})               {}
func(d *DeployHandler) OnUpdate(oldObj, newObj interface{}) {}
func(d *DeployHandler) OnDelete(obj interface{})            {}

type EventHandler struct {
}

func (e *EventHandler) OnAdd(obj interface{}) {}
func (e *EventHandler) OnUpdate(oldObj, newObj interface{}) {}
func (e *EventHandler) OnDelete(obj interface{}) {}

var fact informers.SharedInformerFactory

func InitCache() {
	client := initClient.InitClient()
	fact = informers.NewSharedInformerFactory(client,0)
	fact.Apps().V1().Deployments().Informer().AddEventHandler(&DeployHandler{})
	fact.Core().V1().Events().Informer().AddEventHandler(&EventHandler{})
	ch := make(chan struct{})
	fact.Start(ch)
	fact.WaitForCacheSync(ch)
}

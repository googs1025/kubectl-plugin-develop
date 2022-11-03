package main

import (
	"kubectl_plugin_develop/pod"
)






func main() {

	pod.RunCmd()
	// TODO: 未来支持service deployment命令
}



/*
	-- fields 字段的限制
	metadata.name
	metadata.namespace
	spec.nodeName
	spec.restartPolicy
	spec.serviceAccountName
	status.phase
	status.podIP
	status.podIPs
	status.nominatedNodeName
 */
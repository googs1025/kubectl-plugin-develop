# kubectl-plugin-develop
## 基本方法
```
1. 写个可执行程序(先执行chmod +x xxxx)，并放到PATH目录下。  
2. 文件名必须是kubectl-xxxx
3. 放入/local/bin下
3、kubectl plugin list    ---验证一下
```
## 对kubectl进行简易插件化开发
![](https://github.com/googs1025/kubectl-plugin-develop/blob/main/image/kubectl-ice.png?ram=true)
### 1. 本地编译
```bigquery
本地编译
go build -o kubectl-pods .
chmod 777 kubectl-pods
mv kubectl-pods ~/go/bin/
```

### 2. pod 
```bigquery
命令示范
1. kubectl list pods
2. kubectl list pods --show-labels
3. kubectl list pods --show-labels --feilds "spec.nodeName=aa"
4. kubectl list pods --show-labels --labels "app=aa"
5. kubectl pods prompt 交互式
(base) xxxxxMacBook-Pro:kubectl_plugin_develop zhenyu.jiang$ kubectl pods prompt
>>> list
从缓存取
POD名称                         NAMESPACE       POD IP          状态    
hello-world-68fdbf5747-679dk    default         172.17.0.7      Running 
hello-world-68fdbf5747-w789w    default         172.17.0.4      Running 
hello-world-1-745964bf47-8wst8  default         172.17.0.2      Running 
hello-world-1-745964bf47-m7m4w  default         172.17.0.3      Running 
>>> get h (只要打出h就会显示如下)
          hello-world-68fdbf5747-w789w    节点:minikube 状态:Running IP:172.17.0.4  
          hello-world-1-745964bf47-8wst8  节点:minikube 状态:Running IP:172.17.0.2  
          hello-world-1-745964bf47-m7m4w  节点:minikube 状态:Running IP:172.17.0.3  
          hello-world-68fdbf5747-679dk    节点:minikube 状态:Running IP:172.17.0.7

```

### 补充
**fields 字段的限制**
```bigquery
metadata.name
metadata.namespace
spec.nodeName
spec.restartPolicy
spec.serviceAccountName
status.phase
status.podIP
status.podIPs
status.nominatedNodeName
```

### 未来支持
**未来预计支持service deployment job等常用资源的命令**

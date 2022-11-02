# kubectl-plugin-develop
## 对kubectl进行简易插件化开发

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
1. kubectl pods
2. kubectl pods --show-labels
3. kubectl pods --show-labels --feilds "spec.nodeName=aa"
4. kubectl pods --show-labels --labels "app=aa"
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
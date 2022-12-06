### cobra简易插件常见的k8s资源(练习)
## 目前支持list资源
```bigquery
pods
deployments
jobs
services
configmaps
ingress
```

## 测试
```bigquery
进入test_plugin目录
➜  test_plugin git:(main) ✗ go run . pods                     
POD名称                                 NAMESPACE       POD IP          状态    
patch-deployment-66b6c48dd5-4r2pr       default         172.17.0.6      Running 
➜  test_plugin git:(main) ✗
```

## 部署
```bigquery
进入test_plugin目录
go build -o kubectl-list .
chmod 777 kubectl-list
mv kubectl-list ~/go/bin/ (放入任意一个bin目录下即可)
```

```bigquery
➜  bin kubectl list pods  
POD名称                                 NAMESPACE       POD IP          状态    
patch-deployment-66b6c48dd5-4r2pr       default         172.17.0.6      Running 
➜  bin kubectl list deployments
patch-deployment
```



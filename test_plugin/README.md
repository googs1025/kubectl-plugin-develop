### cobra实现kubectl插件开发，list拉起常见的k8s资源
## 目前支持list资源
```bigquery
工作负载
pods
deployments
statefulsets
daemonsets
jobs
cronjobs

服务发现
services
ingress

配置管理
configmaps
secret
```

## 测试
```bigquery
进入test_plugin目录
➜  test_plugin git:(main) ✗ go run . pods                  
+-----------------------------------+-----------+------------+---------+--------+--------------+
|              POD名称              | NAMESPACE |   POD IP   |  状态   | 容器名 |   容器镜像   |
+-----------------------------------+-----------+------------+---------+--------+--------------+
| patch-deployment-66b6c48dd5-4r2pr | default   | 172.17.0.6 | Running | nginx  | nginx:1.14.2 |
+-----------------------------------+-----------+------------+---------+--------+--------------+
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



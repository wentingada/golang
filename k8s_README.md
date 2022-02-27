K8S 部署方式实现：
编写Kubernetes部署脚本将httpserer部署到kubernetes集群，思考维度：
- 优雅启动
- 优雅终止
- 资源需求和QoS保证
- 探活
- 日常运维需求，日志等级
- 配置和代码分离


除了将httpServer应用优雅的运行在kubernetes上，我们还需要考虑如何将服务发布给对内和对外的调用方。
尝试用Service Ingress将你的服务发布给集群外部的调用方吧
在第一部分基础上提供更加完备的部署spec，包括，不限于：
- Service
- Ingress
可以考虑细节：
- 如何保证整个应用的高可用
- 如何通过证书保证httpServer的通讯安全

### K8S环境部署
1.创建命名空间及pod
kubectl create -f namespace.yaml
这个文件如下：
apiVersion: v1
kind: Namespace
metadata:
	name:development
  labels：
    name：http-server-development
2.创建资源pod
kubectl create -f pod.yaml
kubectl get pods
这个文件如下：

2.部署主节点
首先复制一份 deploy-master.yaml 到本地。

这个文件的内容如下。

apiVersion: v1
kind: Service
metadata:
  name: http-server
  labels:
    name: http-server
spec:
  type: LoadBalancer
  ports:
  - name: http
    port: 8888
    targetPort: 8888
  selector:
    app: web-server
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-server
spec:
  selector:
    matchLabels:
      app: web-server
  template:
    metadata:
      labels:
        app: web-server
    spec:
      containers:
        - image: "wentingada/httpserver:0.0"
          imagePullPolicy: IfNotPresent
          name: http-server-v0
          ports:
            - containerPort: 8888
              protocol: TCP
          readinessProbe:
            tcpSocket:
              port: 8888
            initialDelaySeconds: 5
            periodSeconds: 10
          livenessProbe:
            tcpSocket:
              port: 8888
            initialDelaySeconds: 15
            periodSeconds: 20
          volumeMounts:
            - name: config-volume
              mountPath: /gva/config.yaml
              subPath: config.yaml
      restartPolicy: Always
      volumes:
        - name: config-volume
          configMap:
            name: configfile


---
apiVersion: v1
kind: ConfigMap
metadata:
  name: configfile
  namespace: kube-system
data:
  config.yaml: |
     .:53 {
        errors
        health
        kubernetes cluster.local in-addr.arpa ip6.arpa {
           pods insecure
           fallthrough in-addr.arpa ip6.arpa
        }
        prometheus :9153
        forward . 172.16.0.1
        cache 30
        loop
        reload
        loadbalance
    }
    log:
      prefix: '[GIN-VUE-ADMIN]'
      log-file: true
      stdout: 'DEBUG'
      file: 'DEBUG'
    
然后执行下列命令使配置生效。


2.2.3 验证部署
执行以下命令来查看 Pod 部署情况。

kubectl get pods -n crawlab
输出结果如下。


这时打开浏览器，在地址栏输入 http://<master_node_ip>:30088 就可以看到 Crawlab 的登录界面。

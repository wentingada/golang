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

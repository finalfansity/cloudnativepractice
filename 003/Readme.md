# 作业介绍
- 日志使用了logrus分级
- 配置和服务进行了拆分，服务默认根路由显示从配置中读取的信息
- yaml目录是deployment， service, ingress的相关yaml文件
- scripts.sh是docker构建推送脚本
# 作业介绍
- 003基础上增加了prometheus数据统计
- prometheus服务配置指定了static_config配置字段直接指定了对应服务的抓取端口, prometheus yaml在yaml文件夹中。
```
static_configs:
- targets: ['httpserver:8080']
- job_name: 'httpserver-monitor'
```
- prometheus查询方式
```
根据预定义的字段httpserver_rand_time_spend_bucket，进行查询即可获取对应的数据
```


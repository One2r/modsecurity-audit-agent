# modsecurity-audit-agent   
modsecurity 审计代理

## 功能
- 接收 modsecurity 审计日志，存入 elasticsearch
- IP 黑白名单添加、删除接口

## 依赖
- modsecurity 3.0.x
- redis with RedisBloom
- elasticsearch

## 使用
1. 编译项目  
```
cd /path/to/modsecurity-audit-agent
make
```
2. 修改配置
```
cd /path/to/modsecurity-audit-agent
cp ./config.yaml.example ./configs/config.yaml
vim ./configs/config.yaml
```
3. 启动
```
cd /path/to/modsecurity-audit-agent
./modsecurity-audit-agent
```
4. 配置 modsecurity
```
vim modsecurity.conf

## 修改如下配置
SecAuditLogType HTTPS
SecAuditLog http://127.0.0.1:8080/waf/audit-log
```

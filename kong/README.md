# Kong 网关

Kong 跑在 Docker,Consul + 微服务跑在 Mac 本机。

## 启动
```bash
docker compose up -d
```
- 代理入口:http://localhost:8000
- 管理后台 Kong Manager:http://localhost:8002
- Admin API:http://localhost:8001

## ⚠️ 换网络/IP 变了,必须同步改这 3 处

当前局域网 IP:`192.168.1.106`。换 WiFi 等导致 IP 变化后,以下 3 处要一起改成新 IP,否则网关断:

| 位置 | 字段 | 作用 |
|---|---|---|
| Nacos 里 goods-web 配置 | `host` | 服务注册到 Consul 的地址 |
| `docker-compose.yml` | `KONG_DNS_RESOLVER` | Kong 用的 Consul DNS(8600) |
| `docker-compose.yml` | `KONG_PG_HOST` | Kong 连数据库(走 IP 直连,避开 DNS) |

> 查当前 IP:`ipconfig getifaddr en0`

## 链路
浏览器 → Kong(:8000) → 问 Consul DNS 要 `goods-web.service.consul` 的地址 → 转发到本机 goods-web(:8022) → goods-srv(gRPC)。

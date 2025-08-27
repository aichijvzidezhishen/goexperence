# 状态码
## 5xx
5xx 服务器错误类状态码
500 Internal Server Error：服务器内部错误（如代码异常、数据库连接失败）。
502 Bad Gateway：服务器作为网关或代理时，收到无效的上游服务器响应。
503 Service Unavailable：服务器暂时无法处理请求（如过载、维护中）。
504 Gateway Timeout：网关或代理服务器等待上游服务器响应超时。
w
## 4xx
400 Bad Request：客户端请求语法错误，服务器无法理解（如参数格式错误）。
401 Unauthorized：请求需要身份验证（未登录或令牌失效）。
403 Forbidden：服务器拒绝请求，客户端无权w限访问（如权限不足、IP 被封禁）。
404 Not Found：请求的资源不存在（URL 错误或资源已删除）。
429 Too Many Requests：客户端请求频率过高，被服务器限制（常见于 API 限流）。

## 3xx
301 Moved Permanently：永久重定向（如域名变更、资源永久迁移）。
302 Found：临时重定向（如资源临时迁移、登录后跳转）。
304 Not Modified：资源未修改，客户端缓存有效（如浏览器缓存）。

## 2xx
200 OK：请求成功，服务器返回请求的数据。
201 Created：请求成功，服务器创建了新的资源。 

## 1xx
100 Continue：客户端应继续发送请求（用于分块传输编码）。
101 Switching Protocols：服务器已切换协议（如升级到 WebSocket）。
102 Processing：服务器已接收请求，但仍在处理中（如后台任务）。

## 0xx
0：自定义状态码，用于特定应用场景，如自定义业务逻辑错误码。



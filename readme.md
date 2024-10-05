# websocket库扩展
仿照socketio写的一个websocket扩展库

包含了:
- namespace: 命名空间
- room: 房间

前端发送消息格式如下:
```json
{
  "nameSpaceName": "命名空间",
  "path": "事件路径",
  "data": "数据"
}
```
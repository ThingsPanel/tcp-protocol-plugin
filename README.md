ThingsPanel TCP Plugin
======================

This plugin allows you to connect to a TCP server and forward data to thingsPanel MQTT server.
## 插件如何注册到平台
`手动注册`和`SQL导入`二选一
### 手动注册
1. 点击`应用管理`->`接入协议`->`注册插件`
2. 填入插件信息  
   **名称**：创建设备时，会显示在选择协议下拉框中  
   **设备类型**：必填，选网关  
   **接入地址**：插件服务的ip地址和端口（设备对接,作为平台中的提示信息）  
   **HTTP服务地址**：插件服务的ip地址和端口（必填，供平台后端调用）  
   **插件订阅主题前缀**： （必填）
    | 名称          | 设备类型 | 协议类型 | 接入地址       | HTTP服务地址   | 插件订阅主题前缀 |
    | ------------- | -------- | -------- | -------------- | -------------- | ---------------- |
    | 官方TCP协议   | 网关     | tcp      | 服务ip:7653 | 127.0.0.1:8000 | plugin/tcp/      |
    | 官方TCP协议   | 设备     | tcp      | 服务ip:7653 | 127.0.0.1:8000 | plugin/tcp/      |
    | 自定义TCP协议 | 网关     | raw-tcp  | 服务ip:7654 | 127.0.0.1:8098 | plugin/tcp/      |
    | 自定义TCP协议 | 设备     | raw-tcp  | 服务ip:7654 | 127.0.0.1:8098 | plugin/tcp/      |
### SQL导入
在数据库执行以下sql
  ```sql
  INSERT INTO public.tp_protocol_plugin
  (id, "name", protocol_type, access_address, http_address, sub_topic_prefix, created_at, description, device_type)
  VALUES('de497b74-1bb6-2fc8-237b-75199304ba78', '自定义TCP协议', 'raw-tcp', '服务ip:7654', '127.0.0.1:8098', 'plugin/tcp/', 1670812659, 'docker部署不包含tcp协议插件服务,可根据文档自行部署', '2');
  INSERT INTO public.tp_protocol_plugin
  (id, "name", protocol_type, access_address, http_address, sub_topic_prefix, created_at, description, device_type)
  VALUES('aea3b83a-284d-5738-6d0f-94fc73220c33', '官方TCP协议', 'tcp', '服务ip:7653', '127.0.0.1:8000', 'plugin/tcp/', 1670813735, 'docker部署不包含tcp协议插件服务,可根据文档自行部署', '1');
  INSERT INTO public.tp_protocol_plugin
  (id, "name", protocol_type, access_address, http_address, sub_topic_prefix, created_at, description, device_type)
  VALUES('95b7c0b6-5c5b-4b45-c9ea-5bebda5a48ec', '官方TCP协议', 'tcp', '服务ip:7653', '127.0.0.1:8000', 'plugin/tcp/', 1670813749, 'docker部署不包含tcp协议插件服务,可根据文档自行部署', '2');
  INSERT INTO public.tp_protocol_plugin
  (id, "name", protocol_type, access_address, http_address, sub_topic_prefix, created_at, description, device_type)
  VALUES('95c957bc-a53b-6445-e882-1973bb546b12', '自定义TCP协议', 'raw-tcp', '服务ip:7654', '127.0.0.1:8098', 'plugin/tcp/', 1670809899, 'docker部署不包含tcp协议插件服务,可根据文档自行部署', '1');

  INSERT INTO public.tp_dict
  (id, dict_code, dict_value, "describe", created_at)
  VALUES('fad00d07-63c7-2685-1ee7-3e92d0142c88', 'DRIECT_ATTACHED_PROTOCOL', 'raw-tcp', '自定义TCP协议', 1670809899);
  INSERT INTO public.tp_dict
  (id, dict_code, dict_value, "describe", created_at)
  VALUES('9663bb03-4881-1965-5cf5-17341a4db761', 'GATEWAY_PROTOCOL', 'raw-tcp', '自定义TCP协议', 1670812659);
  INSERT INTO public.tp_dict
  (id, dict_code, dict_value, "describe", created_at)
  VALUES('b9249215-09a2-0298-02c2-0d9085fc40d2', 'DRIECT_ATTACHED_PROTOCOL', 'tcp', '官方TCP协议', 1670813735);
  INSERT INTO public.tp_dict
  (id, dict_code, dict_value, "describe", created_at)
  VALUES('25074e80-b7ca-99a3-e1f7-2fec7ec31b24', 'GATEWAY_PROTOCOL', 'tcp', '官方TCP协议', 1670813749);
  ```
## Installation
```bash
git clone https://github.com/sllt/tp-tcp-plugin.git

mv conf/config.example.yaml conf/config.yaml

cd tp-tcp-plugin && go build -o server cmd/tcp-protoc.go
```

## Usage
```bash
./server
```

## Configuration
see [config.yaml](conf/config.yaml)


## The ThingsPanel TCP Protocol
before a client connect to the server, you must create an iot device in thingsPanel and get the access token.
then the client sends the following structure:
```html
+---------+---------+----------------+----------+----------+----------------+
|  IDENT  |  IDENT  |     TYPE       |  CMD     |  LENGTH  |     PAYLOAD    |
+---------+---------+----------------+----------+----------+----------------+
|   'T'   |  'P'    |       1 byte   |  1 byte  |    4     |     Variable   |
+---------+---------+----------------+----------+----------+----------------+
```

where:
* TYPE:
  * 0x0: data packet
  * 0x1: heartbeat packet
* CMD:
  * 0x0: device auth
  * 0x1: publish attributes
  * 0x2: push events
  * ...
* LENGTH: the length of the payload
* PAYLOAD: the payload
  * if the CMD is 0x0, the payload is the access token
  * if the CMD is 0x1, the payload is the attributes in json format
  * if the CMD is 0x2, the payload is the events in json format

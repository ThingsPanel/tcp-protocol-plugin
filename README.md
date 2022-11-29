ThingsPanel TCP Plugin
======================

This plugin allows you to connect to a TCP server and forward data to thingsPanel MQTT server.

# Installation
```bash
git clone https://github.com/sllt/tp-tcp-plugin.git

cd tp-tcp-plugin && go build -o server cmd/tcp-protoc.go
```

# Usage
```bash
./server
```

# Configuration
see [config.yaml](conf/config.yaml)


# The ThingsPanel TCP Protocol
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



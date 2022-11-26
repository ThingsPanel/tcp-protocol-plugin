package model

import "encoding/binary"

type Packet struct {
	Ident1  byte   // 'T'
	Ident2  byte   // 'P'
	Flag    byte   // 0x0: data packet, 0x1: heartbeat packet
	Cmd     byte   // cmd, 0x0: auth, 0x1: publish attributes, 0x2: publish events
	Length  uint32 // payload length
	Payload []byte // payload
}

func (p *Packet) Parse(data []byte) error {
	p.Ident1 = data[0]
	p.Ident2 = data[1]
	p.Flag = data[2]
	p.Cmd = data[3]
	p.Length = binary.BigEndian.Uint32(data[4:8])
	p.Payload = data[8:]

	return nil
}

func (p *Packet) Serialize() []byte {
	p.EvalPayloadLength()

	data := make([]byte, 0)
	data = append(data, p.Ident1)
	data = append(data, p.Ident2)
	data = append(data, p.Flag)
	data = append(data, p.Cmd)
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(p.Payload)))
	data = append(data, length...)
	data = append(data, p.Payload...)
	return data
}

func (p *Packet) IsHeartbeat() bool {
	return p.Flag == 0x1
}

func (p *Packet) IsPublishAttributes() bool {
	return p.Cmd == 0x1
}

func (p *Packet) IsPublishEvents() bool {
	return p.Cmd == 0x2
}

func (p *Packet) IsDataPacket() bool {
	return p.Flag == 0x0
}

func (p *Packet) IsAuthPacket() bool {
	return p.Cmd == 0x0
}

func (p *Packet) EvalPayloadLength() {
	p.Length = uint32(len(p.Payload))
}

func BuildAuthPacket(accessToken string) *Packet {
	return &Packet{
		Ident1:  'T',
		Ident2:  'P',
		Flag:    0x0,
		Cmd:     0x0,
		Payload: []byte(accessToken),
	}
}

func BuildPublishAttributesPacket(payload []byte) *Packet {
	return &Packet{
		Ident1:  'T',
		Ident2:  'P',
		Flag:    0x0,
		Cmd:     0x1,
		Payload: payload,
	}
}

package protobuf

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

var errWrongValueType = errors.New("protobuf: convert on wrong type value")

// // Processor .
// // -------------------------
// // | id | protobuf message |
// // -------------------------
// type Processor struct {
// 	littleEndian bool
// 	msgInfo      []*MsgInfo
// 	msgID        map[reflect.Type]uint16
// }

// // MsgInfo .
// type MsgInfo struct {
// 	msgType reflect.Type
// 	//msgRouter     *chanrpc.Server
// 	msgHandler MsgHandler
// 	//msgRawHandler MsgHandler
// }

// // MsgHandler .
// type MsgHandler func([]interface{})

// // // MsgRaw .
// // type MsgRaw struct {
// // 	msgID      uint16
// // 	msgRawData []byte
// // }

// // NewProcessor .
// func NewProcessor() *Processor {
// 	p := new(Processor)
// 	p.littleEndian = false
// 	p.msgID = make(map[reflect.Type]uint16)
// 	return p
// }

// // SetByteOrder It's dangerous to call the method on routing or marshaling (unmarshaling)
// func (p *Processor) SetByteOrder(littleEndian bool) {
// 	p.littleEndian = littleEndian
// }

// // Register It's dangerous to call the method on routing or marshaling (unmarshaling)
// func (p *Processor) Register(msg proto.Message) uint16 {
// 	msgType := reflect.TypeOf(msg)
// 	if msgType == nil || msgType.Kind() != reflect.Ptr {
// 		llog.Fatal("protobuf message pointer required")
// 	}
// 	if _, ok := p.msgID[msgType]; ok {
// 		llog.Fatal("message %s is already registered", msgType)
// 	}
// 	if len(p.msgInfo) >= math.MaxUint16 {
// 		llog.Fatal("too many protobuf messages (max = %v)", math.MaxUint16)
// 	}

// 	i := new(MsgInfo)
// 	i.msgType = msgType
// 	p.msgInfo = append(p.msgInfo, i)
// 	id := uint16(len(p.msgInfo) - 1)
// 	p.msgID[msgType] = id
// 	return id
// }

// // // SetRouter It's dangerous to call the method on routing or marshaling (unmarshaling)
// // func (p *Processor) SetRouter(msg proto.Message, msgRouter *chanrpc.Server) {
// // 	msgType := reflect.TypeOf(msg)
// // 	id, ok := p.msgID[msgType]
// // 	if !ok {
// // 		llog.Fatal("message %s not registered", msgType)
// // 	}

// // 	p.msgInfo[id].msgRouter = msgRouter
// // }

// // SetHandler It's dangerous to call the method on routing or marshaling (unmarshaling)
// func (p *Processor) SetHandler(msg proto.Message, msgHandler MsgHandler) {
// 	msgType := reflect.TypeOf(msg)
// 	id, ok := p.msgID[msgType]
// 	if !ok {
// 		llog.Fatal("message %s not registered", msgType)
// 	}

// 	p.msgInfo[id].msgHandler = msgHandler
// }

// // // SetRawHandler It's dangerous to call the method on routing or marshaling (unmarshaling)
// // func (p *Processor) SetRawHandler(id uint16, msgRawHandler MsgHandler) {
// // 	if id >= uint16(len(p.msgInfo)) {
// // 		llog.Fatal("message id %v not registered", id)
// // 	}

// // 	p.msgInfo[id].msgRawHandler = msgRawHandler
// // }

// // Route goroutine safe
// func (p *Processor) Route(msg interface{}, userData interface{}) error {
// 	// raw
// 	// if msgRaw, ok := msg.(MsgRaw); ok {
// 	// 	if msgRaw.msgID >= uint16(len(p.msgInfo)) {
// 	// 		return fmt.Errorf("message id %v not registered", msgRaw.msgID)
// 	// 	}
// 	// 	i := p.msgInfo[msgRaw.msgID]
// 	// 	if i.msgRawHandler != nil {
// 	// 		i.msgRawHandler([]interface{}{msgRaw.msgID, msgRaw.msgRawData, userData})
// 	// 	}
// 	// 	return nil
// 	// }

// 	// protobuf
// 	msgType := reflect.TypeOf(msg)
// 	id, ok := p.msgID[msgType]
// 	if !ok {
// 		return fmt.Errorf("message %s not registered", msgType)
// 	}
// 	i := p.msgInfo[id]
// 	if i.msgHandler != nil {
// 		i.msgHandler([]interface{}{msg, userData})
// 	}
// 	// if i.msgRouter != nil {
// 	// 	i.msgRouter.Go(msgType, msg, userData)
// 	// }
// 	return nil
// }

// // Unmarshal goroutine safe
// func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
// 	if len(data) < 2 {
// 		return nil, errors.New("protobuf data too short")
// 	}

// 	// id
// 	var id uint16
// 	if p.littleEndian {
// 		id = binary.LittleEndian.Uint16(data)
// 	} else {
// 		id = binary.BigEndian.Uint16(data)
// 	}
// 	if id >= uint16(len(p.msgInfo)) {
// 		return nil, fmt.Errorf("message id %v not registered", id)
// 	}

// 	// msg
// 	i := p.msgInfo[id]
// 	if i.msgRawHandler != nil {
// 		return MsgRaw{id, data[2:]}, nil
// 	}
// 	msg := reflect.New(i.msgType.Elem()).Interface()
// 	return msg, proto.UnmarshalMerge(data[2:], msg.(proto.Message))

// }

// // Marshal goroutine safe
// func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
// 	msgType := reflect.TypeOf(msg)

// 	// id
// 	_id, ok := p.msgID[msgType]
// 	if !ok {
// 		err := fmt.Errorf("message %s not registered", msgType)
// 		return nil, err
// 	}

// 	id := make([]byte, 2)
// 	if p.littleEndian {
// 		binary.LittleEndian.PutUint16(id, _id)
// 	} else {
// 		binary.BigEndian.PutUint16(id, _id)
// 	}

// 	// data
// 	data, err := proto.Marshal(msg.(proto.Message))
// 	return [][]byte{id, data}, err
// }

// // Range goroutine safe
// func (p *Processor) Range(f func(id uint16, t reflect.Type)) {
// 	for id, i := range p.msgInfo {
// 		f(uint16(id), i.msgType)
// 	}
// }

//Encode Protobuf编码
func Encode(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, errWrongValueType
	}
	return proto.Marshal(pb)
}

//Decode Protobuf解码
func Decode(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return errWrongValueType
	}
	return proto.Unmarshal(data, pb)
}

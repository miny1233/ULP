package ulp
import "unsafe"
const Version uint8=0

type Kp struct {
	 _version uint8
	 _type uint8    
}
type Dp struct{
	_version uint8
	_type uint8
	id uint16
	len uint16
	data []byte
}
type Ap struct{
	_version uint8
	_type uint8
	len uint16
    UUID []byte
}
type Am struct{
	_version uint8
	_msg uint8
}
type Cp struct{
	_version uint8
	_type uint8
	_ip [4]uint8
	_port uint16
}
type Cm struct{
	_version uint8
	_type uint8
	_ip [4]uint8
	_port uint16
	id uint16
}

func Authentication_and_Conn(UUID string,send chan []byte) int{
    var uid =[]byte(UUID)
	p:=Ap{
        _version: Version,
		_type:0,
		len: uint16(len(uid)),
		UUID: uid,
	}
	send <- *(*[]byte)(unsafe.Pointer(&p))
	return -1
}

func KeepCon(send chan []byte){
   p := Kp{
	   _version: Version,
	   _type: 1,
   }
   pa := *(*[]byte)(unsafe.Pointer(&p))
   for{
     send <- pa
   }
}

func CreatConn(ip [4]uint8,port uint16,send chan []byte)  int{
	p:=Cp{
		_version:Version,
		_type:2,
		_ip:ip,
		_port:port,
	}
	send <- *(*[]byte)(unsafe.Pointer(&p))
	return -1
}
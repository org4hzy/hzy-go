package util

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"log"
	"net"
)

// Base64 return base64 string
func Base64(d string) string {
	en := base64.URLEncoding.EncodeToString([]byte(d))

	dec, err := base64.URLEncoding.DecodeString(en)
	if err != nil {
		log.Println("decode error:", err)
	}
	log.Printf("%s base64(%s, %s)", d, en, dec)
	return en
}

//IP2int ip string to int
func IP2int(ip string) uint32 {
	nIP := net.ParseIP(ip)

	if len(nIP) == 16 {
		return binary.BigEndian.Uint32(nIP[12:16])
	}
	return binary.BigEndian.Uint32(nIP)
}

//Int2IP ip int to string
func Int2IP(nn uint32) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)

	return ip.String()
}

//ToBytes data to bytes
func ToBytes(data interface{}) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, data)

	return buf.Bytes()
}

//ToData bytes to data
func ToData(b []byte) (data interface{}) {
	binary.Read(bytes.NewReader(b), binary.LittleEndian, data)
	return
}

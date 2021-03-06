package connect

import (
	"io"
	"github.com/LilyPad/GoLilyPad/packet"
)

type RequestGetDetails struct {

}

func NewRequestGetDetails() (this *RequestGetDetails) {
	this = new(RequestGetDetails)
	return
}

func (this *RequestGetDetails) Id() int {
	return REQUEST_GET_DETAILS
}

type requestGetDetailsCodec struct {

}

func (this *requestGetDetailsCodec) Decode(reader io.Reader, util []byte) (request Request, err error) {
	request = new(RequestGetDetails)
	return
}

func (this *requestGetDetailsCodec) Encode(writer io.Writer, util []byte, request Request) (err error) {
	return
}

type ResultGetDetails struct {
	Ip string
	Port uint16
	Motd string
	Version string
}

func NewResultGetDetails(ip string, port uint16, motd string, version string) (this *ResultGetDetails) {
	this = new(ResultGetDetails)
	this.Ip = ip
	this.Port = port
	this.Motd = motd
	this.Version = version
	return
}

func (this *ResultGetDetails) Id() int {
	return REQUEST_GET_DETAILS
}

type resultGetDetailsCodec struct {

}

func (this *resultGetDetailsCodec) Decode(reader io.Reader, util []byte) (result Result, err error) {
	resultGetDetails := new(ResultGetDetails)
	resultGetDetails.Ip, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	resultGetDetails.Port, err = packet.ReadUint16(reader, util)
	if err != nil {
		return
	}
	resultGetDetails.Motd, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	resultGetDetails.Version, err = packet.ReadString(reader, util)
	if err != nil {
		return
	}
	result = resultGetDetails
	return
}

func (this *resultGetDetailsCodec) Encode(writer io.Writer, util []byte, result Result) (err error) {
	resultGetDetails := result.(*ResultGetDetails)
	err = packet.WriteString(writer, util, resultGetDetails.Ip)
	if err != nil {
		return
	}
	err = packet.WriteUint16(writer, util, resultGetDetails.Port)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, resultGetDetails.Motd)
	if err != nil {
		return
	}
	err = packet.WriteString(writer, util, resultGetDetails.Version)
	return
}

package ssrdecoder

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Decode subscription decoder
// ss://
// ssr://
// vmess://
// trojan://
func Decode(str string) {

	return
}

type ssConf struct {
	Name       string `json:"-"           yaml:"-"`
	Server     string `json:"server"      yaml:"server"`
	ServerPort int    `json:"server_port" yaml:"server_port"`
	Password   string `json:"password"    yaml:"password"`
	LocalPort  int    `json:"local_port"  yaml:"local_port"`
	Timeout    int    `json:"timeout"     yaml:"timeout"`
	Method     string `json:"method"      yaml:"method"`
}

func ssParser(str string) (ss *ssConf, err error) {
	urlDecodedStr, err := url.QueryUnescape(str)
	if err != nil {
		return
	}

	ss = &ssConf{
		LocalPort: 1080,
		Timeout:   300,
	}

	nameSpilts := strings.Split(urlDecodedStr, "#")
	if len(nameSpilts) == 2 {
		ss.Name = nameSpilts[1]
	}

atLabel:
	atSplits := strings.Split(nameSpilts[0], "@")
	if len(atSplits) == 2 {
		host, port, er := parseHostAndPort(atSplits[1])
		if er != nil {
			err = fmt.Errorf("parseHostAndPort failed %#v", er)
			return
		}

		ss.Server = host
		ss.ServerPort = port

	}

	colonSplits := strings.Split(atSplits[0], ":")
	if len(colonSplits) == 2 {
		ss.Method = colonSplits[0]
		ss.Password = colonSplits[1]
	} else {
		decoded, er := base64Decode(atSplits[0])
		if er != nil {
			ss.Method = colonSplits[0]
			return
		}
		nameSpilts[0] = decoded
		goto atLabel
	}

	return
}

func base64Decode(str string) (decoded string, err error) {
	b, er := base64.URLEncoding.DecodeString(base64Padding(str))
	if er != nil {
		err = fmt.Errorf("base64 decode failed %#v", er)
		return
	}

	decoded = string(b)
	return
}

func base64Padding(s string) string {
	if i := len(s) % 4; i != 0 {
		s += strings.Repeat("=", 4-i)
	}
	return s
}

func parseHostAndPort(addr string) (host string, port int, err error) {
	colonSplits := strings.Split(addr, ":")
	host = colonSplits[0]
	port = 80
	if len(colonSplits) == 2 {
		port, err = strconv.Atoi(colonSplits[1])
	}

	return
}

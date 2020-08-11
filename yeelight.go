package yeelight

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	discoverMSG = "M-SEARCH * HTTP/1.1\r\n HOST:239.255.255.250:1982\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n"

	// timeout value for TCP and UDP commands
	timeout = time.Second * 3

	//SSDP discover address
	ssdpAddr = "239.255.255.250:1982"

	//CR-LF delimiter
	crlf = "\r\n"
)

type EffectType string

const (
	Smooth EffectType = "smooth"
	Sudden EffectType = "sudden"
)

type (
	PropsResult struct {
		ID     int
		Result map[string]string
		Error  *Error
	}

	// Notification represents notification response
	Notification struct {
		Method string            `json:"method"`
		Params map[string]string `json:"params"`
	}

	//Error struct represents error part of response
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

//Yeelight represents device
type YeelightConfig struct {
	Ip     string
	Effect EffectType
}

//Yeelight represents device
type Yeelight struct {
	ip     string
	addr   string
	effect EffectType
	cmdId  int
}

func New(config YeelightConfig) *Yeelight {
	if config.Ip == "" {
		log.Fatalln("Please, add bulb ip to yeelight config")
	}

	y := &Yeelight{
		ip:    config.Ip,
		addr:  fmt.Sprintf("%s:55443", config.Ip),
		cmdId: 0,
	}

	if config.Effect != "" {
		y.effect = config.Effect
	} else {
		y.effect = Smooth
	}

	return y
}

//Discover discovers device in local network via ssdp
func Discover() (*Yeelight, error) {
	var err error

	ssdp, _ := net.ResolveUDPAddr("udp4", ssdpAddr)
	c, _ := net.ListenPacket("udp4", ":0")
	socket := c.(*net.UDPConn)
	socket.WriteToUDP([]byte(discoverMSG), ssdp)
	socket.SetReadDeadline(time.Now().Add(timeout))

	rsBuf := make([]byte, 1024)
	size, _, err := socket.ReadFromUDP(rsBuf)
	if err != nil {
		return nil, errors.New("no devices found")
	}
	rs := rsBuf[0:size]
	addr := parseAddr(string(rs))
	fmt.Printf("Device with ip %s found\n", addr)

	return New(YeelightConfig{Ip: addr}), nil
}

func (y *Yeelight) Discover() (*YeelightParams, error) {
	var err error

	addr := fmt.Sprintf("%s:1982", y.ip)
	msg := fmt.Sprintf("M-SEARCH * HTTP/1.1\r\n HOST:%s\r\n MAN:\"ssdp:discover\"\r\n ST:wifi_bulb\r\n", addr)

	ssdp, _ := net.ResolveUDPAddr("udp4", ssdpAddr)
	c, _ := net.ListenPacket("udp4", ":0")

	socket := c.(*net.UDPConn)
	socket.WriteToUDP([]byte(msg), ssdp)
	socket.SetReadDeadline(time.Now().Add(timeout))

	rsBuf := make([]byte, 1024)
	size, _, err := socket.ReadFromUDP(rsBuf)
	if err != nil {
		return nil, errors.New("no devices found")
	}
	rs := rsBuf[0:size]

	params := parseAnswer(string(rs))
	return params, nil
}

func (y *Yeelight) TurnOn() (*CommandResult, error) {
	return y.ExecuteCommand("set_power", "on")
}

func (y *Yeelight) TurnOff() (*CommandResult, error) {
	return y.ExecuteCommand("set_power", "off")
}

func (y *Yeelight) StartFlow(flow *Flow) (*CommandResult, error) {
	res, err := y.GetProps([]string{"power"})
	if err != nil {
		return nil, err
	}
	power := res.Result["power"]
	if power != "on" {
		_, err = y.TurnOn()
		if err != nil {
			return nil, err
		}
	}
	params := flow.AsStartParams()
	return y.ExecuteCommand("start_cf", params)
}

func (y *Yeelight) StopFlow() (*CommandResult, error) {
	return y.ExecuteCommand("stop_cf", "")
}

func (y *Yeelight) GetProps(props []string) (*PropsResult, error) {
	res, err := y.ExecuteCommand("get_prop", props)
	if err != nil {
		return nil, err
	}

	propsMap := make(map[string]string)

	for i, val := range res.Result {
		key := props[i]
		propsMap[key] = fmt.Sprintf("%v", val)
	}

	return &PropsResult{ID: res.ID, Error: res.Error, Result: propsMap}, nil
}

// Listen connects to device and listens for NOTIFICATION events
func (y *Yeelight) Listen() (<-chan *Notification, chan<- struct{}, error) {
	var err error
	notifCh := make(chan *Notification)
	done := make(chan struct{}, 1)

	conn, err := net.DialTimeout("tcp", y.addr, time.Second*3)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot connect to %s. %s", y.addr, err)
	}

	fmt.Println("Connection established")
	go func(c net.Conn) {
		//make sure connection is closed when method returns
		defer closeConnection(conn)

		connReader := bufio.NewReader(c)
		for {
			select {
			case <-done:
				return
			default:
				data, err := connReader.ReadString('\n')
				if nil == err {
					var rs Notification
					fmt.Println(data)
					json.Unmarshal([]byte(data), &rs)
					select {
					case notifCh <- &rs:
					default:
						fmt.Println("Channel is full")
					}
				}
			}

		}

	}(conn)

	return notifCh, done, nil
}




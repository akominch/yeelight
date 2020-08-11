package yeelight

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"time"
)

//Command represents COMMAND request to Yeelight device
type Command struct {
	ID     int           `json:"id"`
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

// CommandResult represents response from Yeelight device
type CommandResult struct {
	ID     int           `json:"id"`
	Result []interface{} `json:"result,omitempty"`
	Error  *Error        `json:"error,omitempty"`
}

func (y *Yeelight) ExecuteCommand(name string, params ...interface{}) (*CommandResult, error) {
	return y.execute(y.newCommand(name, params))
}

func (y *Yeelight) newCommand(name string, params []interface{}) *Command {
	if len(params) > 0 {
		switch v := params[0].(type) {
		case []interface{}:
			params = v
		case []string:
			s := make([]interface{}, len(v))
			for i, val := range v {
				s[i] = val
			}
			params = s
		default:
		}
	}

	return &Command{
		Method: name,
		ID:     y.getCmdId(),
		Params: params,
	}
}

func (y *Yeelight) execute(cmd *Command) (*CommandResult, error) {
	conn, err := net.Dial("tcp", y.addr)
	if nil != err {
		return nil, fmt.Errorf("cannot open connection to %s. %s", y.addr, err)
	}
	//time.Sleep(time.Second)
	conn.SetReadDeadline(time.Now().Add(timeout))

	//write request/command
	b, _ := json.Marshal(cmd)
	fmt.Fprint(conn, string(b)+crlf)

	//wait and read for response
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("cannot read command result %s", err)
	}
	var rs CommandResult
	err = json.Unmarshal([]byte(res), &rs)
	if nil != err {
		return nil, fmt.Errorf("cannot parse command result %s", err)
	}
	if nil != rs.Error {
		return nil, fmt.Errorf("command execution error. Code: %d, Message: %s", rs.Error.Code, rs.Error.Message)
	}
	return &rs, nil
}

func (y *Yeelight) getCmdId() int {
	if y.cmdId == math.MaxInt32 {
		y.cmdId = 0
	}
	currentId := y.cmdId
	y.cmdId += 1
	return currentId
}
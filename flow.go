package yeelight

import (
	"strings"
)

type Action int8

const (
	Recover Action = 0
	Stay = 1
	Off = 2
)

type Flow struct {
	count int
	action Action
	transitions []string
}

func NewFlow(count int, action Action, transitions []string) *Flow {
	return &Flow{
		count: count,
		action: action,
		transitions: transitions,
	}
}

func (flow *Flow) AsStartParams() []interface{} {
	count := flow.count * len(flow.transitions)
	expr := strings.Join(flow.transitions, ",")

	return []interface{}{count, flow.action, expr}
}
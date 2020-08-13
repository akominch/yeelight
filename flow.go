package yeelight

import (
	t "github.com/akominch/yeelight/transitions"
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
	transitions []t.Transition
}

func NewFlow(count int, action Action, transitions []t.Transition) *Flow {
	return &Flow{
		count: count,
		action: action,
		transitions: transitions,
	}
}

func (flow *Flow) AsStartParams() []interface{} {
	count := flow.count * len(flow.transitions)

	var strTransitions []string

	for _, t := range flow.transitions {
		str := t.AsYeelightParams()
		strTransitions = append(strTransitions, str)
	}

	expr := strings.Join(strTransitions, ",")

	return []interface{}{count, flow.action, expr}
}
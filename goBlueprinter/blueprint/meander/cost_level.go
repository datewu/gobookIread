package meander

import (
	"errors"
	"strings"
)

// Cost is the enum type
type Cost int8

// Cost1 Cost5 are enum elements
const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (level Cost) String() string {
	for k, v := range costStrings {
		if level == v {
			return k
		}
	}
	return "invalid"
}

// ParseCost parse string to Cost
func ParseCost(s string) Cost {
	return costStrings[s]
}

// CostRange is a enum set
type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

// ParseCostRange lol
func ParseCostRange(s string) (CostRange, error) {
	var r CostRange
	segs := strings.Split(s, "...")
	if len(segs) != 2 {
		return r, errors.New("invalid cost range")
	}
	r.From = ParseCost(segs[0])
	r.To = ParseCost(segs[1])
	return r, nil
}

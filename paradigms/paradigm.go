package paradigms

import (
	"fmt"
	"math"
	"math/big"

	"github.com/turnon/elastalarm/response"
)

type Paradigm interface {
	Template() string
	Found(resp *response.Response) (bool, *string)
}

func Names(name string) Paradigm {
	switch name {
	case "percentage":
		return &Percentage{}
	case "count":
		return &Count{}
	case "spike":
		return &Spike{}
	default:
		return nil
	}
}

type Match struct {
	Gt, Lt *float64
	Not    bool
}

func (m *Match) match(v *big.Float) (bool, string) {
	result := true
	verb := "is"

	left := m.leftBound()
	if v.Cmp(configValue(&left)) != 1 {
		result = result && false
	}

	right := m.rightBound()
	if v.Cmp(configValue(&right)) != -1 {
		result = result && false
	}

	if m.Not {
		result = !result
		verb = "is not"
	}

	return result, fmt.Sprintf("%s between (%f, %f)", verb, left, right)
}

func (m *Match) leftBound() float64 {
	if m.Gt == nil {
		return math.Inf(-1)
	}
	return *m.Gt
}

func (m *Match) rightBound() float64 {
	if m.Lt == nil {
		return math.Inf(1)
	}
	return *m.Lt
}

func configValue(v *float64) *big.Float {
	return big.NewFloat(*v)
}

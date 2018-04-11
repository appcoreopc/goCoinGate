package goCoinGate

import "time"

type Gotime struct {

}

func (g Gotime) ToEpoch() int64 {
	return time.Now().UnixNano() / 1000000
}
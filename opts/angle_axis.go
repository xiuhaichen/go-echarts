package opts

import "github.com/xiuhaichen/go-echarts/types"

type AngleAxis struct {
	PolarAxisBase
	Clockwise types.Bool `json:"clockwise,omitempty"`
}

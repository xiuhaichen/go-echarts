package opts

import "github.com/xiuhaichen/go-echarts/types"

func Bool(val bool) types.Bool {
	return &val
}

func Int(val int) types.Int {
	return &val
}

func Float(val float32) types.Float {
	return &val
}

func Str(val string) types.String {
	return types.String(val)
}

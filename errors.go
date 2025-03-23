package carbon

import (
	"fmt"
)

// returns a parse error.
// 解析错误
var parseError = func(value string) error {
	return fmt.Errorf("cannot parse %q as carbon, please make sure the value is valid", value)
}

// returns an empty timezone error.
// 空时区错误
var emptyTimezoneError = func() error {
	return fmt.Errorf("timezone cannot be empty")
}

// returns an invalid timezone error.
// 无效的时区错误
var invalidTimezoneError = func(timezone string) error {
	return fmt.Errorf("invalid timezone %q, please see the file %q for all valid timezones", timezone, "$GOROOT/lib/time/zoneinfo.zip")
}

// returns an invalid location error.
// 无效的地区错误
var invalidLocationError = func() error {
	return fmt.Errorf("invalid location, please make sure the location is valid")
}

// returns an empty duration error.
// 空的时长错误
var emptyDurationError = func() error {
	return fmt.Errorf("duration cannot be empty")
}

// returns an invalid duration error.
// 无效的时长错误
var invalidDurationError = func(duration string) error {
	return fmt.Errorf("invalid duration %q, please make sure the duration is valid", duration)
}

// returns an empty layout error.
// 空的布局模板错误
var emptyLayoutError = func() error {
	return fmt.Errorf("layout cannot be empty")
}

// returns an invalid layout error.
// 无效的布局模板错误
var invalidLayoutError = func(value, layout string) error {
	return fmt.Errorf("cannot parse string %q as carbon by layout %q, please make sure the value and layout match", value, layout)
}

// returns an empty format error.
// 空的格式模板错误
var emptyFormatError = func() error {
	return fmt.Errorf("format cannot be empty")
}

// returns an invalid format error.
// 无效的格式模板错误
var invalidFormatError = func(value, format string) error {
	return fmt.Errorf("cannot parse string %q as carbon by format %q, please make sure the value and format match", value, format)
}

// returns an empty week starts day error.
// 空的周起始日期错误
var emptyWeekStartsDayError = func() error {
	return fmt.Errorf("week start day cannot be empty")
}

// returns an invalid week starts day error.
// 无效的周起始日期错误
var invalidWeekStartsDayError = func(day string) error {
	return fmt.Errorf("invalid week start day %s, please make sure the day is valid", day)
}

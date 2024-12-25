package utils

import "strconv"

func GetInt64FromInterface(val interface{}) int64 {
	if val == nil {
		return 0
	}

	switch v := val.(type) {
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float64:
		return int64(v)
	case float32:
		return int64(v)
	case string:
		num, _ := strconv.ParseInt(val.(string), 10, 64)
		return num
	case bool:
		if v == true {
			return int64(1)
		} else {
			return int64(0)
		}
	}
	return 0
}

func GetStringFromInterface(val interface{}) string {
	if val == nil {
		return ""
	}

	switch v := val.(type) {
	case string:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case bool:
		return strconv.FormatBool(v)
	}
	return ""
}

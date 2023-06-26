package v1

import "reflect"

// DeepEqualsIgnoreValues compares two objects and returns true if they are equal WITHOUT comparing values.
func DeepEqualsIgnoreValues(obj1, obj2 interface{}) bool {
	if reflect.TypeOf(obj1) != reflect.TypeOf(obj2) {
		return false
	}

	switch obj1 := obj1.(type) {
	case map[string]interface{}:
		obj2, ok := obj2.(map[string]interface{})
		if !ok {
			return false
		}

		if len(obj1) != len(obj2) {
			return false
		}

		for key, value1 := range obj1 {
			value2, ok := obj2[key]
			if !ok {
				return false
			}

			if !DeepEqualsIgnoreValues(value1, value2) {
				return false
			}
		}

	case []interface{}:
		obj2, ok := obj2.([]interface{})
		if !ok {
			return false
		}

		if len(obj1) != len(obj2) {
			return false
		}

		for i := 0; i < len(obj1); i++ {
			if !DeepEqualsIgnoreValues(obj1[i], obj2[i]) {
				return false
			}
		}

	default:
		// Ignore the actual values, only compare keys
		return true
	}

	return true
}

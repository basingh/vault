// Code generated by "enumer -type=WhenInconsistentAction -trimprefix=WhenInconsistent"; DO NOT EDIT.

package cache

import (
	"fmt"
)

const _WhenInconsistentActionName = "FailRetryForward"

var _WhenInconsistentActionIndex = [...]uint8{0, 4, 9, 16}

func (i WhenInconsistentAction) String() string {
	if i < 0 || i >= WhenInconsistentAction(len(_WhenInconsistentActionIndex)-1) {
		return fmt.Sprintf("WhenInconsistentAction(%d)", i)
	}
	return _WhenInconsistentActionName[_WhenInconsistentActionIndex[i]:_WhenInconsistentActionIndex[i+1]]
}

var _WhenInconsistentActionValues = []WhenInconsistentAction{0, 1, 2}

var _WhenInconsistentActionNameToValueMap = map[string]WhenInconsistentAction{
	_WhenInconsistentActionName[0:4]:  0,
	_WhenInconsistentActionName[4:9]:  1,
	_WhenInconsistentActionName[9:16]: 2,
}

// WhenInconsistentActionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func WhenInconsistentActionString(s string) (WhenInconsistentAction, error) {
	if val, ok := _WhenInconsistentActionNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to WhenInconsistentAction values", s)
}

// WhenInconsistentActionValues returns all values of the enum
func WhenInconsistentActionValues() []WhenInconsistentAction {
	return _WhenInconsistentActionValues
}

// IsAWhenInconsistentAction returns "true" if the value is listed in the enum definition. "false" otherwise
func (i WhenInconsistentAction) IsAWhenInconsistentAction() bool {
	for _, v := range _WhenInconsistentActionValues {
		if i == v {
			return true
		}
	}
	return false
}

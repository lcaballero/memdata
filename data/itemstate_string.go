// Code generated by "stringer -type=ItemState"; DO NOT EDIT

package data

import "fmt"

const _ItemState_name = "InceptionInArchivedInProgressInBacklogCompleted"

var _ItemState_index = [...]uint8{0, 9, 19, 29, 38, 47}

func (i ItemState) String() string {
	i -= 1
	if i < 0 || i >= ItemState(len(_ItemState_index)-1) {
		return fmt.Sprintf("ItemState(%d)", i+1)
	}
	return _ItemState_name[_ItemState_index[i]:_ItemState_index[i+1]]
}

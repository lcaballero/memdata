// Code generated by "stringer -type=UserRole"; DO NOT EDIT

package data

import "fmt"

const _UserRole_name = "RootRoleNormalRole"

var _UserRole_index = [...]uint8{0, 8, 18}

func (i UserRole) String() string {
	i -= 1
	if i < 0 || i >= UserRole(len(_UserRole_index)-1) {
		return fmt.Sprintf("UserRole(%d)", i+1)
	}
	return _UserRole_name[_UserRole_index[i]:_UserRole_index[i+1]]
}

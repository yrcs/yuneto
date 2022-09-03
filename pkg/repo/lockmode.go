package repo

import "fmt"

type LockMode int32

const (
	// for update
	UPDATE LockMode = iota
	// lock in share mode
	SHARE
)

// Enum name map for LockMode.
var _LockModeMap = map[LockMode]string{
	UPDATE: "UPDATE",
	SHARE:  "SHARE",
}

// String implements the Stringer interface.
func (x LockMode) String() string {
	if str, ok := _LockModeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("LockMode(%d)", x)
}

// Enum value map for LockMode.
var _LockModeValue = map[string]LockMode{
	"UPDATE": UPDATE,
	"SHARE":  SHARE,
}

// ParseLockMode attempts to convert a string to a LockMode.
func ParseLockMode(name string) (LockMode, error) {
	if x, ok := _LockModeValue[name]; ok {
		return x, nil
	}
	return LockMode(0), fmt.Errorf("%s is not a valid LockMode", name)
}

func (x LockMode) Ptr() *LockMode {
	return &x
}

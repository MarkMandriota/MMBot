package error

import (
	"fmt"
)

type NotEnoughPermissions struct {
	permission int64
}

func (e *NotEnoughPermissions) Error() string {
	return fmt.Sprintf("Not enough permissions %d", e.permission)
}

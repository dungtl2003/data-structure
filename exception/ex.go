package exception

import "fmt"

type IndexOutOfBound struct {
	Start uint
	End   uint
}

func (e *IndexOutOfBound) Error() string {
	return fmt.Sprintf("The index is out of range [%d, %d]", e.Start, e.End)
}

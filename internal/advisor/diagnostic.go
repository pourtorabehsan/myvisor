package advisor

import "fmt"

type Diagnostic struct {
	ID          int
	Description string
	Data        string
	Error       error
}

func (c Diagnostic) String() string {
	if c.Error != nil {
		return fmt.Sprintf("[%d] %s\n%s\n", c.ID, c.Description, c.Error)
	}

	return fmt.Sprintf("[%d] %s\n%s\n", c.ID, c.Description, c.Data)
}

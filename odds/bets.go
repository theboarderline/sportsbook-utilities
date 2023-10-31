package odds

import "fmt"

func (o Outcome) Format() string {
	if o.Point > 0 {
		return fmt.Sprintf("%s +%.0f", o.Name, o.Point)
	}

	return fmt.Sprintf("%s %.0f", o.Name, o.Point)
}

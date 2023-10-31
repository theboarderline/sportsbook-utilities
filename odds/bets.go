package odds

import "fmt"

func (o Outcome) FormatPrice() string {
	if o.Price > 0 {
		return fmt.Sprintf("%s +%.0f", o.Name, o.Price)
	}

	return fmt.Sprintf("%s %.0f", o.Name, o.Price)
}

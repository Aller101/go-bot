package e

import "fmt"

func Wrap(op string, msg string, err error) error {
	return fmt.Errorf("(%s) %s: %w", op, msg, err)
}

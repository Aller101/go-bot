package e

import "fmt"

func Wrap(op string, msg string, err error) error {
	return fmt.Errorf("(%s) %s: %w", op, msg, err)
}

func WrapIfErr(op string, msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(op, msg, err)
}

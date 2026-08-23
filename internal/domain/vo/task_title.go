package vo

import (
	"errors"
	"strings"
)

var ErrInvalidTaskTitle = errors.New("title is required")

type TaskTitle string

func NewTaskTitle(value string) (TaskTitle, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return "", ErrInvalidTaskTitle
	}
	return TaskTitle(value), nil
}

func (t TaskTitle) String() string {
	return string(t)
}

package grifts

import (
	"github.com/JDWardle/buffalo_spelunking/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}

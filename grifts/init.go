package grifts

import (
	"github.com/balmanrawat/kopila-api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}

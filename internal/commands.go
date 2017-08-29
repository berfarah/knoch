package internal

import (
	"github.com/berfarah/knoch/internal/command"
	"github.com/berfarah/knoch/internal/git"
)

var Runner = command.NewRunner()
var Git = git.New()

package dialog

import (
	inkwell "git.teletypegames.org/engines/inkwell"
)

type (
	Dialog       = inkwell.Dialog
	Text         = inkwell.Text
	Action       = inkwell.Action
	Given        = inkwell.Given
	DialogNode   = inkwell.DialogNode
	DialogLine   = inkwell.DialogLine
	DialogChoice = inkwell.DialogChoice
)

var DialogManager = inkwell.NewManager[*Dialog]()

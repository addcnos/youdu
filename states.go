package youdu

type EnableState int

const (
	EnableForbidState      EnableState = -1
	EnableAuthorizedState  EnableState = 1
	EnableUnactivatedState EnableState = 0
)

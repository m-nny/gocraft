package datatypes

type State int

const (
	STATE_UNKNOWN State = iota
	STATE_HANDSHAKING
	STATE_STATUS
	STATE_LOGIN
	STATE_TRANSFER
)

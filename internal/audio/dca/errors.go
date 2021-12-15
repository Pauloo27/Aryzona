package dca

import (
	"github.com/Pauloo27/aryzona/internal/utils/errore"
)

var (
	ErrVoiceConnectionClosed = errore.Errore{ID: "VOICE_CONNECTION_CLOSED", Message: "Voice connection was closed"}
)
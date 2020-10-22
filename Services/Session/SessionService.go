package Session

import (
	"time"
)

type SessionService struct {
	MaxSessionLengthInSeconds int
}

var sessionTimer = make(map[string] time.Time)

func (s SessionService) HasSession(cardId string) bool {
	currentTime := time.Now()
	if timeSessionCreated, sessionExists := sessionTimer[cardId]; sessionExists {
		durationOfSession := currentTime.Sub(timeSessionCreated)
		if int(durationOfSession.Seconds()) < s.MaxSessionLengthInSeconds {
			return true
		}else{
			s.DestroySession(cardId)
		}
	}
	return false
}

func (s SessionService) CreateSession(cardId string) {
	sessionTimer[cardId] = time.Now()
}

func (s SessionService) DestroySession(cardId string) {
	delete (sessionTimer, cardId)
}



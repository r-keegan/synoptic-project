package Session

import (
	"time"
)

type SessionService struct {
	MaxSessionLengthInSeconds int
}

var sessionTimer = make(map[string]time.Time)

func (s SessionService) HasSession(cardId string) bool {
	currentTime := time.Now()
	if timeSessionCreated, sessionExists := sessionTimer[cardId]; sessionExists { //if a session with a time is in the map
		durationOfSession := currentTime.Sub(timeSessionCreated)            //Calculate length of time since session creation
		if int(durationOfSession.Seconds()) < s.MaxSessionLengthInSeconds { // if the session was created within the timeout period
			return true
		} else { // if the session was created over the timeout period long ago then destroy it
			s.DestroySession(cardId)
		}
	}
	return false
}

func (s SessionService) CreateSession(cardId string) {
	sessionTimer[cardId] = time.Now()
}

func (s SessionService) DestroySession(cardId string) {
	delete(sessionTimer, cardId)
}

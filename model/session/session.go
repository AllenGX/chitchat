package session

import "strconv"

type SessionManager struct {
	sessions []Session
}

type Session struct {
	Sion      Sesion
	SessionID int
}

type Sesion map[string]string

var sessionManager *SessionManager

func init() {
	var sessions []Session
	sessionManager = &SessionManager{sessions}
}

//Init GetSession
func Init(userID int) (session Sesion) {
	for k, v := range sessionManager.sessions {
		if v.SessionID == userID {
			return sessionManager.sessions[k].Sion

		}
	}
	sess := make(map[string]string)
	sess["userID"] = strconv.Itoa(userID)
	sessionManager.sessions = append(sessionManager.sessions, Session{
		SessionID: userID,
		Sion:      sess,
	})
	return sess
}

func AddSessionManager(userID int) {
	session := make(map[string]string)
	session["userID"] = strconv.Itoa(userID)
	sessionManager.sessions = append(sessionManager.sessions, Session{
		SessionID: userID,
		Sion:      session,
	})
}

func GetSessionManager(userID int) (Sesion, bool) {
	for k, v := range sessionManager.sessions {
		if v.SessionID == userID {
			return sessionManager.sessions[k].Sion, true
		}
	}
	return nil, false
}

func RemoveSessionManager(userID int) {
	for k, v := range sessionManager.sessions {
		if v.SessionID == userID {
			sessionManager.sessions = append(
				sessionManager.sessions[:k],
				sessionManager.sessions[k+1:]...)
			break
		}
	}
}

func (session Sesion) GetSession(key string) (value string) {
	return session[key]
}

func (session Sesion) SetSession(key string, value string) {
	session[key] = value
}

func (session Sesion) RemoveSession(key string) {
	if session[key] != "" {
		delete(session, "key")
	}
}

package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log"
	"time"
)

type Member struct {
	id        int
	email     string
	password  string
	firstName string
}

func (m *Member) Id() int {
	return m.id
}

func (m *Member) Email() string {
	return m.email
}

func (m *Member) Password() string {
	return m.password
}

func (m *Member) FirstName() string {
	return m.firstName
}

func (m *Member) SetId(value int) {
	m.id = value
}

func (m *Member) SetEmail(value string) {
	m.email = value
}

func (m *Member) SetPassword(value string) {
	m.password = value
}

func (m *Member) SetFirstName(value string) {
	m.firstName = value
}

type Session struct {
	id        int
	memberId  int
	sessionId string
}

func (s *Session) Id() int {
	return s.id
}

func (s *Session) MemberId() int {
	return s.memberId
}

func (s *Session) SessionId() string {
	return s.sessionId
}

func (s *Session) SetId(value int) {
	s.id = value
}

func (s *Session) SetMemberId(value int) {
	s.memberId = value
}

func (s *Session) SetSessionId(value string) {
	s.sessionId = value
}

func GetMember(email, password string) (Member, error) {
	log.Printf("Get member '%s' ('%s')", email, password)
	db, err := GetDBConnection()
	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		log.Printf("Encrypted password: %s", hex.EncodeToString(pwd[:]))
		row := db.QueryRow(`SELECT id, email, first_name
			FROM Member
			WHERE email = $1 AND password = $2`,
			email,
			hex.EncodeToString(pwd[:]),
		)
		result := Member{}
		err = row.Scan(&result.id, &result.email, &result.firstName)
		log.Printf("Err: %v", err)
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Unable to find Member with email: " + email)
		}
	} else {
		return Member{}, errors.New("Unable to get database connection")
	}
}

func CreateSession(m Member) (Session, error) {
	result := Session{}
	result.memberId = m.Id()
	sessionId := sha256.Sum256([]byte(m.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])

	db, err := GetDBConnection()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO Session
			(member_id, session_id)
			VALUES ($1, $2)
			RETURNING id`,
			m.Id(),
			result.sessionId,
		).Scan(&result.id)
		log.Print(err)
		if err == nil {
			return result, nil
		} else {
			return Session{}, errors.New("Unable to save session to database")
		}
	} else {
		return Session{}, errors.New("Unable to get database connection")
	}
}

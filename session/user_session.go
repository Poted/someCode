package session

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
)

// // data from DB
// const sessionID = "1"
// const userID = "1"
// const lang = "english"
// const currency = "storage url"
// const country = "images url"

type UserSession struct {
	Session session
	User    user
	Locale  locale
}

func (s UserSession) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

var (
	userSess *UserSession
	once     sync.Once
)

func StartSession() *UserSession {

	if userSess == nil {

		once.Do(func() {

			userSess = &UserSession{
				Session: session{ID: func() uuid.UUID {
					id, err := uuid.NewV7()
					if err != nil {
						fmt.Errorf("cannot create session ID")
						return uuid.Nil
					}
					return id
				}()},
				User: user{
					urls: urls{
						url: "",
					},
					doownload: download{},
				},
				Locale: locale{
					urls:     urls{},
					download: download{},
				},
			}
		})
	}
	return userSess
}

type urls struct {
	url string // private field for const/private data
}

type download struct {
	FileID int // public fields that may be changed freely
	File   []byte
}

type session struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Lang     string
	Currency string
	Country  string
}

type user struct {
	urls      urls // fields are not inherited
	doownload download
}

type locale struct {
	urls
	download
}

// getters and setters for encapsulated operations
func (s *session) SessionID() uuid.UUID {
	return s.ID
}

func (c *session) UpdateUrl(id string) *session {
	return c // put
}

func (s *user) UpdateUrl(id string) []byte { // Different return type for returning only

	s.urls.url = "new url with id: " + id

	fmt.Printf("s.urls.url: %v\n", s.urls.url)

	return []byte(id) // read modified only
}

func (u *urls) OtherFunction(number int) string {

	return fmt.Sprintf("soem: %d", 220+number)

}

func (d *download) FileDownload(id int) ([]byte, error) {

	fmt.Printf("downloading file with id: %d!\n", id)

	return d.File, d.saveFile(id)
}

func (d *download) saveFile(id int) error { // private method

	d.File = []byte{byte(id)}
	d.FileID = id

	fmt.Printf("file with id %d saved!\n", id)
	return nil
}

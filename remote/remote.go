package remote

import (
	"fmt"
	"sync"
)

const connection_url = "connection url"
const storage_url = "storage url"
const images_url = "images url"

type remote struct {
	Connection connection
	Storage    storage
	Images     images
}

var (
	Remote *remote
	once   sync.Once
)

func Connect() *remote {

	if Remote != nil {
		return Remote
	}

	once.Do(func() {

		Remote = &remote{

			Connection: connection{
				urls{
					Url: connection_url,
				},
			},

			Storage: storage{
				urls{
					Url: storage_url,
				},
				download{
					Download: fileDownload,
				},
			},
			Images: images{
				urls{
					Url: images_url,
				},
				download{
					Download: fileDownload,
				},
			},
		}
	})

	return Remote
}

type urls struct {
	Url string
}

type download struct {
	Download func(id int) error
}

type connection struct {
	urls
}

type storage struct {
	urls
	download
}

type images struct {
	urls
	download
}

func (c *connection) URLWithID(id string) *connection {
	c.urls.Url = "connection url with id: %s" + id
	return c
}

func (s *storage) StorageURLWithID(id string) string {
	return storage_url + id
}

func (u *urls) OtherFunction(number int) string {

	return fmt.Sprintf("soem: %d", 220+number)

}

func fileDownload(id int) error {

	fmt.Printf("file with %d downloaded!")
	return nil
}

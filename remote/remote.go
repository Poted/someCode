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

	if Remote == nil {

		once.Do(func() {

			Remote = &remote{

				Connection: connection{
					urls{
						url: connection_url,
					},
				},

				Storage: storage{
					urls{
						storage_url,
					},
					download{
						// FileID: Remote.Images.FileID, // access between fields in initialization? cause panic
					},
				},
				Images: images{
					urls{
						images_url,
					},
					download{},
				},
			}
		})
	}
	return Remote
}

type urls struct {
	url string // private field for const/private data
}

type download struct {
	FileID int // public fields that may be changed freely
	File   []byte
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

// getters and setters for encapsulated operations
func (c *connection) URL() string {
	return c.urls.url
}

func (c *connection) UpdateUrl(id string) *connection { // Same name different type and probably different logic
	c.urls.url = "connection url with id: " + id // this connection return is to add direct access to .Url() getter after calling this function
	return c                                     // put
}

func (s *storage) UpdateUrl(id string) string { // Different return type
	return storage_url + id // read modified only
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

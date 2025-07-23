package cloud

type Clouddb struct {
	filename string
}

func NewCloudDB(url string) *Clouddb {
	return &Clouddb{
		filename: url,
	}

}
func (db *Clouddb) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *Clouddb) Write(content []byte) {

}

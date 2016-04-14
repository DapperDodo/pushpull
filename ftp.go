package pushpull

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/jlaffaye/ftp"
)

/*
	Ftp is a simple PushPuller implementation that talks to a regular ftp server
*/
type Ftp struct {
	srvr, port, user, pass string
}

func NewFtp(s, p, u, w string) *Ftp {
	return &Ftp{
		srvr: s,
		port: p,
		user: u,
		pass: w,
	}
}

func (f *Ftp) Push(filepath string, data []byte) error {

	c, err := f.connect()
	if err != nil {
		return err
	}
	defer c.Quit()

	buf := bytes.NewBuffer(data)
	err = c.Stor(filepath, buf)
	if err != nil {
		return err
	}

	return nil
}

func (f *Ftp) Pull(filepath string) ([]byte, error) {

	c, err := f.connect()
	if err != nil {
		return nil, err
	}
	defer c.Quit()

	r, err := c.Retr(filepath)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (f *Ftp) connect() (*ftp.ServerConn, error) {

	c, err := ftp.DialTimeout(f.srvr+":"+f.port, 30*time.Second)
	if err != nil {
		return nil, err
	}

	err = c.Login(f.user, f.pass)
	if err != nil {
		return nil, err
	}

	return c, nil
}

package user

import (
	"io"
	"io/ioutil"
	"net/http"
)

type User struct {
	done <-chan struct{}
}

func New(done <-chan struct{}) *User {
	return &User{
		done: done,
	}
}

func (u User) Start(url string) <-chan int {
	status := make(chan int)
	go func() {
		for {
			s, err := u.send(url)
			if err != nil {
				continue
			}

			select {
			case <-u.done:
				close(status)
				return
			case status <- s:
			}
		}
	}()

	return status
}

func (u User) send(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	return resp.StatusCode, nil
}

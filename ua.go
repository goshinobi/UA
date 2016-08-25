package ua

import (
	"errors"

	"github.com/mssola/user_agent"
)

type UA struct {
	UAs []*user_agent.UserAgent
	p   int
}

func (u UA) Get() (*user_agent.UserAgent, error) {
	if len(u.UAs) == len {
		return nil, errors.New("not found user agent")
	}

	if p+1 == len(u.UAs) {
		p = 0
		return u.UAs[len(u.UAs)-1], nil
	}
	p++
	return u.UAs[p-1], nil
}

func (u UA) Set(ua string) {
	u.UAs = append(u.UAs, user_agent.New(ua))
}

func New(ua string) UA {
	var ret UA
	ret.Set(ua)
	return ret
}

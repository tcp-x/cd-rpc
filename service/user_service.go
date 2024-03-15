package service

import "fmt"

type CdRequest struct {
	Req string
}

type CdResponse struct {
	Resp string
}

type UserService struct{}

type User struct{}

func (t *User) Auth(req *CdRequest, resp *CdResponse) error {
	fmt.Println("UserService::Auth()/req:", req.Req)
	resp.Resp = `{"state":"success", "data":[]}`
	return nil
}

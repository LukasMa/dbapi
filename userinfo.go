package dbapi

import "net/http"

// The UserInfoService binds to the HTTP endpoints which belong to the userInfo resource.
type UserInfoService struct {
	client *Client
}

// UserInfo represents personal information of a user.
type UserInfo struct {
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	DateOfBirth string `json:"dateOfBirth,omitempty"`
	Gender      string `json:"gender,omitempty"`
}

// Get retrieves personal information (e.g. first name, family name date of
// birth) about the current user.
func (s *UserInfoService) Get() (*UserInfo, *Response, error) {
	u := "/userInfo"
	r := new(UserInfo)

	resp, err := s.client.Call(http.MethodGet, u, nil, r)
	return r, resp, err
}

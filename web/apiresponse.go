package web

import "fmt"

// apiResponse is a concrete implementation of web.APIResponse
type apiResponse struct {
	OK  bool   `json:"ok"`
	Err string `json:"error"`
	Raw []byte
}

// Success returns true if no error occured
func (t apiResponse) Success() bool {
	return t.OK
}

// Error returns an error containing details of the fault, if there is one
func (t apiResponse) Error() (err error) {
	if !t.Success() {
		err = fmt.Errorf(t.Err)
	}
	return
}

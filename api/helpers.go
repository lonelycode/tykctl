package api

import (
	"fmt"
	"github.com/TykTechnologies/tykctl/logger"
	"gopkg.in/resty.v1"
)

var log = logger.GetLogger("api.helpers")

type ServiceError struct {
	msg string
}

func (s *ServiceError) Error() string {
	return s.msg
}

func (s ServiceError) New(resp *resty.Response) error {
	return &ServiceError{
		msg: fmt.Sprintf("[API ERR]: status code %v, body: %v", resp.StatusCode(), resp.String()),
	}
}

func HandleResponseAndError(resp *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if resp.StatusCode() > 300 {
		return fmt.Errorf("[API ERR]: status code %v, body: %v", resp.StatusCode(), resp.String())
	}

	return nil
}

package dash

import (
	"github.com/TykTechnologies/tykctl/logger"
	"gopkg.in/resty.v1"
)

var log = logger.GetLogger("api.dash")

type Client struct {
	token string
	host  string
	api   *resty.Client
}

func New(host, token string) *Client {
	cl := &Client{token: token, host: host, api: resty.New()}
	err := cl.Init()
	if err != nil {
		log.Fatal(err)
	}

	return cl
}

func (c *Client) Init() error {
	c.api.OnBeforeRequest(func(rc *resty.Client, req *resty.Request) error {
		req.Header.Add("Authorization", c.token)
		return nil
	})

	return nil
}

func (c *Client) Api() *Api {
	apiClient := &Api{cl: c}
	return apiClient
}

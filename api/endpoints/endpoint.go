package endpoints

import (
	"fmt"
	"github.com/TykTechnologies/tykctl/logger"
	"net/url"
	"strings"
)

var log = logger.GetLogger("api.endpoint")

type Endpoint struct {
	u           *url.URL
	p           string
	params      map[string]string
	queries     map[string]string
	URLTemplate string
	basePath    string
}

func (e *Endpoint) P(name, value string) *Endpoint {
	v := fmt.Sprintf("{%s}", name)
	e.p = strings.Replace(e.p, v, value, 1)
	e.params[v] = value
	return e
}

func (e *Endpoint) Q(name, value string) *Endpoint {
	e.queries[name] = value
	return e
}

func (e *Endpoint) String() string {
	e.p = fmt.Sprintf("%s%s", e.basePath, e.p)

	for v, val := range e.params {
		e.p = strings.Replace(e.p, v, val, 1)
	}

	var err error
	e.u, err = url.Parse(e.p)
	if err != nil {
		log.Fatal(err)
	}

	for n, q := range e.queries {
		e.u.Query().Add(n, q)
	}

	log.Debug("calling: ", e.u.String())

	return e.u.String()
}

func (e *Endpoint) NewFrom() *Endpoint {
	ne := &Endpoint{
		URLTemplate: e.URLTemplate,
		u:           &url.URL{},
		p:           e.URLTemplate,
		queries:     map[string]string{},
		params:      map[string]string{},
	}
	return ne
}

func (e *Endpoint) Base(pth string) *Endpoint {
	n := e.NewFrom()
	n.basePath = pth
	return n
}

func EP(tpl string) *Endpoint {
	ne := &Endpoint{
		URLTemplate: tpl,
		u:           &url.URL{},
		p:           tpl,
		queries:     map[string]string{},
		params:      map[string]string{},
	}
	return ne
}

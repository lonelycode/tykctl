package dash

import (
	"fmt"
	"github.com/TykTechnologies/tyk/apidef"
	"github.com/TykTechnologies/tykctl/api"
)

type Api struct {
	cl *Client
}

func (a *Api) List() ([]*apidef.APIDefinition, error) {
	apiList := &APIDefinitions{}
	if err := api.HandleResponseAndError(a.cl.api.R().
		SetQueryParam("p", "-2").
		SetResult(apiList).
		Get(APIs.Base(a.cl.host).String())); err != nil {
		return nil, err
	}

	return apiList.AsStandardDef(), nil
}

func (a *Api) Search(query string) ([]*apidef.APIDefinition, error) {
	apiList := &APIDefinitions{}
	if err := api.HandleResponseAndError(a.cl.api.R().
		SetResult(&apiList).
		SetQueryParam("q", query).
		Get(APISearch.Base(a.cl.host).String())); err != nil {
		return nil, err
	}

	return apiList.AsStandardDef(), nil
}

func (a *Api) Fetch(id string) (*apidef.APIDefinition, error) {
	apiDef := &DashApiDefinition{}
	if err := api.HandleResponseAndError(a.cl.api.R().
		SetResult(apiDef).
		Get(API.Base(a.cl.host).P("id", id).String())); err != nil {
		return nil, err
	}

	return apiDef.APIDefinition, nil
}

func (a *Api) Update(dbId string, def *apidef.APIDefinition) error {
	apiDef := &DashApiDefinition{}
	if err := api.HandleResponseAndError(a.cl.api.R().
		SetResult(apiDef).
		Get(API.Base(a.cl.host).P("id", dbId).String())); err != nil {
		return err
	}

	respData := &APIResp{}
	apiDef.APIDefinition = def
	if err := api.HandleResponseAndError(a.cl.api.R().
		SetBody(apiDef).
		SetResult(respData).
		Put(API.Base(a.cl.host).P("id", dbId).String())); err != nil {
		return err
	}

	return nil
}

func (a *Api) Create(def *apidef.APIDefinition) (string, error) {
	if def.TagHeaders == nil {
		def.TagHeaders = []string{}
	}

	if def.ResponseProcessors == nil {
		def.ResponseProcessors = []apidef.ResponseProcessor{}
	}

	if len(def.VersionData.Versions) == 0 {
		return "", fmt.Errorf("api must have at least one version")
	}

	if def.PinnedPublicKeys == nil {
		def.PinnedPublicKeys = map[string]string{}
	}

	respData := &APICreateResp{}
	sendData := &DashApiDefinition{
		APIDefinition: def,
	}
	if err := api.HandleResponseAndError(a.cl.api.R().
		SetBody(sendData).
		SetResult(respData).
		Post(APIs.Base(a.cl.host).String())); err != nil {
		return "", err
	}

	return respData.Meta, nil
}

func (a *Api) Delete(dbId string) error {
	if err := api.HandleResponseAndError(a.cl.api.R().
		Delete(API.Base(a.cl.host).P("id", dbId).String())); err != nil {
		return err
	}

	return nil
}

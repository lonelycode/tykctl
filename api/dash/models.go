package dash

import (
	"github.com/TykTechnologies/tyk/apidef"
	"gopkg.in/mgo.v2"
)

type APIResp struct {
	Status  string
	Message string
	Meta    interface{}
}

type APICreateResp struct {
	Status  string
	Message string
	Meta    string
}

type APIDefinitions struct {
	Apis    []*DashApiDefinition `json:"apis"`
	Pages   int                  `json:"pages"`
	session *mgo.Session
}

func (ad *APIDefinitions) AsStandardDef() []*apidef.APIDefinition {
	newAd := make([]*apidef.APIDefinition, len(ad.Apis))
	for i, d := range ad.Apis {
		newAd[i] = d.APIDefinition
	}

	return newAd
}

type DashApiDefinition struct {
	*apidef.APIDefinition `bson:"api_definition,inline" json:"api_definition,inline"`
	IsSite                bool `bson:"is_site" json:"is_site"`
	SortBy                int  `bson:"sort_by" json:"sort_by"`
}

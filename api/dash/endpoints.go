package dash

import "github.com/TykTechnologies/tykctl/api/endpoints"

var (
	API          = endpoints.EP("/api/apis/{id}")
	APIs         = endpoints.EP("/api/apis")
	APISearch    = endpoints.EP("/api/apis")
	Policies     = endpoints.EP("/api/portal/policies/")
	PolicySearch = endpoints.EP("/api/portal/policies/")
	Policy       = endpoints.EP("/api/portal/policies/{id}")
	Token        = endpoints.EP("/api/apis/dummy-api/keys/{key-id}")
	CreateToken  = endpoints.EP("/api/keys")
)

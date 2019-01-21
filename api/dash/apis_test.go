package dash

import (
	"github.com/TykTechnologies/tyk/apidef"
	"github.com/TykTechnologies/tykctl/api/_test_util"
	"os"
	"testing"
)

var testHost = "http://localhost:8989"
var testToken = "foo"

func setup() *_test_util.DashServerMock {
	wkDir := os.Getenv("TYKCTRL_WKDIR")
	if wkDir == "" {
		panic("set TYKCTRL_WKDIR otherwise validation tests will fail")
	}
	s := &_test_util.DashServerMock{}
	s.Start(":8989")

	return s
}

func tearDown(s *_test_util.DashServerMock) {
	s.Stop()
}

func TestNew(t *testing.T) {
	x := New(testHost, testToken)
	if x == nil {
		t.Fatal("client should be created")
	}
}

func TestClient_Api(t *testing.T) {
	x := New(testHost, testToken)
	a := x.Api()
	if a == nil {
		t.Fatal("API func should return object")
	}

	if a.cl != x {
		t.Fatal("APi client should be same client as New()")
	}
}

func TestApi_List(t *testing.T) {
	s := setup()
	defer tearDown(s)

	x := New(testHost, testToken)
	list, err := x.Api().List()
	if err != nil {
		t.Fatalf("should have returned list: %v", err)
	}

	if len(list) == 0 {
		t.Fatalf("should have returned list > 0")
	}
}

func TestApi_Fetch(t *testing.T) {
	s := setup()
	defer tearDown(s)

	var cases = []struct {
		ApiID        string
		Code         int
		ExpectSVcErr bool
	}{
		{
			ApiID:        "foo",
			Code:         404,
			ExpectSVcErr: true,
		},
		{
			ApiID:        "",
			Code:         500,
			ExpectSVcErr: true,
		},
		{
			ApiID:        "581b5e91854a610001a2d3ff",
			Code:         200,
			ExpectSVcErr: false,
		},
	}

	x := New(testHost, testToken)
	for _, cs := range cases {
		apiDef, err := x.Api().Fetch(cs.ApiID)
		if cs.ExpectSVcErr {
			if err == nil {
				t.Fatal("expected service error, got nil")
			}
		} else {
			if err != nil {
				t.Fatal("expected no error, got: ", err)
			}

			if apiDef.Id.Hex() != cs.ApiID {
				t.Fatalf("API of decoded object ID must match, got: %v", apiDef.Id.Hex())
			}
		}
	}

}

func TestApi_Create(t *testing.T) {
	s := setup()
	defer tearDown(s)

	ad := &apidef.APIDefinition{}
	ad.Name = "foo"
	ad.Slug = "foo"
	ad.Proxy.ListenPath = "/"
	ad.Proxy.TargetURL = "http://example.com"

	defaultVersion := apidef.VersionInfo{
		Name: "Default",
	}
	defaultVersion.Paths.Ignored = []string{}
	defaultVersion.Paths.BlackList = []string{}
	defaultVersion.Paths.WhiteList = []string{}

	ad.VersionData.Versions = map[string]apidef.VersionInfo{
		"default": defaultVersion,
	}

	ad.AllowedIPs = []string{}
	ad.UpstreamCertificates = map[string]string{}
	ad.BlacklistedIPs = []string{}
	ad.ClientCertificates = []string{}
	ad.Tags = []string{}
	ad.ConfigData = map[string]interface{}{}
	ad.CustomMiddleware.Pre = []apidef.MiddlewareDefinition{}
	ad.CustomMiddleware.Post = []apidef.MiddlewareDefinition{}

	x := New(testHost, testToken)
	id, err := x.Api().Create(ad)
	if err != nil {
		t.Fatal(err)
	}

	if id != "5c43f0ffd1f3fd0001ff797e" {
		t.Fatalf("expected ID: %s, got %s", "5c43f0ffd1f3fd0001ff797e", id)
	}
}

func TestApi_Update(t *testing.T) {
	s := setup()
	defer tearDown(s)

	ad := &apidef.APIDefinition{}
	ad.Name = "foo"
	ad.Slug = "foo"
	ad.Proxy.ListenPath = "/"
	ad.Proxy.TargetURL = "http://example.com"

	x := New(testHost, testToken)
	err := x.Api().Update("581b5e91854a610001a2d3ff", ad)
	if err != nil {
		t.Fatal(err)
	}

	// Invalid
	ad.Proxy.TargetURL = "foo.bar:123"

	_, err2 := x.Api().Create(ad)
	if err2 == nil {
		t.Fatal("handler should ahve reported an error")
	}
}

func TestApi_Delete(t *testing.T) {
	s := setup()
	defer tearDown(s)

	x := New(testHost, testToken)
	err := x.Api().Delete("581b5e91854a610001a2d3ff")
	if err != nil {
		t.Fatal(err)
	}
}

func TestApi_Search(t *testing.T) {
	s := setup()
	defer tearDown(s)

	x := New(testHost, testToken)
	res, err := x.Api().Search("P1")
	if err != nil {
		t.Fatal(err)
	}

	if len(res) == 0 {
		t.Fatal("expected results longer than 0")
	}
}

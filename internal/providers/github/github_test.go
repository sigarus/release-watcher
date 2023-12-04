package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getReleaseByPath(t *testing.T) {
	type testCase struct {
		name        string
		path        string
		client      *http.Client
		wantRelease ReleaseInfo
		wantErr     string
	}

	getReleaseByPath_Case := func(
		tc testCase,
	) func(t *testing.T) {
		return func(t *testing.T) {
			t.Helper()
			t.Parallel()
			req := require.New(t)

			gp := GithubProvider{
				Path:   tc.path,
				client: tc.client,
			}

			release, err := gp.getRelease()
			fmt.Println(release)
			req.Equal(release, tc.wantRelease)
			if err != nil {
				req.EqualError(err, tc.wantErr)
			}
		}
	}

	testCases := []testCase{
		{
			name:   "Simple test",
			path:   "test/test1",
			client: getHTTPClient_Test(),
			wantRelease: ReleaseInfo{
				TagName: "v0.0.1",
			},
		},
		{
			name:        "Not 200 status test",
			path:        "test/test2",
			client:      getHTTPClient_Test(),
			wantRelease: ReleaseInfo{},
			wantErr:     errNo200.Error(),
		},
		{
			name:        "Unmarshal error test",
			path:        "test/test3",
			client:      getHTTPClient_Test(),
			wantRelease: ReleaseInfo{},
			wantErr:     "unexpected end of JSON input",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, getReleaseByPath_Case(tc))
	}

}

func getHTTPClient_Test() *http.Client {
	test1_struct := ReleaseInfo{
		TagName: "v0.0.1",
	}
	test1_json, _ := json.Marshal(test1_struct)
	test1_Path := fmt.Sprintf("%v/%v", "/repos/test/test1", lastReleaseReq)

	test2_Path := fmt.Sprintf("%v/%v", "/repos/test/test2", lastReleaseReq)
	test3_Path := fmt.Sprintf("%v/%v", "/repos/test/test3", lastReleaseReq)

	client := &http.Client{
		Transport: RoundTripFunc(func(req *http.Request) *http.Response {
			switch req.URL.Path {
			case test1_Path:
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(test1_json)),
				}
			case test2_Path:
				return &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       nil,
				}
			case test3_Path:
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       nil,
				}
			}

			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader("KEK")),
			}
		})}

	return client
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

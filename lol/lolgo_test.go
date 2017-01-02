package lol

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	apiKey := randString(10)
	region := Region(r1.Intn(int(Pbe) + 1))
	client := NewClient(apiKey, region, nil)

	if client.APIKey != apiKey {
		t.Errorf("expected %v, got %v", apiKey, client.APIKey)
	}
	if client.region != region {
		t.Errorf("expected %v, got %v", region, client.region)
	}
}

func TestClient_Region(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	region := Region(r1.Intn(int(Pbe) + 1))
	client := NewClient("", region, nil)

	if region != client.Region() {
		t.Errorf("expected %v, got %v", region, client.Region())
	}
}

func TestRiotAPIError_Error(t *testing.T) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	err := RiotAPIError{
		Status: riotAPIErrorStatus{
			StatusCode: http.StatusTooManyRequests,
			Message:    "Too many requests",
		},
		RetryAfter:      r1.Int(),
		XRateLimitCount: r1.Int(),
		XRateLimitType:  randString(10),
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, strconv.Itoa(err.RetryAfter)) {
		t.Error("error does not contain Retry-After header")
	}
	if !strings.Contains(errMsg, strconv.Itoa(err.XRateLimitCount)) {
		t.Error("error does not contain X-Rate-Limit-Count header")
	}
	if !strings.Contains(errMsg, err.XRateLimitType) {
		t.Error("error does not contain X-Rate-Limit-Type header")
	}
	if !strings.Contains(errMsg, strconv.Itoa(err.Status.StatusCode)) {
		t.Error("error does not contain Status Code value")
	}
	if !strings.Contains(errMsg, err.Status.Message) {
		t.Error("error does not contain Status Message")
	}
}

func mockClient(region Region) (*http.Client, *http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	httpClient := &http.Client{Transport: transport}

	return httpClient, mux, server, NewClient("", region, httpClient)
}

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

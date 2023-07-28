package uni

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/unimtx/uni-go-sdk/meta"
)

const defaultEndpoint = "https://api.unimtx.com"
const defaultSigningAlgorithm = "hmac-sha256"
const requestIdHeaderKey = "x-uni-request-id"
const userAgent = "uni-go-sdk" + "/" + meta.VERSION

type UniClient struct {
	AccessKeyId			string
	AccessKeySecret		string
	Endpoint			string
	SigningAlgorithm	string
	Messages            *MessageService
	Otp					*OtpService
}

type UniResponse struct {
	Raw			*http.Response
	Status		int
	Code		string
	Message		string
	Data		map[string]interface{}
	RequestId	string
	Valid		bool
}

func NewClient(params ...string) *UniClient {
	accessKeyId := os.Getenv("UNIMTX_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("UNIMTX_ACCESS_KEY_SECRET")
	endpoint := os.Getenv("UNIMTX_ENDPOINT")

	if len(params) > 0 {
		accessKeyId = params[0]
	}

	if len(params) > 1 {
		accessKeySecret = params[1]
	}

	if len(endpoint) == 0 {
		endpoint = defaultEndpoint
	}

	client := &UniClient{
		AccessKeyId: accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint: endpoint,
		SigningAlgorithm: defaultSigningAlgorithm,
	}
	client.Messages = &MessageService{client: client}
	client.Otp = &OtpService{client: client}

	return client
}

func NewResponse(res *http.Response) (*UniResponse, error) {
	var data map[string]interface{}
	var code string
	var message string

	status := res.StatusCode
	requestId := res.Header.Get(requestIdHeaderKey)
	rawBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if rawBody != nil {
		body := make(map[string]interface{})
		err := json.Unmarshal([]byte(rawBody), &body)

		if err != nil {
			return nil, err
		}

		code = body["code"].(string)
		message = body["message"].(string)

		if code != "0" {
			return nil, errors.New(fmt.Sprintf("[%s] %s, RequestId: %s", code, message, requestId))
		}
		data = body["data"].(map[string]interface{})
	}

	return &UniResponse{
		Raw: res,
		Status: status,
		Code: code,
		Message: message,
		Data: data,
		RequestId: requestId,
	}, nil
}

func (c *UniClient) GenerateRandomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}

func (c *UniClient) Sign(query url.Values) url.Values  {
	if c.AccessKeySecret != "" {
		query.Add("algorithm", c.SigningAlgorithm)
		query.Add("timestamp", strconv.FormatInt(time.Now().Unix(), 10))
		query.Add("nonce", c.GenerateRandomString(8))

		message := query.Encode()
		mac := hmac.New(sha256.New, []byte(c.AccessKeySecret))
		mac.Write([]byte(message))
		query.Add("signature", hex.EncodeToString(mac.Sum(nil)))
	}

	return query
}

func (c *UniClient) Request(action string, data map[string]interface{}) (response *UniResponse, err error) {
	u := c.Endpoint
	query := url.Values{}
	query.Add("action", action)
	query.Add("accessKeyId", c.AccessKeyId)
	query = c.Sign(query)
	querystr := query.Encode()
	jsonbytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(jsonbytes)
	client := &http.Client {
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("POST", u + "/?" + querystr, reader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return NewResponse(res)
}

func (c *UniClient) SetEndpoint(endpoint string) {
	c.Endpoint = endpoint
}

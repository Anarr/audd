package audd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	auddURL               = "https://api.audd.io/"
	statusErr             = "error"
	statusSuccess         = "success"
	errServiceUnavailable = "service unavailable"
	//PlatformSpoify display spotify platform
	PlatformSpoify = "spotify"
	//PlatformAppleMusic display apple music platform
	PlatformAppleMusic = "apple_music"
	//PlatformDeezer display deezer music platform
	PlatformDeezer = "deezer"
)

//AudioReader containst audio client methods
type AudioReader interface {
	SetToken(token string) *AudioClient
}

//AudioClient contains audio client struct
type AudioClient struct {
	token    string
	platform string
}

//SetToken set token to audio client
func (ac *AudioClient) SetToken(token string) *AudioClient {
	ac.token = token
	return ac
}

//SetPlatform set platform to audio client
func (ac *AudioClient) SetPlatform(platform ...string) *AudioClient {
	ac.platform = strings.Join(platform, ",")
	return ac
}

//Response display cutsom response type
type Response struct {
	Title string `json:"title"`
}

type auddResponse struct {
	Status string
	Result *artist
	Error  *auddError
}

type artist struct {
	Artist      string `json:"artist"`
	Title       string `json:"title"`
	Album       string `json:"album"`
	ReleaseDate string `json:"release_date"`
	Label       string `json:"label"`
	TimeCode    string `json:"time_code"`
	SongLink    string `json:"song_link"`
}

type auddError struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

//Parse start audio detection process
func (ac *AudioClient) Parse(source string) (*Response, error) {
	data := url.Values{
		"url":       {source},
		"return":    {ac.platform},
		"api_token": {ac.token},
	}

	res, err := http.PostForm(auddURL, data)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if len(body) > 0 {
		var auddRes auddResponse

		err := auddRes.unMarshal(body)
		if err != nil {
			return nil, err
		}

		switch auddRes.Status {
		case statusSuccess:
			if auddRes.Result != nil {
				return &Response{Title: auddRes.Result.Title}, nil
			}
		case statusErr:
			return nil, errors.New(auddRes.Error.ErrorMessage)
		}
	}

	return nil, errors.New(errServiceUnavailable)
}

func (audd *auddResponse) unMarshal(data []byte) error {
	err := json.Unmarshal(data, audd)

	if err != nil {
		return err
	}

	return nil
}

//NewClient create new audio detection client
func NewClient(reader AudioReader) *AudioClient {
	return &AudioClient{}
}

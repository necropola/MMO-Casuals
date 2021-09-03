package gw2api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	API_BASE       = "https://api.guildwars2.com"
	IMG_BASE       = "https://render.guildwars2.com"
	CLIENT_TIMEOUT = 15 // Overall request timeout [seconds]
	DIAL_TIMEOUT   = 5  // Timeout [seconds] for establishing the TCP connection
	KEEP_ALIVE     = 10 // How long [seconds] to keep the TCP connection alive (for further requests)
	TLS_TIMEOUT    = 3  // Timeout [seconds] for establishing TLS on the TCP connection
	LOG_INDENT     = 4  // Indent Width (spaces) for (multi line) diagnostic output
)

type GW2API struct {
	client *http.Client
	key    string
	logger *log.Logger
}

type APIOption func(*GW2API)

func WithAuth(key string) APIOption {
	return func(api *GW2API) {
		api.key = key
	}
}

func WithLogger(logger *log.Logger) APIOption {
	return func(api *GW2API) {
		api.logger = logger
	}
}

func New(opts ...APIOption) (api *GW2API) {
	api = &GW2API{
		client: &http.Client{
			Timeout: time.Duration(CLIENT_TIMEOUT * time.Second),
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(DIAL_TIMEOUT * time.Second),
					KeepAlive: time.Duration(KEEP_ALIVE * time.Second),
				}).DialContext,
				TLSHandshakeTimeout: time.Duration(TLS_TIMEOUT * time.Second),
			},
		},
	}
	for _, opt := range opts {
		opt(api)
	}
	return api
}

func trim_indent_multi(s string, i int) string {
	return strings.ReplaceAll("\n"+strings.Trim(s, "\n"), "\n", "\n"+strings.Repeat(" ", i))
}

func log_msg(logger *log.Logger, message string, details string) {
	if logger == nil {
		return
	}
	if details != "" {
		logger.Printf("%s: %s\n", message, trim_indent_multi(details, LOG_INDENT))
	} else {
		logger.Println(message)
	}
}

func (api *GW2API) fetch(query string, result interface{}) (err error) {
	req, err := http.NewRequest("GET", API_BASE+query, nil)
	if err != nil {
		return err
	}
	if api.key != "" {
		req.Header.Set("Authorization", "Bearer "+api.key)
	}
	log_msg(api.logger, "HTTP Request", "Method: "+req.Method+"\n"+
		"Proto : "+req.Proto+"\n"+
		"Scheme: "+req.URL.Scheme+"\n"+
		"Host  : "+req.URL.Host+"\n"+
		"Path  : "+req.URL.Path+"\n"+
		"Params: "+req.URL.RawQuery+"\n"+
		"Auth  : "+req.Header.Get("Authorization"))
	resp, err := api.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	// Using ioUtil.ReadAll() + json.Unmarshal() instead of json.NewDecoder().Decode()
	// to make logging of the response easier
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log_msg(api.logger, "HTTP Response", "Body:\n"+
		string(data))
	if err = json.Unmarshal(data, &result); err != nil {
		return err
	}
	return nil
}

func (api *GW2API) Anything(query string) (result interface{}, err error) {
	err = api.fetch(query, &result)
	return
}

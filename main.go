/*
Package relay implments a simple client to report errors using the Mailgun mailing service
*/
package relay

import (
	"encoding/json"
	"errors"
	"net/http"
	//"net/url"
	"os"
	//"time"
)

var ErrBadConfig = errors.New("The config does not contain the necessary information")

// Relay is a client to send error messages with
type Relay struct {
	c    *http.Client
	to   string
	from string
	key  string
}

// Config contains the information used to initialize a Relay
type Config struct {
	To   string `json:"to"`
	From string `json:"from"`
	Key  string `json:"api_key"`
}

// New is used to generate a new Relay. If called with argument nil, it
// reads from config.json
func New(c *Config) (*Relay, error) {
	// make a new relay
	r := &Relay{
		c: &http.Client{},
	}

	// if they gave us a config, use it
	if c != nil {
		r.to = c.To
		r.from = c.From
		r.key = c.Key
	} else { // otherwise read config.json
		infile, err := os.Open("config.json")
		defer infile.Close()
		if err != nil {
			return nil, ErrBadConfig
		}

		dec := json.NewDecoder(infile)

		config := new(Config)
		err = dec.Decode(config)
		if err != nil {
			return nil, ErrBadConfig
		}

		r.to = config.To
		r.from = config.From
		r.key = config.Key
	}

	// ensure all necessary fields are set
	if r.to == "" || r.from == "" || r.key == "" {
		return nil, ErrBadConfig
	} else {
		return r, nil
	}
}

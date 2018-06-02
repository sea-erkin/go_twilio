package go_twilio

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	print = fmt.Println
)

func New(accountId, accountKey string, verbose bool) (retval *Twilio, err error) {

	if accountId == "" || accountKey == "" {
		return nil, errors.New("AccountId or Key cannot be empty strings. Check yo strings")
	}

	retval = &Twilio{
		AccountId:  accountId,
		AccountKey: accountKey,
		Verbose:    verbose,
		Endpoints: Endpoints{
			Messages: "https://api.twilio.com/2010-04-01/Accounts/{accountId}/Messages.json",
		},
	}
	return retval, nil
}

func (this *Twilio) SendMessage(to, from, smsBody string) error {
	url := strings.Replace(this.Endpoints.Messages, "{accountId}", this.AccountId, 1)

	// can validate that these are good phone numbers

	var builder strings.Builder
	builder.WriteString("To=")
	builder.WriteString(to)
	builder.WriteString("&From=")
	builder.WriteString(from)
	builder.WriteString("&Body=")
	builder.WriteString(smsBody)

	payload := strings.NewReader(builder.String())

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(this.AccountId+":"+this.AccountKey))
	print(basicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", basicAuth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if this.Verbose {
		fmt.Println(res)
		fmt.Println(string(body))
	}

	if res.StatusCode != http.StatusCreated {
		return errors.New("Non 201 response received")
	}

	return nil

}

type Twilio struct {
	AccountId  string
	AccountKey string
	Verbose    bool
	Endpoints
}

type Endpoints struct {
	Messages string
}

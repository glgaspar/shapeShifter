package main

import (
	"encoding/json"
	"errors"
	"strings"
)

type Request struct {
	Data          json.RawMessage         `json:"data"`
	ResponseLang  []string                `json:"responseLang"`
	ResponseData  []string                `json:"-"`
	ParseSelector map[string]ParseHandler `json:"-"`
}

func (r *Request) Parse(lang int) error {
	if handler, ok := r.ParseSelector[strings.ToLower(r.ResponseLang[lang])]; ok {
		if err := handler(r); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("language not available to shape shift into")
	}
}

func (r *Request) SetUpSelectors() {
	r.ParseSelector[Mssql] = MssqlParser
	r.ParseSelector[Go] = GoParser
}

func NewRequest() *Request {
	r := &Request{}
	r.SetUpSelectors()
	return r
}

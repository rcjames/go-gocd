package gocd

import (
	"encoding/json"
)

type Links struct {
	Self string
	Doc  string
	Find string
}

func (l *Links) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	selfMap, ok := m["self"]
	if ok {
		selfData := selfMap.(map[string]interface{})
		l.Self = selfData["href"].(string)
	}

	docMap, ok := m["doc"]
	if ok {
		docData := docMap.(map[string]interface{})
		l.Doc = docData["href"].(string)
	}

	findMap, ok := m["find"]
	if ok {
		findData := findMap.(map[string]interface{})
		l.Find = findData["href"].(string)
	}

	return nil
}

type DeleteMessage struct {
	Message string `json:"message"`
}

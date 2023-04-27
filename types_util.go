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

	selfMap := m["self"]
	selfData := selfMap.(map[string]interface{})
	l.Self = selfData["href"].(string)

	docMap := m["doc"]
	docData := docMap.(map[string]interface{})
	l.Doc = docData["href"].(string)

	findMap := m["find"]
	findData := findMap.(map[string]interface{})
	l.Find = findData["href"].(string)

	return nil
}

type DeleteMessage struct {
	Message string `json:"message"`
}

package gocd

import (
	"encoding/json"
)

// A Links object contains some response metadata which is returned from
// requests to the GoCD API.
type Links struct {
	Self string `json:"self,omitempty"`
	Doc  string `json:"doc,omitempty"`
	Find string `json:"find,omitempty"`
}

// The Links UnmarshalJSON function is used to unmarshall the nested structure
// returned by the API into a more friendly format.
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

// A DeleteMessage object is used for handling the message reponse recieved
// when a delete request is made.
type DeleteMessage struct {
	Message string `json:"message"`
}

// A ConfigurationProperty maps to a [configuration property] oject.
//
// [configuration property]: https://api.gocd.org/current/#the-configuration-property-object
type ConfigurationProperty struct {
	Key            string `json:"key"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
	IsSecure       bool   `json:"is_secure,omitempty"`
}

// A Pagination is used to handle the pagination section of the response to
// requests for history, such as [get material modifications] and [get job
// history]
//
// [get material modifications]: https://api.gocd.org/current/#get-material-modifications
// [get job history]: https://api.gocd.org/current/#get-job-history
type Pagination struct {
	Offset   int `json:"offset,omitempty"`
	Total    int `json:"total"`
	PageSize int `json:"page_size,omitempty"`
}

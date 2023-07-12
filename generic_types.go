package gocd

// TODO - Validation for Value or EncryptedValue
type ConfigurationProperty struct {
	Key            string `json:"key"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
	IsSecure       bool   `json:"is_secure,omitempty"`
}

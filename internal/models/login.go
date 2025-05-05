package models

type Login struct {
	ID        string `json:"id,omitempty"`
	ID2 string `json:"id2,omitempty"`
	USERNAME  string `json:"username,omitempty"`
	PASSWORD     string `json:"password,omitempty"`
	ACCOUNT_STATUS     string `json:"account_status,omitempty"`
	ACCOUNT_GROUP   string `json:"account_group,omitempty"`
	EMAIL   string `json:"email,omitempty"`
}

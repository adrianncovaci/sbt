package models

type CustomItem struct {
	Id          uint   `json:"id"`
	System      string `json:"system"`
	ItemType    string `json:"type"`
	Cmd         string `json:"cmd"`
	Regex       string `json:"regex"`
	Expect      string `json:"expect"`
	Description string `json:"description"`
	Info        string `json:"info"`
	Reference   string `json:"reference"`
	See_also    string `json:"see_also"`
	File        string `json:"file"`
	Owner       string `json:"owner"`
	Mask        string `json:"mask"`
	Group       string `json:"group"`
	AuditID     uint
}

type Audit struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	CustomItems []CustomItem
}

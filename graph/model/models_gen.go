// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewQuery struct {
	Text string `json:"text"`
	ID   string `json:"id"`
}

type Vehicle struct {
	ID        string `json:"id"`
	MakeID    string `json:"makeId"`
	MakeName  string `json:"makeName"`
	ModelName string `json:"modelName"`
}

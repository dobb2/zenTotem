package entity

type User struct {
	Id   int    `json:"id,omitempty"`
	Age  int    `json:"age,omitempty"`
	Name string `json:"name,omitempty"`
}

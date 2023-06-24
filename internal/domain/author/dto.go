package author

type CreateAuthorDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateAuthorDto struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

package data

type NewDomain struct {
	ID      int    `json:"id"`
	OwnerId int    `json:"owner_id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

type Domain struct {
	OwnerId int    `json:"owner_id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

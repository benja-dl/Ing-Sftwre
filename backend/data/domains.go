package data

type NewDomain struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Domain struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

type DomainID struct {
	ID int `json:"domain_id"`
}

type DomainName struct {
	ID   int    `json:"domain_id"`
	Name string `json:"name"`
}

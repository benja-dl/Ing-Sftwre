package data

// NewAccount Se ocupa el json despues de cada campo, por que permite que al codificar a JSON se asigne un campo con ese nombre o que al decodificar se asigne
// el valor de un campo de la solicitud json al parametro coincidente/*
type NewAccount struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

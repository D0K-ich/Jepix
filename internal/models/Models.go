package Models

// TODO сделать конфиг для админа и сервера. Переработать конфиг бд
type //Users depends
(
	Customer struct {
		Name           string `json:"name" `
		Surname        string `json:"surname"`
		Email          string `json:"email"`
		Dateofregister int    `json:"dateofregister"`
		Balancemoney   int    `json:"balancemoney"`
		Balancebonus   int    `json:"balancebonus"`
		Promocodes     int    `json:"promocodes"`
	}
	Promocode struct {
		Datecreated int `json:"Dateofregister"`
		Datedead    int `json:"Datedead"`
		Percent     int `json:"Percent"`
		UUID        int `json:"uuid"`
	}
)

type ConfigDB struct {
	User       string
	Pass       string
	Collection string
}

type TU struct {
	Name     string `json:"name" `
	Password string `json:"password"`
	Email    string `json:"email"`
}

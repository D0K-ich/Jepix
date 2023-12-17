package Models

// TODO сделать конфиг для админа и сервера. Переработать конфиг бд
type //Users depends
(
	Customer struct {
		Name           string    `json:"name" `
		Surname        string    `json:"surname"`
		Email          string    `json:"email"`
		Phone          string    `json:"phone"`
		Password       string    `json:"password"`
		Dateofregister int64     `json:"dateofregister"`
		Balancemoney   int       `json:"balancemoney"`
		Balancebonus   int       `json:"balancebonus"`
		Promocodes     Promocode `json:"promocodes"`
	}
	Promocode struct {
		Datecreated int64  `json:"Dateofregister"`
		Datedead    int64  `json:"Datedead"`
		Percent     int    `json:"Percent"`
		UUID        string `json:"uuid"`
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

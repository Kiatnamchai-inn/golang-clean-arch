package entities

type UsersUsecase interface {
	GetUserAndOrderListById(id string) (*GetUserAndOrderListByIdRes, error)
	UserLogin(email string, password string) (*UserLoginRes, error)
	UserLogin2(email string, password string) (*UserLoginRes, error)
}

type UsersRepository interface {
	GetUserAndOrderListById(id string) (*GetUserAndOrderListByIdRes, error)
	UserLogin(email string, password string) (*UserLoginRes, error)
}

type GetUserAndOrderListByIdRes struct {
	ID     int32        `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	Name   string       `gorm:"column:name;type:character varying(255);not null" json:"name"`
	Email  string       `gorm:"column:email;type:character varying(255);not null" json:"email"`
	Orders []UserOrders `json:"orders" db:"orders" gorm:"foreignKey:UserID"`
}

type UserLoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserLoginRes struct {
	ID    int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	Name  string `gorm:"column:name;type:character varying(255);not null" json:"name"`
	Email string `gorm:"column:email;type:character varying(255);not null" json:"email"`
}

type UserOrders struct {
	ID           int32   `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	UserID       int32   `gorm:"column:user_id;type:integer;not null" json:"-"` // don't show user_id in json
	ProductName  string  `gorm:"product_name" json:"product_name"`              // from products table
	ProductPrice float64 `gorm:"product_price" json:"product_price"`            // from products table
	Quantity     int32   `gorm:"column:quantity;type:integer;not null" json:"quantity"`
}

func (*UserOrders) TableName() string {
	return "orders"
}

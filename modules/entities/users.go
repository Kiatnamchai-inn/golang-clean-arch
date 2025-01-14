package entities

type UsersUsecase interface {
	GetUserAndOrderListById(id string) (*GetUserAndOrderListByIdRes, error)
	// GetUserProductPaymentById(id string) (*UserProductPaymentByIdRes, error)
}

type UsersRepository interface {
	GetUserAndOrderListById(id string) (*GetUserAndOrderListByIdRes, error)
}

type GetUserAndOrderListByIdRes struct {
	ID     int32        `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
	Name   string       `gorm:"column:name;type:character varying(255);not null" json:"name"`
	Email  string       `gorm:"column:email;type:character varying(255);not null" json:"email"`
	Orders []UserOrders `json:"orders" db:"orders" gorm:"foreignKey:UserID"`
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

// type UserProductPaymentByIdRes struct {
// 	ID       uint64       `json:"id" db:"id"`
// 	Name     string       `json:"name" db:"name"`
// 	Email    string       `json:"email" db:"email"`
// 	Payments []UserOrders `json:"orders" db:"orders" gorm:"foreignKey:UserID"`
// }

// type UserProductPayments struct {
// 	ID        int32 `gorm:"column:id;type:integer;primaryKey;autoIncrement:true" json:"id"`
// 	UserID    int32 `gorm:"column:user_id;type:integer;not null" json:"user_id"`
// 	ProductID int32 `gorm:"column:product_id;type:integer;not null" json:"product_id"`

// 	Amount float64 `gorm:"column:amount;type:numeric(10,2);not null" json:"amount"`
// }

package models

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

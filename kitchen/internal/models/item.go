package models

type Item struct {
	ID        string  `json:"id" bson:"_id,omitempty"`
	Name      string  `json:"name" bson:"name"`
	Price     float32 `json:"price" bson:"price"`
	Weight    float32 `json:"weight" bson:"weight"`
	CreatedAt string  `json:"created_at" bson:"created_at"`
	UpdatedAt string  `json:"updated_at" bson:"updated_at"`
	DeletedAt string  `json:"deleted_at" bson:"deleted_at"`
	IsActive  bool    `json:"is_active" bson:"is_active"`
}

package infrastructure

import (
	"github.com/google/uuid"
)

// PlayerModel es el modelo de persistencia GORM mapeado a la tabla players.
type PlayerModel struct {
	PlayerID   uuid.UUID `gorm:"column:player_id;type:uuid;primaryKey"`
	Email      string    `gorm:"column:email;type:varchar(255);uniqueIndex;not null"`
	Password   string    `gorm:"column:password;type:varchar(255);not null"`
	Nickname   string    `gorm:"column:nickname;type:varchar(100);not null"`
	Bio        *string   `gorm:"column:bio;type:varchar(500)"`
	Role       string    `gorm:"column:role;type:varchar(50);not null;default:'player'"`
	Exp        int       `gorm:"column:exp;not null;default:0"`
	Tobycoin   int       `gorm:"column:tobycoin;not null;default:0"`
	AvatarPath string    `gorm:"column:avatar_path;type:varchar(255);not null;default:'default_avatar.png'"`
}

// TableName retorna el nombre de la tabla en PostgreSQL para PlayerModel.
func (PlayerModel) TableName() string {
	return "players"
}

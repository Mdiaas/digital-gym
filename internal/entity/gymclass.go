package entity

import (
	"errors"
	"time"

	"github.com/mdiaas/goapi/pkg/entity"
)

var (
	ErrIDIsRequired   = errors.New("gym class id is required")
	ErrIDIsInvalid    = errors.New("gym class id is invalid")
	ErrNameIsRequired = errors.New("gym class name is required")
	ErrLinkIsRequired = errors.New("gym class link is required")
)

type GymClass struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
}

func NewGymClass(name, link string) (*GymClass, error) {
	gymClass := &GymClass{
		ID:        entity.NewID(),
		Name:      name,
		Link:      link,
		CreatedAt: time.Now(),
	}
	err := gymClass.Validate()
	if err != nil {
		return nil, err
	}
	return gymClass, nil
}

func (g *GymClass) Validate() error {
	if g.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(g.ID.String()); err != nil {
		return ErrIDIsInvalid
	}
	if g.Name == "" {
		return ErrNameIsRequired
	}
	if g.Link == "" {
		return ErrLinkIsRequired
	}
	return nil
}

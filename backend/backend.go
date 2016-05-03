package backend

import (
	"errors"
	"github.com/gonum/blas"
)

var (
	ErrNotImpl = errors.New("Not implemented")
)

type Rec struct {
}

type FilterReq struct {
	Users       string
	Categories  []string
	NumFeatures uint
}

func Format(path string) error {
	return ErrNotImpl
}

func Open(path string) (*Rec, error) {
	return ErrNotImpl
}

func (rec *Rec) Close() error {
	return ErrNotImpl
}

func (rec *Rec) AddUser(email string) error {
	return ErrNotImpl
}

func (rec *Rec) RmUser(email string) error {
	return ErrNotImpl
}

func (rec *Rec) AddItem(item string) error {
	return ErrNotImpl
}

func (rec *Rec) CatItem(item, category string) error {
	return ErrNotImpl
}

func (rec *Rec) AddReview(email, item, rating float64) error {
	return ErrNotImpl
}

func (rec *Rec) ChReview(email, item, rating float64) error {
	return ErrNotImpl
}

func (rec *Rec) GuessRating(email, item string) (float64, error) {
	return ErrNotImpl
}

func (rec *Rec) Filter(FilterReq req) error {
	return ErrNotImpl
}

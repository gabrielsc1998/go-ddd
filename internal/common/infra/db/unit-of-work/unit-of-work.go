package unit_of_work

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type RepositoryFactory func(db *gorm.DB) interface{}

type UowInterface interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	Commit() error
	Rollback() error
	UnRegister(name string)
	GetCtx() context.Context
}

type Uow struct {
	Db           *gorm.DB
	Tx           *gorm.DB
	Repositories map[string]RepositoryFactory
}

func NewUow(ctx context.Context, db *gorm.DB) *Uow {
	return &Uow{
		Db:           db,
		Tx:           nil,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) begin(ctx context.Context) {
	tx := u.Db.Begin()
	u.Tx = tx
	u.Tx.Statement.Context = ctx
}

func (u *Uow) reset() {
	u.Tx = nil
}

func (u *Uow) Register(name string, repositoryFactory RepositoryFactory) {
	u.Repositories[name] = repositoryFactory
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil || u.Tx.Statement.Context == nil {
		u.begin(ctx)
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(Uow *Uow) error) error {
	isANewCtx := u.Tx.Statement.Context != ctx
	if u.Tx.Statement.Context != nil && isANewCtx {
		return fmt.Errorf("transaction already started")
	}
	if isANewCtx {
		u.begin(ctx)
	}
	err := fn(u)
	if err != nil {
		errRb := u.Rollback()
		if errRb != nil {
			return fmt.Errorf("original error: %s, rollback error: %s", err.Error(), errRb.Error())
		}
		return err
	}
	return u.Commit()
}

func (u *Uow) Rollback() error {
	if u.Tx.Statement.Context == nil {
		return errors.New("no transaction to rollback")
	}
	u.Tx.Rollback()
	u.reset()
	return nil
}

func (u *Uow) Commit() error {
	u.Tx.Commit()
	u.reset()
	return nil
}

func (u *Uow) GetCtx() context.Context {
	return u.Tx.Statement.Context
}

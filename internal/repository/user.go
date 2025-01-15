package repository

import (
	"context"
	"database/sql"
	"ecommerce/internal/model"
	"ecommerce/pkg/infrastructure/query"
	"fmt"
)

type userImpl struct {
	db  *sql.DB
	ctx context.Context
}

func NewUserRepository(db *sql.DB, ctx context.Context) UserRepository {
	return &userImpl{db: db, ctx: ctx}
}
func (r *userImpl) RegisterUser(user *model.User) error {
	stmt, err := r.db.PrepareContext(r.ctx, query.UserLogin)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *userImpl) CountUser(email *string) (int, error) {
	stmt, err := r.db.PrepareContext(r.ctx, query.UserCount)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &email)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	} else {
		return 0, err
	}
	return count, nil
}
func (r *userImpl) GetUserByEmail(email *string) (*model.User, error) {
	user := new(model.User)
	stmt, err := r.db.PrepareContext(r.ctx, query.UserGetByEmail)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user email %s not found", *email)
	}
	return user, nil
}
func (r *userImpl) GetUserById(userId *string) (*model.User, error) {
	user := new(model.User)
	stmt, err := r.db.PrepareContext(r.ctx, query.UserGetById)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(r.ctx, &userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user %s not found", *userId)
	}
	return user, nil
}

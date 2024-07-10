package main

import (
	"context"
	"database/sql"

	"casethree/helper"
	"casethree/model"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, entity model.MainEntity) interface{}
	FindTheMostSpentCountry(ctx context.Context, tx *sql.Tx) []model.MainEntity
	FindTheMostUserCreditCardType(ctx context.Context, tx *sql.Tx) interface{}
}

type RepositoryImpl struct{}

// FindTheMostUserCreditCardType implements Repository.

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (repository *RepositoryImpl) FindTheMostUserCreditCardType(ctx context.Context, tx *sql.Tx) interface{} {
	SQL := "select credit_card_type from users group by(credit_card_type) order by count(credit_card_type) desc limit 1;"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var entities []string
	for rows.Next() {
		entity := ""
		err := rows.Scan(&entity)
		helper.PanicIfError(err)
		entities = append(entities, entity)
	}
	return entities
}

func (repository *RepositoryImpl) Save(ctx context.Context, tx *sql.Tx, entity model.MainEntity) interface{} {
	SQL := "insert into users (country, credit_card_type, credit_card, first_name, last_name) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, entity.Country, entity.CreditCardType, entity.CreditCard, entity.FirstName, entity.LastName)
	helper.PanicIfError(err)

	_, err = result.LastInsertId()
	helper.PanicIfError(err)

	return nil
}

func (repository *RepositoryImpl) FindTheMostSpentCountry(ctx context.Context, tx *sql.Tx) []model.MainEntity {
	SQL := "select id, country, credit_card_type, credit_card, first_name, last_name from users where country = (select u.country from users as u join transactions as t on u.id = t.id_user group by(u.country) order by sum(t.total_buy) desc limit 1);"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var entities []model.MainEntity
	for rows.Next() {
		entity := model.MainEntity{}
		err := rows.Scan(&entity.Id, &entity.Country, &entity.CreditCardType, &entity.CreditCard, &entity.FirstName, &entity.LastName)
		helper.PanicIfError(err)
		entities = append(entities, entity)
	}
	return entities
}

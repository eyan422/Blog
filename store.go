package main

import (
	"database/sql"
	"github.com/eyan422/Blog/CommonStruct"
	"log"
)

type Store interface {
	CreateArticle(article *CommonStruct.Article) (int64, error)
	GetArticles() (articles []*CommonStruct.Article, err error)
	GetArticle(id uint64) (articles *CommonStruct.Article, err error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateArticle(article *CommonStruct.Article) (int64, error) {
	query := "INSERT INTO article(title, content, author) VALUES (?, ?, ?)"

	stmt, err := store.db.Prepare(query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(article.Title, article.Content, article.Author)
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0, err
	}
	log.Printf("%d record created ", rows)

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, err
}

func (store *dbStore) GetArticles() (articles []*CommonStruct.Article, err error) {
	rows, err := store.db.Query("SELECT id, title, content, author from article order by id")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		article := &CommonStruct.Article{}
		if err := rows.Scan(&article.Id, &article.Title, &article.Content, &article.Author); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (store *dbStore) GetArticle(id uint64) (article *CommonStruct.Article, err error) {
	res, err := store.db.Query("SELECT id, title, content, author from article where id = ?", id)
	defer res.Close()
	if err != nil {
		return nil, err
	}

	if res.Next() {
		article = &CommonStruct.Article{}
		if err := res.Scan(&article.Id, &article.Title, &article.Content, &article.Author); err != nil {
			return nil, err
		}
	}

	return article, nil
}

var store Store

func InitStore(s Store) {
	store = s
}

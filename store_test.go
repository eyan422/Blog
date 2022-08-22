package main

import (
	"database/sql"
	"fmt"
	"github.com/eyan422/Blog/CommonStruct"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct {
	suite.Suite
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "root:12345678@tcp(localhost:3306)/assignment"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Exec("DELETE FROM article")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	_, err := s.db.Exec("DELETE FROM article")
	if err != nil {
		s.T().Fatal(err)
	}

	err = s.db.Close()
	if err != nil {
		return
	}
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateArticle() {
	_, err := s.store.CreateArticle(&CommonStruct.Article{
		Title:   "war and peace",
		Content: "novel",
		Author:  "frankie",
	})
	if err != nil {
		return
	}

	res, err := s.db.Query(`SELECT COUNT(*) FROM article WHERE title='war and peace' AND content='novel' AND author = 'frankie'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got [%d]", count)
	}
}

func (s *StoreSuite) TestGetArticles() {

	_, err := s.db.Exec("INSERT INTO article(title, content, author) VALUES ('blue', 'novel', 'feng')")
	if err != nil {
		s.T().Fatal(err)
	}

	_, err = s.db.Exec("INSERT INTO article(title, content, author) VALUES ('pink', 'novel', 'Frankie')")
	if err != nil {
		s.T().Fatal(err)
	}

	articleRecords, err := s.store.GetArticles()
	if err != nil {
		s.T().Fatal(err)
	}

	nArticles := len(articleRecords)
	if nArticles != 2 {
		s.T().Errorf("incorrect count, wanted 2, got [%d]", nArticles)
	}

	expectedArticle1 := CommonStruct.Article{
		Id:      1,
		Title:   "blue",
		Content: "novel",
		Author:  "feng",
	}

	if articleRecords[0].Title != expectedArticle1.Title || articleRecords[0].Author != expectedArticle1.Author || articleRecords[0].Content != expectedArticle1.Content {
		s.T().Errorf("incorrect details, expected [%v], got [%v]", expectedArticle1, articleRecords[0])
	}

	expectedArticle2 := CommonStruct.Article{
		Id:      1,
		Title:   "pink",
		Content: "novel",
		Author:  "Frankie",
	}

	if articleRecords[1].Title != expectedArticle2.Title || articleRecords[1].Author != expectedArticle2.Author || articleRecords[1].Content != expectedArticle2.Content {
		s.T().Errorf("incorrect details, expected [%v], got [%v]", expectedArticle2, articleRecords[1])
	}
}

func (s *StoreSuite) TestGetArticle() {

	res, err := s.db.Exec("INSERT INTO article(title, content, author) VALUES ('black', 'Sci-Fi', 'Rockey')")
	if err != nil {
		s.T().Fatal(err)
	}

	res, err = s.db.Exec("INSERT INTO article(title, content, author) VALUES ('red', 'Sci-Fi', 'yan')")
	if err != nil {
		s.T().Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		s.T().Fatal(err)
	}

	article, err := s.store.GetArticle(uint64(lastId))
	if err != nil {
		s.T().Fatal(err)
	}

	if article == nil {
		s.T().Errorf("invalid article, got [%v]", article)
	}

	fmt.Printf("retrieved article id: [%v], created article id: [%v] \n", article.Id, lastId)

	if article.Id != uint64(lastId) {
		s.T().Errorf("incorrect article id, got [%v], want [%v]", article.Id, lastId)
	}
}

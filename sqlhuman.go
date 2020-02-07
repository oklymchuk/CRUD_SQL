package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SQLHuman struct {
	db *sql.DB
}

func newSQLHuman() *SQLHuman {
	database, err := sql.Open("mysql", "root:root@/mydb")
	if err != nil {
		log.Fatal(err)
	}

	return &SQLHuman{
		db: database,
	}
}

func (sqlhuman *SQLHuman) Add(w http.ResponseWriter, r *http.Request) {
	var human Human

	err := json.NewDecoder(r.Body).Decode(&human)
	if err != nil {
		log.Println(err)
	}

	human.ID = uuid.New().String()
	res, err := sqlhuman.db.Exec("INSERT INTO User(ID,Firstname,Lastname,Age) VALUES(?,?,?,?)", human.ID, human.Firstname, human.Lastname, human.Age)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
	}
	if rows != 1 {
		log.Printf("expected to insert 1 row, insert %d", rows)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func (sqlhuman *SQLHuman) BeforeQueryGet(tName string) (rPP int, rC int) {

	rowCount := 0
	err := sqlhuman.db.QueryRow("SELECT COUNT(*) FROM ?", tName).Scan(&rowCount)

	if err == sql.ErrNoRows {
		log.Printf("no data in table %q\n", tName)
	} else if err != nil {
		log.Printf("query error: %v\n", err)
	}
	return FileConfig.LinesPerPage, rowCount
}

func (sqlhuman *SQLHuman) GetAll(w http.ResponseWriter, r *http.Request) {
	var humanity []Human
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["PAGE"])
	if err != nil {
		log.Println(err)
		page = 1
	}

	linePerPage, rowCount := sqlhuman.BeforeQueryGet("User")

	if rowCount > 0 {
		rows, err := sqlhuman.db.Query("SELECT ID,Firstname,Lastname,Age FROM User LIMIT ? OFSET ?", linePerPage, page*linePerPage)
		if err != nil {
			log.Println(err)
		}

		for rows.Next() {
			var h Human
			err = rows.Scan(&h.ID, &h.Firstname, &h.Lastname, &h.Age)

			if err != nil {
				log.Println(err)
			}

			humanity = append(humanity, h)
		}

		err = json.NewEncoder(w).Encode(humanity)
		if err != nil {
			log.Println(err)
		}
		w.Header().Add("X-Paging-PageNo", string(page))
		w.Header().Add("X-Paging-PageSize", string(linePerPage))
		w.Header().Add("X-Paging-PageCount", string(rowCount/linePerPage))
		w.Header().Add("X-Paging-TotalRecordCount", string(rowCount))
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (sqlhuman *SQLHuman) GetOne(w http.ResponseWriter, r *http.Request) {
	var h Human

	params := mux.Vars(r)

	rows, err := sqlhuman.db.Query("SELECT ID,Firstname,Lastname,Age FROM User WHERE ID=?", params["ID"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for rows.Next() {
		err = rows.Scan(&h.ID, &h.Firstname, &h.Lastname, &h.Age)
		if err != nil {
			log.Println(err)
		}
	}

	err = json.NewEncoder(w).Encode(h)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (sqlhuman *SQLHuman) UpdateOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var h Human

	err := json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		log.Println(err)
	}

	h.ID = params["ID"]

	res, err := sqlhuman.db.Exec("UPDATE User SET Firstname=?,Lastname=?,Age=? WHERE ID=?", h.Firstname, h.Lastname, h.Age, h.ID)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
	}
	if rows != 1 {
		log.Printf("expected to update 1 row, update %d", rows)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (sqlhuman *SQLHuman) DeleteOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	res, err := sqlhuman.db.Exec("UPDATE User SET Salted=1 WHERE ID=?", params["ID"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotImplemented)
	}
	if rows != 1 {
		log.Printf("expected to delete 1 row, delete %d", rows)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

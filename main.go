package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
}

func initDB() {
	var err error

	db, err = gorm.Open(mysql.Open("root:admin@tcp(127.0.0.1:3306)/gorm"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashPassword)
	err = db.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created SuccessFully")
}
func login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var storedUser User
	result := db.Where("username = ?", user.Username).First(&storedUser)
	if result.Error != nil {
		http.Error(w, "User Not found", http.StatusNotFound)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Wrong Password", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func main() {
	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/signup", signup).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	fmt.Println("server started at, 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

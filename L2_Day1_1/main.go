package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
"/info"   // Информация об API
"/first"  // Случайное число
"/second" // Случайное число
"/add"    // Сумма двух случайных чисел
"/sub"    // Разность
"/mul"    // Произведение
"/div"    // Деление


*результат вернуть в виде JSON

"math/rand"
number := rand.Intn(100)

// Queries
GET http://127.0.0.1:1234/first
GET http://127.0.0.1:1234/second
GET http://127.0.0.1:1234/add
GET http://127.0.0.1:1234/sub
GET http://127.0.0.1:1234/mul
GET http://127.0.0.1:1234/div
GET http://127.0.0.1:1234/info
*/

var (
	vFirst  int = -1
	vSecond int = -1
)

func Rfirst(w http.ResponseWriter, r *http.Request) {
	vFirst = rand.Intn(100)

	HandleSetOK(w, r, "Установлено первое значение равное "+strconv.Itoa(vFirst))
}

func Rsecond(w http.ResponseWriter, r *http.Request) {
	vSecond = rand.Intn(100)
	HandleSetOK(w, r, "Установлено первое значение равное "+strconv.Itoa(vSecond))
}

func Radd(w http.ResponseWriter, r *http.Request) {
	err := Validate()
	if err != nil {
		HandleSetBadrequest(w, r, err.Error())
		return
	}

	HandleSetOK(w, r, "Сумма равна "+strconv.Itoa(vFirst+vSecond))
}

func Rsub(w http.ResponseWriter, r *http.Request) {
	err := Validate()
	if err != nil {
		HandleSetBadrequest(w, r, err.Error())
		return
	}

	HandleSetOK(w, r, "Разность равна "+strconv.Itoa(vFirst-vSecond))
}

func Rmul(w http.ResponseWriter, r *http.Request) {
	err := Validate()
	if err != nil {
		HandleSetBadrequest(w, r, err.Error())
		return
	}

	HandleSetOK(w, r, "Умножение равно "+strconv.Itoa(vFirst*vSecond))
}

func Rdiv(w http.ResponseWriter, r *http.Request) {
	err := Validate()
	if err != nil {
		HandleSetBadrequest(w, r, err.Error())
		return
	}

	if vSecond == 0 {
		HandleSetBadrequest(w, r, "Деление на 0, сгенерируйте второе число!")
		return
	}

	HandleSetOK(w, r, "Деление равно "+strconv.Itoa(vFirst/vSecond))
}

type Res struct {
	Message   string `json:"message,omitempty"`
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}

func HandleSetOK(w http.ResponseWriter, r *http.Request, mes string) {
	w.WriteHeader(http.StatusOK)

	err := Res{
		Message:   mes,
		ErrorCode: 0,
		Error:     "",
	}
	res, _ := json.Marshal(err)
	w.Write(res)
}

func HandleSetBadrequest(w http.ResponseWriter, r *http.Request, err_m string) {
	w.WriteHeader(http.StatusBadRequest)

	err := Res{
		ErrorCode: 0,
		Error:     err_m,
	}
	res, _ := json.Marshal(err)
	w.Write(res)
}

func Rinfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `<h1>Endpoints:</h1>
		<a href="/first" target="_blank">first</a></br>
		<a href="/second" target="_blank">second</a></br>
		<a href="/add" target="_blank">add</a></br>
		<a href="/sub" target="_blank">sub</a></br>
		<a href="/mul" target="_blank">mul</a></br>
		<a href="/div" target="_blank">div</a></br>
		<a href="/info" target="_blank">info</a></br>
	`)
}

func Validate() (err error) {

	if vFirst == -1 || vSecond == -1 {
		return errors.New("Установите значения переменных")
	}

	return nil
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/first", Rfirst).Methods("GET")
	r.HandleFunc("/second", Rsecond).Methods("GET")
	r.HandleFunc("/add", Radd).Methods("GET")
	r.HandleFunc("/sub", Rsub).Methods("GET")
	r.HandleFunc("/mul", Rmul).Methods("GET")
	r.HandleFunc("/div", Rdiv).Methods("GET")
	r.HandleFunc("/info", Rinfo).Methods("GET")

	http.ListenAndServe(":1234", r)
}

package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)

type User struct {
	UserId   int
	Username string
	Password string
}
type Product struct {
	ProductId int
	Name      string
	Price     float64
	Info      string
	Rating    float64
}

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	password = "postgres"
	dbname   = "asignment-1"
)

var (
	db  *sql.DB
	cur = User{}
)

func database() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/register.html")
	tmpl.Execute(w, nil)
}
func home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, cur)
}
func checking(w http.ResponseWriter, r *http.Request) {

	err := db.QueryRow(`select password, id from users where "login" = $1;`, r.FormValue("login")).Scan(&cur.Password, &cur.UserId)
	if r.FormValue("type") == "Login" {
		cur.Username = r.FormValue("login")
		if err == nil && r.FormValue("password") == cur.Password {
			http.Redirect(w, r, "/home/", http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "Incorrect username or password.")
		}
	} else {
		err := db.QueryRow(`select login, id from users where "login" = $1;`, r.FormValue("login")).Scan(&cur.Username, &cur.UserId)
		tmpl := template.New("data")
		message := ""
		if err != nil {
			_ = db.QueryRow(`insert into users("login", "password") values ($1, $2);`, r.FormValue("login"), r.FormValue("password"))

			message = "Successfully registered."
		} else {
			message = fmt.Sprintf(`User %s already exists.\n`, cur.Username)
		}
		tmpl.Parse(`<h1>{{.}}</h1>`)
		tmpl.Execute(w, message)
	}

}
func addProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/add_product.html")
	tmpl.Execute(w, cur)
}
func saveProduct(w http.ResponseWriter, r *http.Request) {
	_, _ = db.Exec(`insert into products(Name, Price, Info,  user_id) values ($1,$2,$3,$4 )`, r.FormValue("name"), r.FormValue("price"), r.FormValue("info"), cur.UserId)
	http.Redirect(w, r, "/home/", http.StatusSeeOther)
}
func searchProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/search_product.html")
	row, err := db.Query(`select  ProductId,Name, Price, Info, Rating from products where Name = $1 order by $2`, r.FormValue("name"), r.FormValue("filter"))

	var products []Product
	temp := Product{}
	if err == nil {
		defer row.Close()
		for row.Next() {
			err = row.Scan(&temp.ProductId, &temp.Name, &temp.Price, &temp.Info, &temp.Rating)
			if err != nil {
				panic(err)
			}
			products = append(products, temp)
			//println(temp.Price)
		}
	}
	tmpl.Execute(w, products)
}

func main() {
	rtr := mux.NewRouter()
	http.Handle("/", rtr)

	db = database()
	defer db.Close()
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/register/", register)
	rtr.HandleFunc("/home/", home)
	rtr.HandleFunc("/checking/", checking)
	rtr.HandleFunc("/add_product/", addProduct)
	rtr.HandleFunc("/save_product/", saveProduct)
	rtr.HandleFunc("/search_product/", searchProduct)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/register/static/", http.StripPrefix("/register/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}

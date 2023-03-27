// package main
//
// import (
//
//	"database/sql"
//	"fmt"
//	_ "github.com/lib/pq"
//
// )
//
// const (
//
//	host     = "localhost"
//	port     = 5432
//	username = "postgres"
//	password = "postgres"
//	dbname   = "asignment-1"
//
// )
//
//	type User struct {
//		UserId   int
//		Username string
//		Password string
//	}
//
//	type Product struct {
//		ProductId int
//		Name      string
//		Price     float64
//		Info      string
//		Rating    float64
//	}
//
// var (
//
//	db     *sql.DB
//	cur    = User{}
//	temp   = Product{Info: ""}
//	filter = "Price"
//	input  string
//
// )
//
//	func main() {
//		fmt.Println("Welcome to the Web-Shop")
//		db = database()
//		for {
//			if loginPage() {
//				homePage()
//			} else {
//				break
//			}
//		}
//		err := db.Close()
//		if err != nil {
//			return
//		}
//	}
//
//	func database() *sql.DB {
//		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//			"password=%s dbname=%s sslmode=disable",
//			host, port, username, password, dbname)
//
//		db, err := sql.Open("postgres", psqlInfo)
//		if err != nil {
//			panic(err)
//		}
//
//		err = db.Ping()
//		if err != nil {
//			panic(err)
//		}
//		return db
//	}
//
//	func homePage() {
//		fmt.Println("1)Add product")
//		fmt.Println("2)View products")
//		fmt.Println("3)Exit")
//		fmt.Print("Enter your choice: ")
//
//		var option string
//		_, _ = fmt.Scan(&option)
//
//		switch option {
//		case "1":
//			addProduct()
//		case "2":
//			viewProducts()
//		case "3":
//			return
//		default:
//			fmt.Println("wrong option")
//		}
//
//		homePage()
//	}
//
//	func loginPage() bool {
//		for {
//
//			fmt.Println("1. Register")
//			fmt.Println("2. Login")
//			fmt.Println("3. Exit")
//			fmt.Print("Enter your choice: ")
//
//			_, _ = fmt.Scan(&input)
//
//			switch input {
//			case "1":
//				register()
//			case "2":
//				if login() {
//					return true
//				}
//			case "3":
//				return false
//			default:
//				fmt.Println("Invalid choice. Try again.")
//			}
//		}
//	}
//
//	func register() {
//		fmt.Print("Enter your username: ")
//		_, _ = fmt.Scan(&cur.Username)
//		username := ""
//		err := db.QueryRow(`select login from users where "login" = $1;`, cur.Username).Scan(&username)
//		if err != nil {
//			_ = db.QueryRow(`insert into users("login", "password") values ($1, $2) returning id;`, cur.Username, cur.Password).Scan(&cur.UserId)
//
//			fmt.Print("Enter your password: ")
//			_, _ = fmt.Scan(&cur.Password)
//			fmt.Println("Successfully registered.")
//			return
//		}
//		fmt.Printf("User %s already exists.\n\n", cur.Username)
//
// }
//
//	func login() bool {
//		fmt.Print("Enter your username: ")
//		_, _ = fmt.Scan(&cur.Username)
//		var password string
//		err := db.QueryRow(`select password,id from users where "login" = $1;`, cur.Username).Scan(&password, &cur.UserId)
//
//		fmt.Print("Enter your password: ")
//		_, _ = fmt.Scan(&cur.Password)
//		if err == nil && password == cur.Password {
//			return true
//		}
//		fmt.Print("Incorrect username or password.\n\n")
//		return false
//	}
//
//	func addProduct() {
//		fmt.Println("Enter product Name:")
//		_, _ = fmt.Scan(&temp.Name)
//
//		fmt.Println("Enter product Price:")
//		_, _ = fmt.Scan(&temp.Price)
//
//		fmt.Println("Enter product Description:")
//		_, _ = fmt.Scan(&temp.Info)
//
//		_, _ = db.Exec(`insert into products(Name, Price, Info,  user_id) values ($1,$2,$3,$4 )`, temp.Name, temp.Price, temp.Info, cur.UserId)
//		fmt.Println("Product added")
//	}
//
//	func viewProducts() {
//		for {
//			fmt.Println("1)Search")
//			fmt.Println("2)Set Filters")
//			fmt.Println("3)Exit")
//			fmt.Print("Enter your choice: ")
//
//			_, _ = fmt.Scan(&input)
//			switch input {
//			case "1":
//				search()
//			case "2":
//				setFilter()
//			case "3":
//				return
//			default:
//				fmt.Println("Invalid choice. Try again.")
//			}
//		}
//	}
//
//	func search() {
//		fmt.Println("Enter product Name:")
//		var productName string
//		_, _ = fmt.Scan(&productName)
//		row, err := db.Query(`select  ProductId,Name, Price, Info, Rating from products where Name = $1 order by $2`, productName, filter)
//		i := 1
//		productIds := []int{}
//		if err == nil {
//			defer row.Close()
//			for row.Next() {
//				err = row.Scan(&temp.ProductId, &temp.Name, &temp.Price, &temp.Info, &temp.Rating)
//				if err != nil {
//					panic(err)
//				}
//				productIds = append(productIds, temp.ProductId)
//				fmt.Printf("\n%d)Name: %s\nPrice: %f\nDescription: %s\nRating: %f\n", i, temp.Name, temp.Price, temp.Info, temp.Rating)
//				i++
//			}
//		}
//		if i == 1 {
//			fmt.Println("No results")
//		} else {
//			if action() {
//				fmt.Println("Select product number:")
//				var j, k int
//				_, _ = fmt.Scan(&j)
//				if j >= i || j < 1 {
//					fmt.Println("Invalid number!")
//				} else {
//					fmt.Println("Put rating:")
//					_, _ = fmt.Scan(&k)
//					db.Exec(`update products set Rating = $1 where ProductId = $2`, k, productIds[j-1])
//					action()
//				}
//			}
//		}
//	}
//
//	func setFilter() {
//		fmt.Println("Current Filter: ", filter)
//		fmt.Println("1)Price\n2)Rating\nSelect filter:")
//		_, _ = fmt.Scan(&input)
//		switch input {
//		case "1":
//			filter = "Price"
//		case "2":
//			filter = "Rating"
//		default:
//			fmt.Println("Wrong option")
//		}
//	}
//
//	func action() bool {
//		for {
//
//			fmt.Println("1)Rate product")
//			fmt.Println("2)Exit")
//			_, _ = fmt.Scan(&input)
//			switch input {
//			case "1":
//				return true
//			case "2":
//				return false
//			default:
//				fmt.Println("Invalid choice. Try again.")
//			}
//		}
//	}
package main

func main() {
	var val int
	for i := 0; i < 10; i++ {
		go func() {
			val++
		}()
	}
	print(val)
}

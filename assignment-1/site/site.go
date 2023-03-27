package main

import (
	"fmt"
)

type User struct {
	Username string
	Password string
}

var Users []User

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

var Items []Item

func main() {
	fmt.Println("\nWelcome to my Store\n")

	for true {
		fmt.Print("1. Login\n2. Register\n3. Break\nEnter your choice: ")

		var index string
		fmt.Scan(&index)

		switch index {
		case "1":
			Login()

		case "2":
			Register()

		case "3":
			fmt.Println("See you soon!")
			return
		case "admin":
			fmt.Println("You not admin!!!")
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func Register() {
	fmt.Print("Enter your username: ")
	var name, pass string
	fmt.Scan(&name)
	fmt.Print("Enter your password: ")
	fmt.Scan(&pass)

	newUser := User{name, pass}
	Users = append(Users, newUser)
	fmt.Println("Successfully registered.")
}

func Login() {
	fmt.Print("Enter your username: ")
	var name, pass string
	fmt.Scan(&name)
	fmt.Print("Enter your password: ")
	fmt.Scan(&pass)
	for _, user := range Users {
		if user.Username == name && user.Password == pass {
			fmt.Println("Successfully logged in.")
			for true {
				fmt.Print("1. Add item\n2. Searching items\n3. Filtering items\n4. Giving rating\n0. Exit\nEnter your choice: ")

				var index string
				fmt.Scan(&index)
				switch index {
				case "1":
					AddItem()
				case "2":
					Searching()
				case "3":
					Filter()
				case "4":
					GiveRating()
				case "0":
					return
				}
			}
		}
	}
	fmt.Println("Incorrect username or password.")
}

func AddItem() {
	var name string
	var price, rating float64
	fmt.Print("Enter item name: ")
	fmt.Scan(&name)
	fmt.Print("Enter item price: ")
	fmt.Scan(&price)
	fmt.Print("Enter item rating: ")
	fmt.Scan(&rating)

	newItem := Item{name, price, rating}
	Items = append(Items, newItem)
	fmt.Println("You have successfully added the item.")
}

func Searching() {
	var name string
	fmt.Print("Enter item name: ")
	fmt.Scan(&name)

	for _, item := range Items {
		if item.Name == name {
			fmt.Println("Search Result: ", item.Name, item.Price, item.Rating)
			return
		}
		fmt.Println("This item is not exist.")
	}
}

func Filter() {
	var price, rating float64
	fmt.Print("Enter item price: ")
	fmt.Scan(&price)
	fmt.Print("Enter item rating: ")
	fmt.Scan(&rating)
	for _, item := range Items {
		if item.Price <= price && item.Rating >= rating {
			fmt.Println(item.Name, item.Price, item.Rating)
		}
	}
}
func GiveRating() {
	var name string
	var rating float64
	fmt.Print("Enter item name: ")
	fmt.Scan(&name)
	fmt.Print("Enter item rating: ")
	fmt.Scan(&rating)

	for _, item := range Items {
		if item.Name == name {
			item.Rating = rating
			fmt.Println("Rating edited successfully.")
			fmt.Println(item.Name, item.Price, item.Rating)
			return
		}
	}
}

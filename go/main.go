package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// User represents a user's data
type User struct {
	Name    string
	Results []string
}

// Map to store user data
var users map[string]*User

func init() {
	users = make(map[string]*User)
}

// homeHandler handles the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		if name != "" {
			user := getUser(name)
			if user != nil {
				renderResultsPage(w, user)
				return
			}
			user = &User{Name: name}
			users[name] = user
			renderInputPage(w, user)
			return
		}
	}
	renderHomePage(w)
}

// inputHandler handles the input page
func inputHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	num1 := r.FormValue("num1")
	num2 := r.FormValue("num2")

	if name != "" && num1 != "" && num2 != "" {
		user := getUser(name)
		if user != nil {
			result := processInput(num1, num2)
			user.Results = append(user.Results, result)
			renderResultsPage(w, user)
			return
		}
	}
	renderHomePage(w)
}

// getUser retrieves user data from the map
func getUser(name string) *User {
	return users[name]
}

// processInput performs addition and multiplication
func processInput(num1, num2 string) string {
	n1 := parseInt(num1)
	n2 := parseInt(num2)
	addition := n1 + n2
	multiplication := n1 * n2
	return fmt.Sprintf("Addition: %d, Multiplication: %d", addition, multiplication)
}

// parseInt converts string to integer
func parseInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

// renderHomePage renders the home page
func renderHomePage(w http.ResponseWriter) {
	tmpl := `<html>
<head><title>Calculator App</title></head>
<body>
<h1>Welcome to the Calculator App!</h1>
<form action="/" method="post">
  <label for="name">Please enter your name:</label><br>
  <input type="text" id="name" name="name"><br>
  <input type="submit" value="Submit">
</form>
</body>
</html>`
	fmt.Fprint(w, tmpl)
}

// renderInputPage renders the input page
func renderInputPage(w http.ResponseWriter, user *User) {
	tmpl := `<html>
<head><title>Calculator App</title></head>
<body>
<h1>Welcome back, {{.Name}}!</h1>
<form action="/input" method="post">
  <input type="hidden" id="name" name="name" value="{{.Name}}">
  <label for="num1">Please enter first number:</label><br>
  <input type="number" id="num1" name="num1"><br>
  <label for="num2">Please enter second number:</label><br>
  <input type="number" id="num2" name="num2"><br>
  <input type="submit" value="Submit">
</form>
</body>
</html>`
	t := template.Must(template.New("input").Parse(tmpl))
	err := t.Execute(w, user)
	if err != nil {
		log.Println("Error rendering template:", err)
	}
}

// renderResultsPage renders the results page
func renderResultsPage(w http.ResponseWriter, user *User) {
	tmpl := `<html>
<head><title>Calculator App</title></head>
<body>
<h1>Welcome back, {{.Name}}!</h1>
<h2>Previous Search Results:</h2>
<ul>
  {{range .Results}}
  <li>{{.}}</li>
  {{end}}
</ul>
<h2>Enter new numbers:</h2>
<form action="/input" method="post">
  <input type="hidden" id="name" name="name" value="{{.Name}}">
  <label for="num1">Please enter first number:</label><br>
  <input type="number" id="num1" name="num1"><br>
  <label for="num2">Please enter second number:</label><br>
  <input type="number" id="num2" name="num2"><br>
  <input type="submit" value="Submit">
</form>
</body>
</html>`
	t := template.Must(template.New("results").Parse(tmpl))
	err := t.Execute(w, user)
	if err != nil {
		log.Println("Error rendering template:", err)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/input", inputHandler)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

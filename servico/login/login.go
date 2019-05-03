package login

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
)

func Inicializa(router *mux.Router) {
	router.HandleFunc("/api/login", loginHandler)
	router.HandleFunc("/api/user", userHandler)
	router.HandleFunc("/login", indexPageHandler)
}

const indexPage = `
<h1>Login</h1>
<form method="post" action="/api/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func loginHandler(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	pass := req.FormValue("password")
	redirectTarget := "/"
	if name != "" && pass != "" {
		// .. check credentials ..
		setSession(name, res)
		redirectTarget = "/#/"
	}
	http.Redirect(res, req, redirectTarget, 302)
}

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, indexPage)
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func getUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("session"); err == nil {
		fmt.Println(cookie)
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			fmt.Println(cookieValue)
			userName = cookieValue["name"]
		}
	}
	return userName
}

func userHandler(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		fmt.Fprint(response, userName)
	} /* else {
		http.Redirect(response, request, "/login", 302)
	}*/
}

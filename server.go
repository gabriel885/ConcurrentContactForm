package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	port       = 9090
	emailRoute = "/send-email"
)

func serveContactForm(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" { // server http form
		http.ServeFile(w, r, "contact-form.html")
	} else { // return method not allowed
		http.Error(w, "Method Now Allowed", http.StatusMethodNotAllowed)
	}
}

func sendMail(w http.ResponseWriter, r *http.Request) { // contact form handler
	if r.Method == "POST" {
		r.ParseForm()
		form := r.Form // form parameters of type url.Values

		conf := LoadSmtpConfigurations("config.json")

		go func() {
			conf.HandleSendMail(form) // send email
		}()

		fmt.Fprintf(w, "Email sent.") // write response
		http.Redirect(w, r, "/", 200)
	}
	http.Redirect(w, r, "/", 404)

}

func main() {

	// set up http routes handlers
	http.HandleFunc("/", serveContactForm)
	http.HandleFunc(emailRoute, sendMail) // set email route
	log.Println("listening on :", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil) // setting listening port
	if err != nil {
		log.Fatalf("Couldn't listen to port %d", port)
	}
	fmt.Printf("Running server on port %d", port)

}

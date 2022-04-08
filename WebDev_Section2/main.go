package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to this awesome channel!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>Contact us at: <a href=\"mailto:dw@gmail.com\">dw@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {

	// If you only want to serve 1 file and not a full directory, you can use http.ServeFile
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//     http.ServeFile(w, r, "index.html")
	// })

	handler := http.StripPrefix("/faq", http.FileServer(http.Dir("./static")))
	handler.ServeHTTP(w, r)
}

// CASE - 1, 2	// Standalone handler created to handle different routes
// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("content-type", "text/html; charset=utf-8")
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		message := fmt.Sprintf("<h1> %d - Page Not Found</h1><p>Please redirect to:<ul><li><a href=\"\\\">Home</a></li><li><a href=\"\\contact\">Contact</a></li></ul></p>", http.StatusNotFound)
// 		fmt.Fprint(w, message)
// 		// http.Error(w, message, http.StatusNotFound)	// only sends plain text error
// 	}
// }

// CASE - 1
// func main() {
// 	var router http.HandlerFunc
// 	router = pathHandler
// 	http.ListenAndServe(":3000", router)
// }

// CASE - 2
// func main() {
// 	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
// }

// CASE - 3, 4, 5
type Router struct{}

func (route Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		message := fmt.Sprintf("<h1> %d - Page Not Found</h1><p>Please redirect to:<ul><li><a href=\"\\\">Home</a></li><li><a href=\"\\contact\">Contact</a></li><li><a href=\"\\faq\">FAQs</a></li></ul></p>", http.StatusNotFound)
		fmt.Fprint(w, message)
	}
}

// CASE - 3 		// Define the router and pass directly to port
// func main() {
// 	var router Router
// 	http.ListenAndServe(":3000", router)
// }

// CASE - 4		// Define the router and pass to Handle(). Helpful in cases where multiple routers are created for specific purposes, but listening to same port.
// func main() {
// 	var router Router
// 	http.Handle("/", router)
// 	http.ListenAndServe(":3000", nil)
// }

// CASE - 5		// Similar to case 2; just passing the particular function directly to HandleFunc()
func main() {
	var router Router
	http.HandleFunc("/", router.ServeHTTP)
	http.ListenAndServe(":3000", nil)
}

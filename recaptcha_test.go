package recaptcha

import (
	"fmt"
	"log"
	"net/http"

	"github.com/haisum/recaptcha"
)

func ExampleR_Verify() {
	sitekey := "{Your site key here}"
	re := recaptcha.R{
		Secret: "{Your secret here}",
	}

	form := fmt.Sprintf(`
		<html>
			<head>
				<script src='https://www.google.com/recaptcha/api.js'></script>
			</head>
			<body>
				<form action="/submit" method="post">
					<div class="g-recaptcha" data-sitekey="%s"></div>
					<input type="submit">
				</form>
			</body>
		</html>
	`, sitekey)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, form)
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		isValid := re.Verify(r)
		if isValid {
			fmt.Fprintf(w, "Valid")
		} else {
			fmt.Fprintf(w, "Invalid! These errors ocurred: %v", re.LastError())
		}
	})

	err := http.ListenAndServe(":8100", nil)

	if err != nil {
		log.Printf("Could not start server. %s", err)
	}
}

package recaptcha

import (
	"net/http"
)

type Recaptcha struct {
	Secret string
}

var url string = "https://www.google.com/recaptcha/api/siteverify"

func (r Recaptcha) Verify(req http.Request) {

}

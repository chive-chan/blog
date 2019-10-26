package web

import (
	"fmt"
	"github.com/MccGithub/blog/web/details"
	"github.com/MccGithub/blog/web/index"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Opt struct {
	Address		string

	DBDriver 	string
	DBConn		string
}

func home(w http.ResponseWriter, r *http.Request) {

}

func Serve(opt Opt) error {
	handler := chi.NewRouter()

	logrus.Tracef("%+v", opt)

	handler.Mount("/", index.Router())
	handler.Mount("/details", details.Router())

	fmt.Println("Listening at ", opt.Address)
	return http.ListenAndServe(opt.Address, handler)
}
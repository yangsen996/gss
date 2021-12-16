package example

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/yangsen996/gss"
)

func TestGss(t *testing.T) {
	r := gss.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "url.path=%q\n", req.URL.Path)
	})
	r.Run(":9999")
}

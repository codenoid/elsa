package ops

import (
	"log"

	"github.com/elsaland/quickjs"
	"github.com/imroc/req"
)

func Fetch(ctx *quickjs.Context, url quickjs.Value) quickjs.Value {
	r, err := req.Get(url.String())
	if err != nil {
		log.Fatal(err)
	}

	resp, err := r.ToString()
	if err != nil {
		log.Fatal(err)
	}
	return ctx.String(resp)
}

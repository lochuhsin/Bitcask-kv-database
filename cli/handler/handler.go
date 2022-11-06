package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"rebitcask/src"
)

type QueryHandler struct{}

func (query *QueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	switch r.Method {
	case "GET":
		res, status := query.Get(r.Form)
		if !status {
			fmt.Fprintf(w, "400")
		} else {
			fmt.Fprintf(w, res)
		}
	case "POST":
		err := query.Post(r.Form)
		if err != nil {
			fmt.Fprintf(w, "400")
		} else {
			fmt.Fprintf(w, "ok")
		}
	case "PATCH":
		err := query.Post(r.Form)
		if err != nil {
			fmt.Fprintf(w, "400")
		} else {
			fmt.Fprintf(w, "ok")
		}

	case "DELETE":
		err := query.Post(r.Form)
		if err != nil {
			fmt.Fprintf(w, "400")
		} else {
			fmt.Fprintf(w, "ok")
		}
	default:
		fmt.Fprintf(w, "400")
	}

}

func (query *QueryHandler) Get(rawQuery url.Values) (result string, status bool) {
	if key, ok := rawQuery["key"]; ok {
		return src.Get(key[0])
	}
	return "", false
}

func (query *QueryHandler) Post(rawQuery url.Values) error {
	key, keyStatus := rawQuery["key"]
	val, valStatus := rawQuery["val"]
	fmt.Println(keyStatus, valStatus)
	if keyStatus && valStatus {
		return src.Set(key[0], val[0])
	}
	return errors.New("invalid query")
}

func (query *QueryHandler) Patch(rawQuery url.Values) error {
	key, keyStatus := rawQuery["key"]
	val, valStatus := rawQuery["val"]

	if keyStatus && valStatus {
		return src.Set(key[0], val[0])
	}
	return errors.New("invalid query")
}

func (query *QueryHandler) Delete(rawQuery url.Values) error {
	if key, keyStatus := rawQuery["key"]; keyStatus {
		return src.Delete(key[0])
	}
	return errors.New("invalid query")
}

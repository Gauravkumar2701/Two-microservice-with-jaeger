package handler

import (
	"ProxyService/models"
	"ProxyService/tracing/tracing"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"
)



var itemIDKey = "itemID"

func student(router chi.Router) {
	router.Get("/", getAllStudent)
}
func Do(req *http.Request) (string, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}
	return string(body), nil
}
func getAllStudent(w http.ResponseWriter, r *http.Request) {
	//res, err := http.Get( "http://localhost:8080/student" )
	// read all response body
	span := tracing.StartSpanFromRequest(gTracer, r)
	defer span.Finish()

	//ctx := opentracing.ContextWithSpan(context.Background(), span)
	req,_:= http.NewRequest("GET","http://localhost:8080/student" , nil)
	tracing.Inject(span, req)
	res,err:= Do(req)

	//data, _ := ioutil.ReadAll( res.Body )

	// close response body
	//res.Body.Close()

	// print `data` as a string
	//fmt.Printf( "%s\n", data )
	var temp2 models.ItemList
	json.Unmarshal([]byte(res), &temp2)
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, &temp2); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}
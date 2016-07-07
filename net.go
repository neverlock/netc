package main

/*
typedef struct http_request {
	        char *method;
	        char *path;
		char *query;
		} http_request;

typedef void (*handler_func)(http_request *);

extern inline void call_handler(handler_func fn, http_request *req) {
		fn(req);
}

*/
import "C"
import (
	"net/http"
	"unsafe"
)

//export http_handle_func
func http_handle_func(path *C.char, handler unsafe.Pointer) {
	http.HandleFunc(C.GoString(path), func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var req C.struct_http_request
		req.method = C.CString(r.Method)
		req.path = C.CString(r.URL.Path)
		req.query = C.CString(r.Form.Get("data"))
		C.call_handler(C.handler_func(handler), &req)
	})
}

//export http_listen_and_serve
func http_listen_and_serve(listen *C.char) {
	http.ListenAndServe(C.GoString(listen), nil)
}

func main() {}

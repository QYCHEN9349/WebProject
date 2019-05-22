package webUtil

import "net/http"

type RouterHandler struct {
}

var routerMap = make(map[string]func(w http.ResponseWriter, r *http.Request))

func (p *RouterHandler) Router(urlPath string, handler func(w http.ResponseWriter, r *http.Request)) {
	routerMap[urlPath] = handler
}

//转发请求
func (p *RouterHandler) RequestDispatch(w http.ResponseWriter, r *http.Request) {
	if function, ok := routerMap[r.URL.Path]; ok {
		function(w, r)
	}
}

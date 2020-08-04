package cmd

func TestHealthLogic(t *Testing.T) {
	h := healthzHandler{d: d}
	h.ServeHTTP(fakeResponseWriter, fakeRequest)
}

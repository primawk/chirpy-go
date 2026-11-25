package main

type returnErr struct {
	Error string `json:"error"`
}

type validateResp struct {
	Body string `json:"body"`
}

type returnCleaned struct {
	Cleaned_body string `json:"cleaned_body"`
}

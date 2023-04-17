package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type PageData struct {
    Title string
    Name  string
}

type FormData struct {
    Name     string
    Email    string
    PhoneNum string
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if r.Method != http.MethodPost {

            data := PageData{
                Title: "Main form",
                Name:  "",
            }


            templatePath := "static/index.html"

            tmpl, err := template.ParseFiles(templatePath)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            err = tmpl.Execute(w, data)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            return
        }

        name := r.FormValue("name")
        email := r.FormValue("email")
        phoneNum := r.FormValue("phoneNum")


//        data := PageData{
//            Title: "Результаты отправки формы",
//            Name:  name,
//        }


        formData := FormData{
            Name:     name,
            Email:    email,
            PhoneNum: phoneNum,
        }


        templatePath := "static/results.html"


        tmpl, err := template.ParseFiles(templatePath)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        err = tmpl.Execute(w, formData)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    })

    fmt.Println("Starting server on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

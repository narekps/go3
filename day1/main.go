package main

import (
	"html/template"
	"log"
	"net/http"
)

type Profile struct {
	Name            string
	Gender          string
	Age             int
	DOB             string
	City            string
	MetroStation    string
	Contacts        []*Contact
	Position        string
	Specializations []string
	WorkSchedule    string
	RemoteWork      bool
	Experience      int
	Experiences     []*Experience
}

type Contact struct {
	Value   string
	Default bool
}

type Experience struct {
	DateFrom    string
	DateTo      string
	Duration    string
	Company     string
	CompanySite string
	City        string
	Position    string
	Description string
}

var profile = &Profile{
	Name:         "Погосян Нарек Суренович",
	Gender:       "Мужчина",
	Age:          32,
	DOB:          "28 июля 1991",
	City:         "Москва",
	MetroStation: "Молодежная",
	Contacts: []*Contact{
		&Contact{
			Value:   "+7 (926) 123-45-67",
			Default: true,
		},
		&Contact{
			Value: "narek.p@live.ru",
		},
	},
	Position:        "Senior Golang developer",
	Specializations: []string{"Программист, разработчик"},
	WorkSchedule:    "Полный день",
	RemoteWork:      true,
	Experience:      13,
	Experiences: []*Experience{
		&Experience{
			DateFrom:    "Апрель 2016",
			DateTo:      "Июнь 2017",
			Duration:    "1 год 2 месяца",
			Company:     "Google",
			CompanySite: "https://google.ru",
			City:        "Москва",
			Position:    "Senior Golang developer",
			Description: "Разработка сервисов для Google",
		},
		&Experience{
			DateFrom:    "Апрель 2015",
			DateTo:      "Январь 2016",
			Duration:    "1 год",
			Company:     "Яндекс",
			CompanySite: "https://ya.ru",
			City:        "Москва",
			Position:    "Senior Golang developer",
			Description: "Разработка сервисов для Яндекса",
		},
	},
}

func main() {
	http.HandleFunc("/", profileHandler)
	http.HandleFunc("/experience", experienceHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/layout.gohtml", "web/template/profile.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, profile)
}

func experienceHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/layout.gohtml", "web/template/experience.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, profile)
}

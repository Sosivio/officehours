package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Title string
	Image string
}

type TodoPageData struct {
	PageTitle  string
	AppVersion string
	Products   []Product
}

func runHomePage(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("assets/html/home.html"))
	data := TodoPageData{PageTitle: "My TODO list",
		AppVersion: AppVersion,
		Products: []Product{
			{Title: "Kingston DT100G3 8GB 16GB 32GB 64GB Data Traveler 100 G3 USB 3.0 Flash Pen Drive", Image: "assets/img/s-l225.webp"},
			{Title: "SanDisk 128GB 64GB 32GB 16GB 8GB Cruzer Blade CZ50 USB Flash Pen Drive OTG Lot", Image: "assets/img/s-l226.webp"},
			{Title: "Flash Memory 64GB 32GB 16GB 8GB 4GB USB 2.0 Stick Pen Drive U Disk Swivel Key", Image: "assets/img/s-l226.webp"},
			{Title: "Kingston 128GB 64GB 32GB 16GB 8GB DT 50 Flash USB 3.0 3.1 Drive 110MB/s OTG Lot", Image: "assets/img/s-l227.webp"},
			{Title: "SanDisk OTG M3.0 128GB 64GB 32GB 16GB Ultra Dual USB 3.0 Flash Drive 150MBs Lot", Image: "assets/img/s-l228.webp"},
			{Title: "SanDisk 256GB 128GB 64GB 32GB 16GB Cruzer Ultra Flair CZ73 USB 3.0 Drive OTG Lot", Image: "assets/img/s-l229.webp"},
			{Title: "USB Flash Drive (2 TB) High-Speed Data Storage Thumb Stick Store Movies, Picture", Image: "assets/img/s-l230.webp"},
			{Title: "32GB 64GB 128GB SanDisk USB 3.0 Flash Drive 130MB/s Ultra Fit Mini Nano", Image: "assets/img/s-l231.webp"},
			{Title: "Usb flash drive super stick mini pen drive 64gb 32gb 16gb 8gb 4gb", Image: "assets/img/s-l232.webp"},
		},
	}
	tmpl.Execute(w, data)
}

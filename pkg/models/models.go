package models

type GkuData struct {
	Postav string `dbase:"POSTAV"`
	Lschet string `dbase:"LSCHET"`
	Sity   string `dbase:"CITY"`
	Street string `dbase:"STREET"`
	Dom    string `dbase:"DOM"`
	KV     string `dbase:"KV"`
}

package auth

import "time"

type User struct {
	ID           	int
	Fullname     	string
	LegalName    	string
	TempatLahir  	string
	TanggalLahir 	time.Time
	Gaji			int
	FotoKtp			string
	FotoSelfie		string
}
package models

type MessageRecieveRMQ struct {
	Id      int     `json:"Id"`
	UserId  int     
	Product string  
	Price   float32 
	Time    string  
}
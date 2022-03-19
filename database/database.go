package database

var connection string


//function init merupakan function yang akan langsung di jalankan ketika packagenya di import ke package lain
func init(){
	connection = "Mysql"
}

func GetDatabase()string{
	return connection
}
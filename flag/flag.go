package flag

import "flag"

var MongoUrl string
var MongoUsername string
var MongoPassword string
func init(){
	mongoUrl:=flag.String("mongo_url","192.168.2.184:27017","mongodb url")
	mongoUsername:=flag.String("mongo_username","","mongodb username")
	mongoPassword:=flag.String("mongo_password","","mongodb password")
	flag.Parse()
	MongoUrl = *mongoUrl
	MongoUsername = *mongoUsername
	MongoPassword = *mongoPassword
}


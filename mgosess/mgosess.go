package mgosess
import (

	"gopkg.in/mgo.v2"
	"go-alipay/flag"
)

const(
	DB="trade"
)
func OpenSession() *mgo.Session{
	session,err:=mgo.Dial(flag.MongoUrl)
	if err != nil { panic(err) }
	session.SetMode(mgo.Monotonic, true)
	return session
}
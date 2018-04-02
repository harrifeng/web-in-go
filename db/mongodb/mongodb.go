package mongodb

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

//global
var GlobalMgoSession *mgo.Session

func init() {
	globalMgoSession, err := mgo.DialWithTimeout("mongodb://localhost:27027/web-in-go", 10*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(300)
}

func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

package dba

import (
	"gopkg.in/mgo.v2"
)

const (
	MONOLAB_URL = "mongodb://gophermongo:kombatgopher123@ds037451.mongolab.com:37451/gopherkombat" //os.Env!!!
)

func Execute(col string, f func(*mgo.Collection)) {
	session, err := mgo.Dial(MONOLAB_URL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("gopherkombat").C(col)
	f(c)
}

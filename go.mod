module go-repository-pattern

go 1.14

require (
	github.com/labstack/echo/v4 v4.1.16
	github.com/spf13/viper v1.7.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	transport/lib v0.0.0-20200609030708-5cbccf123a48
)

replace transport/lib => gitlab.com/transport4/lib.git v0.0.0-20200609030708-5cbccf123a48

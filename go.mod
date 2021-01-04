module github.com/dora1998/snail-bot

go 1.13

require (
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/golang/mock v1.4.3
	github.com/google/uuid v1.1.4
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rubenv/sql-migrate v0.0.0-20190902133344-8926f37f0bc1
	github.com/spf13/cobra v0.0.7
	github.com/ziutek/mymysql v1.5.4 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/gorp.v1 v1.7.2 // indirect
)

replace github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f => github.com/dora1998/go-twitter v0.0.0-20191014181053-ad41280878c1

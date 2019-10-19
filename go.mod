module snail-bot

go 1.13

require (
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/go-sql-driver/mysql v1.4.0
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	google.golang.org/appengine v1.6.5 // indirect
)

replace github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f => github.com/dora1998/go-twitter v0.0.0-20191014181053-ad41280878c1

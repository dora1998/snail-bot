module snail-bot

go 1.13

require (
	github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f
	github.com/dghubble/oauth1 v0.6.0
	github.com/google/uuid v1.1.1
	github.com/kelseyhightower/envconfig v1.4.0
)

replace github.com/dghubble/go-twitter v0.0.0-20190719072343-39e5462e111f => ../go-twitter

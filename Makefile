code:
	goctl api go -api ./dsl/main.api -dir .  --style=goZero

swg:
	goctl api plugin -plugin goctl-swagger="swagger -filename api.json" -api dsl/main.api -dir ./doc/swagger/

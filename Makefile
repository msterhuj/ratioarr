css:
	@npx @tailwindcss/cli -i app.css -o internal/static/app.css

templ:
	@templ generate

sqlc:
	@sqlc generate

docker:
	@docker build -t ratioarr .

dev:
	@air
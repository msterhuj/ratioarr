css:
	@echo "Generating CSS with Tailwind..."
	@npm install
	@npx --yes @tailwindcss/cli -i app.css -o internal/static/app.css

templ:
	@echo "Generating templ code..."
	@templ generate

sqlc:
	@echo "Generating sqlc code..."
	@sqlc generate

docker:
	@docker build -t ratioarr .

deps: css templ sqlc

dev:
	@air

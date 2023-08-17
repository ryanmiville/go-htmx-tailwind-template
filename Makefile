.PHONY: all
all:
	npx tailwindcss -i ./assets/app.css -o ./public/app.css --minify
	go build

.PHONY: clean
clean:
	rm -f ./public/app.css
	rm -f ./go-htmx-tailwind-template

.PHONY: run
run:
	DEV=true air

.PHONY: css
css:
	npx tailwindcss -i ./assets/app.css -o ./public/app.css --watch
.PHONY: all
all: ./public/app.css
	go build
	
.PHONY: clean
clean:
	rm -f ./public/app.css
	rm -f ./go-htmx-tailwind-template

.PHONY: air-cmd
air-cmd: public/app.css
	go build -o ./tmp/main .
	
.PHONY: run
run:
	air

public/app.css:
	npx tailwindcss -i ./assets/app.css -o ./public/app.css --minify

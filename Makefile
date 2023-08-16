.PHONY: all
all: ./views/output.css
	go build
	
.PHONY: clean
clean:
	rm -f ./views/output.css
	rm -f ./go-htmx-tailwind-template

.PHONY: air
air: views/output.css
	go build -o ./tmp/main .
	
views/output.css:
	npx tailwindcss -i ./views/global.css -o ./views/output.css --minify

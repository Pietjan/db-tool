install-htmx:
	mkdir -p ./web/static/js
	wget -O ./web/static/js/htmx.min.js https://unpkg.com/htmx.org/dist/htmx.min.js 

install-tailwindcss:
	mkdir -p ~/bin
	wget -O ~/bin/tailwindcss https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x ~/bin/tailwindcss

build: 
	go run main.go start
extract:
	tar xvf data.tar.gz
clean:
	rm assets/communes.json assets/country.json assets/districts.json assets/fokontany.json assets/regions.json 
serve:
	go run main.go serve

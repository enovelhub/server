BINDATA_PATH=${PWD}/bindata.go
WEBSITE_PATH=enovelhub-website

all: website bindata

bindata: clear-bindata
	cd ${WEBSITE_PATH} && \
	go-bindata -ignore .git  -o ${BINDATA_PATH} ./...

clear-bindata:
	rm -rf ${BINDATA_PATH}
	
website: clear-website
	git clone https://github.com/enovelhub/enovelhub-website.git ${WEBSITE_PATH}

clear-website: 
	rm -rf ${WEBSITE_PATH}

clear: clear-bindata clear-website

.DEFAULT_GOAL				:= all
name 								:= "infraviz"

all: vendor build tools cover finish

.PHONY: test
test:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests"
	@go test ./drawyed -v

.PHONY: cover
cover: 
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run the unit-tests + coverage"
	@go test ./drawyed -v -covermode=count -coverprofile=coverage.out

cover.upload:
	# for this to get working you have to export the repo_token for your repo at coveralls.io
	# i.e. export INFRA_VIZ_COVERALLS_REPO_TOKEN=<your token>
	@${GOPATH}/bin/goveralls -coverprofile=coverage.out -service=circleci -repotoken=${INFRA_VIZ_COVERALLS_REPO_TOKEN}
	


#-----------------
#-- build
#-----------------
build:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Build the $(name)"
	@go build -o $(name) .

#------------------
#-- dependencies
#------------------
depend.update:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> updating dependencies from Gopkg.lock"
	@dep ensure -update

depend.install:
	@echo "----------------------------------------------------------------------------------"
	@echo "--> install dependencies as listed in Gopkg.toml"
	@dep ensure

#------------------
#-- Tools
#------------------
tools:
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls	

vendor: depend.install

run: build
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run ${name}"
	@./${name}

finish:
	@echo "=================================================================================="
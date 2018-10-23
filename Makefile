.DEFAULT_GOAL				:= all
name 								:= "infraviz"

all: vendor build finish

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

vendor: depend.install

run: build
	@echo "----------------------------------------------------------------------------------"
	@echo "--> Run ${name}"
	@./${name}

finish:
	@echo "=================================================================================="
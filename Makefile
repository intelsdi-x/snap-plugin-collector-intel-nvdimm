default:
	bash -c "./scripts/build.sh"
clean:
	rm intel-nvdimm/constants.go
	rm intel-nvdimm/*_string.go
	rm intel-nvdimm/type/*_string.go
	go clean	

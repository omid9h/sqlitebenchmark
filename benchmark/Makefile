get:
	wrk -t4 -c100 -d10s --timeout 15s --script=./get.lua http://localhost:1323/api/v1/terminals/get
set:
	wrk -t4 -c100 -d10s --timeout 15s --script=./set.lua http://localhost:1323/api/v1/terminals/set
.PHONY: get set
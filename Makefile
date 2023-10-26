.PHONY:gen
gen:
	ls -F | grep '/' | grep -v orm01 | xargs -tI {} sh -c "cd {} && rm -f *.tmpl && go generate types.go"

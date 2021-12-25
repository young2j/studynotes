gen:
	rm -rf ./protos/gen;
	cd protos && buf lint;
	cd protos && buf generate --exclude-path thirdapis/openapiv2,thirdapis/google,thirdapis/validate;
	cp protos/gen/swagger/* swagger/;






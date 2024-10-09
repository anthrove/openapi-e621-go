package openapi

//go:generate {JAVA_HOME} -jar openapi-generator-cli-7.7.0.jar generate -g go -i {openapi.yaml} -o {E:/Anthrove/_github/openapi-e621-go} --additional-properties=packageName=e621,packageVersion=1.0.1,packageName=openapi,enumClassPrefix=true,structPrefix=true

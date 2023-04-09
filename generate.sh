#!/bin/bash
wget https://github.com/5GCoreNet/openapi-generator/raw/master/modules/openapi-generator-cli/target/openapi-generator-cli.jar
for specs in $(ls ./5GC_APIs/*.yaml); do
    package_name=$(basename "${specs%.*}" | cut -d'_' -f2- | tr '-' '_')
    package_name="openapi_$package_name" # Adding the prefix openapi_ to avoid conflicts with numbers in the name (e.g. 5GC_APIs/1.yaml)
    echo "Generating $package_name"
    rm -rf ./"$package_name" # Removing the folder if it already exists
    java -Xmx1024M -jar openapi-generator-cli.jar generate -i "$specs" -g go -o ./"$package_name" --skip-validate-spec --additional-properties=isGoSubmodule=true,packageName="$package_name",useOneOfDiscriminatorLookup=false --global-property apiDocs=false,modelDocs=false,apiTests=false,modelTests=false --openapi-normalizer REF_AS_PARENT_IN_ALLOF=true
    rm -rf ./"$package_name"/go.* # Removing the go.mod and go.sum files
    echo "Done"
done
rm openapi-generator-cli.jar

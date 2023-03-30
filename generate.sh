#!/bin/bash
for specs in $(ls /specs/5GC_APIs/*.yaml); do
    package_name=$(basename "${specs%.*}" | cut -d'_' -f2-)
    package_name="openapi_$package_name" # Adding the prefix openapi_ to avoid conflicts with numbers in the name (e.g. 5GC_APIs/1.yaml)
    echo "Generating $package_name"
    java -Xmx1024M -DloggerPath=conf/log4j.properties -jar opt/openapi-generator/modules/openapi-generator-cli/target/openapi-generator-cli.jar generate -i "$specs" -g go -o /specs/"$package_name"  --skip-validate-spec --additional-properties=onlyInterfaces=true,isGoSubmodule=true,packageName="$package_name"
    rm -rf /specs/"$package_name"/go.* # Removing the go.mod and go.sum files
    rm -rf /specs/"$package_name"/test # Removing the test folder
    rm -rf /specs/"$package_name"/docs # Removing the docs folder as it causes conflicts due to name collisions. Plus, we don't need it and it makes the package bigger.
    echo "Done"
done

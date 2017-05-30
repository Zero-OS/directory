#!/bin/bash
set -e
pushd directory
    echo "Run tests"
    go get ./...
    bash ../codecov.sh
popd
echo "Generate docs"
pushd specs
    raml2html directory.raml > directory.html
popd
echo "Install go-raml"
pushd $GOPATH/src/github.com/Jumpscale/go-raml
    bash install.sh
popd
echo "Validate raml server generation"
go-raml server -l go --api-file-per-method --dir servertmp --ramlfile specs/directory.raml --lib-root-urls https://raw.githubusercontent.com/Jumpscale-Cockpit/raml-definitions/master/
echo "Validate raml client generation"
go-raml client -l python --ramlfile specs/directory.raml --dir clienttmp

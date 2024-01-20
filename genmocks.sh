#!/bin/bash

go install github.com/vektra/mockery/v2@latest
rm -fr tests/mocks
if [ -z "${GOPATH}"]; then
  mockery --dir internal --all --keeptree --with-expecter --output ./tests/mocks
else 
  ${GOPATH}/bin/mockery --dir internal --all --keeptree --with-expecter --output ./tests/mocks
fi
#mv tests/mocks/internal/modules/* tests/mocks/
#rm -fr tests/mocks/internal/modules
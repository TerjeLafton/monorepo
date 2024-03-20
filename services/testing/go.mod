module github.com/terjelafton/monorepo/services/testing

go 1.22.1

require golang.org/x/example/hello v0.0.0-20240205180059-32022caedd6a
require github.com/terjelafton/monorepo/libs/real v0.0.0

replace github.com/terjelafton/monorepo/libs/real => ../../libs/real


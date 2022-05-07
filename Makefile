GO_FILES := $(shell find . -name '*.go' ! -name '*_test.go')

GOLANG := go

.PHONY: run
run: listener
	./listener

.PHONY: run-replay
run-replay: replay
	./replay -path ./session-example.sqlite3 -sid 8714292215998828815 -rate 20

GO_TEST_FLAGS ?= -v -count 1
GO_TEST_RUN ?= .

.PHONY: test
test:
	${GOLANG} test ${GO_TEST_FLAGS} -run ${GO_TEST_RUN} ./...

listener: ${GO_FILES}
	${GOLANG} build -o ./listener ./cmd/listener

replay: ${GO_FILES}
	${GOLANG} build -o ./replay ./cmd/replay

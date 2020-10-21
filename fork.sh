#!/usr/bin/env sh

set -e

cd "$(dirname "$0")"

getLog() {
  git log --abbrev-commit --max-count=1 --date=iso | head -n4
}

getAuthor() {
  echo $(getLog) | sed -e 's|.*Author: \(.*\) Date: \(.*\)|\1 \2|'
}

getVersion() {
  git describe --tags --long --always --dirty
}

ls | grep -v fork.sh | xargs -I {} rm -rf {}

git clone --depth=1 https://github.com/galaxy-book/feishu-sdk-golang.git tmp

cd tmp
sourceVersion=$(getVersion)
sourceAuthor=$(getAuthor)

find . -name '*_test.go' \
    -or -name '*.sh' \
    -or -name '*.bat' \
    -or -name '*.yml' \
    -or -name '.*.yml' \
    -or -name '.*.toml' \
    -or -name '*.toml' \
    -or -name 'go.sum' \
    -or -name 'Makefile' \
    -or -name 'Guardfile' \
    -or -name 'CONTRIBUT*' \
    -or -name 'CONTRIBUT*' \
    -or -name 'Gopkg*' \
    -or -name 'tests' \
    -or -name '_example' \
    -or -name '_examples' \
    -or -name 'examples' \
    -or -name 'benchmark' \
    -or -name 'benchmarks' \
    -or -name 'PULL_REQUEST_TEMPLATE.md' \
    -or -name '*_tests' \
    -or -name 'testdata' \
    -or -name 'travis' \
    -or -name '.travis' \
    -or -name '.git' \
    -or -name '.github' \
    -or -name '.circleci' | xargs -I {} rm -rf {}

find . -name '*.go' | while read -r f; do
  sed -i.bak 's?github.com/galaxy-book/feishu-sdk-golang?amzcrm/crawler/sms/feishu?g;s?package sdk?package feishu?g;s?fmt.Println(1111, reqBody)??g' "${f}"
  rm "${f}.bak"
done

mv sdk/* .
rm -rf sdk

sed -i.bak '/fmt.Println(1111, reqBody)/d' "robot_send_msg.go"
sed -i.bak '/"fmt"/d' "robot_send_msg.go"
rm -f "robot_send_msg.go.bak"

cat >core/util/log/log.go <<EOF
package log

type logger interface {
	Error(msg ...interface{})
	Errorf(msg string, args ...interface{})
	Info(msg ...interface{})
	Infof(msg string, args ...interface{})
}

var log logger

func Init(l logger) {
	log = l
}

func Error(msg ...interface{}) {
	if log != nil {
		log.Error(msg...)
	}
}

func ErrorF(msg string, args ...interface{}) {
	if log != nil {
		log.Errorf(msg, args...)
	}
}

func Info(msg ...interface{}) {
	if log != nil {
		log.Info(msg...)
	}
}

func InfoF(msg string, args ...interface{}) {
	if log != nil {
		log.Infof(msg, args...)
	}
}
EOF

find . -name go.mod -or -name go.sum | xargs -I {} rm -f {}
#go mod init shu.run/gopkg/sdk/feishu
#go mod tidy

cd ../
ls tmp | xargs -I {} mv tmp/{} .
rm -rf tmp

cat >fork.txt <<EOF
fork time: 
	$(date -R)
fork from:
	https://github.com/galaxy-book/feishu-sdk-golang.git
fork version:
    ${sourceVersion} ${sourceAuthor}
EOF

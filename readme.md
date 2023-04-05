## 개요

Testcontainers for Go is a Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done.

## 사용 용도

- 통합테스트에 사용하기 가장 적합하다
- E2E 테스트

## 레퍼런스

패키지: https://pkg.go.dev/github.com/testcontainers/testcontainers-go

## 사용방법

go test 프레임워크와 병행해서 사용할 수 있다.
도커를 통해서 테스트가 이루어진다

## 세팅

1. pkg 설치

```
go get github.com/testcontainers/testcontainers-go
```

2. 테스트 파일 작성

3. go test ./... -v 명령어 실행(자세히 볼수있음)

```
ok      testcont        9.674s --1차
ok      testcont        3.800s --2차
```

컨테이너 띄우고 내리는데 한참 걸리는듯(이미지 없는 경우엔)

package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",       //컨테이너가 사용하는 도커이미지
		ExposedPorts: []string{"6379/tcp"}, //컨테이너가 쓰는 포트. t.Parallel 넣으면 여러개의 랜덤 포트 받아서 테스트 가능
		WaitingFor:   wait.ForLog("Ready to accept connections"),
		//컨테이너가 테스트 요청 트래픽을 받을 준비가 언제 되는지 알려주는 역할을 하므로,
		//꼭 이 값을 넣어줘야 한다.
		//위 레디스 예시에서는 이 로그가 찍힌 시점 이후에 테스트 준비가 되었다는 것을 알려준다.
	}

	//컨테이너를 생성한다
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		//true는 컨테이너 함수가 컨테이너가 준비 후 실행되기까지 기다려준다는 뜻
		//false는 언제 시작할지 사용자가 결정해줘야 함
	})
	if err != nil {
		t.Error(err)
	}

	//테스트 하고 싶은 내용: 레디스 엔드포인트가 얻어지는지, 클라이언트 생성되는지 테스트.
	endpoint, err := redisC.Endpoint(ctx, "")
	fmt.Println(endpoint, "엔드포인트 찍히나")
	if err != nil {
		t.Error(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	_ = client

	//테스트 실행 후엔 더 이상 쓰지 않으니까 무조건 컨테이너를 삭제해야 한다. 따라서 아래의 코드 실행한다
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()
}

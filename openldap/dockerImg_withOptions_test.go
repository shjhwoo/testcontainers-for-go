package openldap

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

//https://medium.com/@philippe_andreas/how-to-customized-ldap-schema-docker-image-for-a-symfony-4-project-df6efc806867

func TestWithImageOptions(t *testing.T) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "osixia/openldap:1.3.0",
		Hostname:     "sf4app",
		ExposedPorts: []string{"389/ldap"},
		Cmd:          []string{"--copy-service"},
		WaitingFor:   wait.ForListeningPort(nat.Port("389/ldap")),
		Env: map[string]string{
			"LDAP_ADMIN_PASSWORD": "myadminpasswd",
			"LDAP_DOMAIN":         "sf4app.org",
		},
		Files: []testcontainers.ContainerFile{
			{HostFilePath: "./volumes/ldap/test.schema", ContainerFilePath: "/container/service/slapd/assets/config/bootstrap/schema/test.schema", FileMode: 700},
			{HostFilePath: "./volumes/ldap/test.ldif", ContainerFilePath: "/container/service/slapd/assets/config/bootstrap/ldif/test.ldif", FileMode: 700},
		},
	}

	//컨테이너를 생성한다
	openLDAPCon, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	fmt.Println(err, "에러 내용도 안 찍히나ㅠㅠ")
	if err != nil {
		t.Error(err)
	}

	//테스트 하고 싶은 내용: 레디스 엔드포인트가 얻어지는지, 클라이언트 생성되는지 테스트.
	endpoint, err := openLDAPCon.Endpoint(ctx, "")
	fmt.Println(endpoint, "엔드포인트 찍히나")
	if err != nil {
		t.Error(err)
	}

	//테스트 실행 후엔 더 이상 쓰지 않으니까 무조건 컨테이너를 삭제해야 한다. 따라서 아래의 코드 실행한다
	defer func() {
		if err := openLDAPCon.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()
}

package sandboxFactory_test

import (
	"nginx_debugger/internal/sandboxFactory"
	"testing"
	"time"
)

var configuration = `

server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;

    #access_log  /var/log/nginx/host.access.log  main;

    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}

`

func TestNewSandboxFactory(t *testing.T) {
	factory, err := sandboxFactory.NewSandboxFactory()
	if err != nil {
		t.Fatal(err)
	}

	sandbox, err := factory.CreateSandbox()
	if err != nil {
		t.Fatal(err)
	}

	if err = factory.InitializeDocker(); err != nil {
		t.Fatal(err)
	}

	if err = sandbox.Start(); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Second * 5)

	if err = sandbox.StopAndDestroy(); err != nil {
		t.Fatal(err)
	}
}

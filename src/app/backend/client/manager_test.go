package client

/**
 * Using minikube/kubeadm start a single node k8s cluster in the remote host
 * and then using `kubectl proxy --port=8080 &` to proxy apiserver at port 8080 with protocol http(not https)
 * In the local dev mode, we use ssh's Local Port Forwarding to make a SSH Tunnel between localhost:8080 and
 * remote host's port 8080
 */
import (
	"github.com/emicklei/go-restful"
	"net/http"
	"testing"
)

func TestNewClientManager(t *testing.T) {
	cases := []struct {
		kubeConfigPath, apiserverHost string
	}{
		{"", "test"},
	}

	for _, c := range cases {
		manager := NewClientManager(c.kubeConfigPath, c.apiserverHost)

		if manager == nil {
			t.Fatalf("NewClientManager(%s, %s): Expected manager not to be nil",
				c.kubeConfigPath, c.apiserverHost)
		}
	}
}

func TestClient(t *testing.T) {
	cases := []struct {
		request *restful.Request
	}{
		{
			&restful.Request{
				Request: &http.Request{
					Header: http.Header(map[string][]string{}),
				},
			},
		},
		{nil},
	}

	for _, c := range cases {
		manager := NewClientManager("", "http://localhost:8080")
		_, err := manager.Client(c.request)

		if err != nil {
			t.Fatalf("Client(%v): Expected client to be created but error was thrown:"+
				" %s", c.request, err.Error())
		}
	}
}

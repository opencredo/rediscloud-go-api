package rediscloud_api

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/RedisLabs/rediscloud-go-api/redis"
	"github.com/RedisLabs/rediscloud-go-api/service/cloud_accounts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudAccount_Create(t *testing.T) {
	expected := 1235
	s := httptest.NewServer(testServer("key", "secret", postRequest(t, "/cloud-accounts", `{
  "accessKeyId": "123456",
  "accessSecretKey": "765432",
  "consoleUsername": "foo",
  "consolePassword": "bar",
  "name": "cumulus nimbus",
  "provider": "AWS",
  "signInLoginUrl": "http://example.org/foo"
}`, `{
  "taskId": "task-id",
  "commandType": "cloudAccountCreateRequest",
  "status": "received",
  "description": "Task request received and is being queued for processing.",
  "timestamp": "2020-11-02T09:05:34.3Z",
  "_links": {
    "task": {
      "href": "https://example.org",
      "title": "getTaskStatusUpdates",
      "type": "GET"
    }
  }
}`), getRequest(t, "/tasks/task-id", fmt.Sprintf(`{
  "taskId": "e02b40d6-1395-4861-a3b9-ecf829d835fd",
  "commandType": "subscriptionCreateRequest",
  "status": "processing-completed",
  "timestamp": "2020-10-28T09:58:16.798Z",
  "response": {
    "resourceId": %d
  },
  "_links": {
    "self": {
      "href": "https://example.com",
      "type": "GET"
    }
  }
}`, expected))))

	subject, err := clientFromTestServer(s, "key", "secret")
	require.NoError(t, err)

	actual, err := subject.CloudAccount.Create(context.TODO(), cloud_accounts.CreateCloudAccount{
		AccessKeyID:     redis.String("123456"),
		AccessSecretKey: redis.String("765432"),
		ConsoleUsername: redis.String("foo"),
		ConsolePassword: redis.String("bar"),
		Name:            redis.String("cumulus nimbus"),
		Provider:        redis.String("AWS"),
		SignInLoginURL:  redis.String("http://example.org/foo"),
	})
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCloudAccount_Update(t *testing.T) {
	s := httptest.NewServer(testServer("key", "secret", putRequest(t, "/cloud-accounts/642", `{
  "accessKeyId": "tfvbjuyg",
  "accessSecretKey": "gyujmnbvgy",
  "consoleUsername": "baz",
  "consolePassword": "bar",
  "name": "stratocumulus",
  "signInLoginUrl": "http://example.org/foo"
}`, `{
  "taskId": "task-id",
  "commandType": "cloudAccountUpdateRequest",
  "status": "received",
  "description": "Task request received and is being queued for processing.",
  "timestamp": "2020-11-02T09:05:34.3Z",
  "_links": {
    "task": {
      "href": "https://example.org",
      "title": "getTaskStatusUpdates",
      "type": "GET"
    }
  }
}`), getRequest(t, "/tasks/task-id", `{
  "taskId": "e02b40d6-1395-4861-a3b9-ecf829d835fd",
  "commandType": "cloudAccountUpdateRequest",
  "status": "processing-completed",
  "timestamp": "2020-10-28T09:58:16.798Z",
  "response": {
  },
  "_links": {
    "self": {
      "href": "https://example.com",
      "type": "GET"
    }
  }
}`)))

	subject, err := clientFromTestServer(s, "key", "secret")
	require.NoError(t, err)

	err = subject.CloudAccount.Update(context.TODO(), 642, cloud_accounts.UpdateCloudAccount{
		AccessKeyID:     redis.String("tfvbjuyg"),
		AccessSecretKey: redis.String("gyujmnbvgy"),
		ConsoleUsername: redis.String("baz"),
		ConsolePassword: redis.String("bar"),
		Name:            redis.String("stratocumulus"),
		SignInLoginURL:  redis.String("http://example.org/foo"),
	})
	require.NoError(t, err)
}

func TestCloudAccount_Get(t *testing.T) {
	s := httptest.NewServer(testServer("apiKey", "secret", getRequest(t, "/cloud-accounts/98765", `{
  "id": 97643,
  "name": "Frank",
  "provider": "GCP",
  "status": "active",
  "accessKeyId": "keyId",
  "_links": {
    "self": {
      "href": "https://example.org",
      "type": "GET"
    }
  }
}`)))

	subject, err := clientFromTestServer(s, "apiKey", "secret")
	require.NoError(t, err)

	actual, err := subject.CloudAccount.Get(context.TODO(), 98765)
	require.NoError(t, err)

	assert.Equal(t, &cloud_accounts.CloudAccount{
		Name:        redis.String("Frank"),
		Provider:    redis.String("GCP"),
		Status:      redis.String("active"),
		AccessKeyID: redis.String("keyId"),
	}, actual)
}

func TestCloudAccount_Delete(t *testing.T) {
	s := httptest.NewServer(testServer("apiKey", "secret", deleteRequest(t, "/cloud-accounts/98765", `{
  "taskId": "task",
  "commandType": "cloudAccountDeleteRequest",
  "status": "received",
  "description": "Task request received and is being queued for processing.",
  "timestamp": "2020-11-02T09:05:34.3Z",
  "_links": {
    "task": {
      "href": "https://example.org",
      "title": "getTaskStatusUpdates",
      "type": "GET"
    }
  }
}`), getRequest(t, "/tasks/task", `{
  "taskId": "e02b40d6-1395-4861-a3b9-ecf829d835fd",
  "commandType": "cloudAccountDeleteRequest",
  "status": "processing-completed",
  "timestamp": "2020-10-28T09:58:16.798Z",
  "response": {
  },
  "_links": {
    "self": {
      "href": "https://example.com",
      "type": "GET"
    }
  }
}`)))

	subject, err := clientFromTestServer(s, "apiKey", "secret")
	require.NoError(t, err)

	err = subject.CloudAccount.Delete(context.TODO(), 98765)
	require.NoError(t, err)
}
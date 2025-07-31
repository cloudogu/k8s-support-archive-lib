package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"k8s.io/apimachinery/pkg/api/meta"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	v1 "github.com/cloudogu/k8s-support-archive-lib/api/v1"
)

var testCtx = context.Background()

func Test_supportArchiveClient_Get(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/testsupportArchive", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)

			writer.Header().Add("content-type", "application/json")
			supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "testsupportArchive", Namespace: "test"}}
			supportArchiveBytes, err := json.Marshal(supportArchive)
			require.NoError(t, err)
			_, err = writer.Write(supportArchiveBytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.Get(testCtx, "testsupportArchive", metav1.GetOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_List(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodGet, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)

			writer.Header().Add("content-type", "application/json")
			supportArchiveList := v1.SupportArchiveList{}
			supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "testsupportArchive", Namespace: "test"}}
			supportArchiveList.Items = append(supportArchiveList.Items, *supportArchive)
			supportArchiveBytes, err := json.Marshal(supportArchiveList)
			require.NoError(t, err)
			_, err = writer.Write(supportArchiveBytes)
			require.NoError(t, err)
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		timeout := int64(5)

		// when
		_, err = sClient.List(testCtx, metav1.ListOptions{TimeoutSeconds: &timeout})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPost, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "tocreate", createdSupportArchive.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.Create(testCtx, supportArchive, metav1.CreateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/tocreate", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "tocreate", createdSupportArchive.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.Update(testCtx, supportArchive, metav1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_UpdateStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "tocreate", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/tocreate/status", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "tocreate", createdSupportArchive.Name)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.UpdateStatus(testCtx, supportArchive, metav1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/testsupportArchive", request.URL.Path)

			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		err = sClient.Delete(testCtx, "testsupportArchive", metav1.DeleteOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_DeleteCollection(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodDelete, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives", request.URL.Path)
			assert.Equal(t, "labelSelector=test&timeout=5s&timeoutSeconds=5", request.URL.RawQuery)
			writer.Header().Add("content-type", "application/json")
			writer.WriteHeader(200)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")
		timeout := int64(5)

		// when
		err = sClient.DeleteCollection(testCtx, metav1.DeleteOptions{}, metav1.ListOptions{LabelSelector: "test", TimeoutSeconds: &timeout})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_Patch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPatch, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/testsupportArchive", request.URL.Path)
			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)
			assert.Equal(t, []byte("test"), bytes)
			result, err := json.Marshal(v1.SupportArchive{})
			require.NoError(t, err)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(result)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		patchData := []byte("test")

		// when
		_, err = sClient.Patch(testCtx, "testsupportArchive", types.JSONPatchType, patchData, metav1.PatchOptions{})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_Watch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "GET", request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives", request.URL.Path)
			assert.Equal(t, http.NoBody, request.Body)
			assert.Equal(t, "labelSelector=test&timeout=5s&timeoutSeconds=5&watch=true", request.URL.RawQuery)

			writer.Header().Add("content-type", "application/json")
			_, err := writer.Write([]byte("egal"))
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		timeout := int64(5)

		// when
		_, err = sClient.Watch(testCtx, metav1.ListOptions{LabelSelector: "test", TimeoutSeconds: &timeout})

		// then
		require.NoError(t, err)
	})
}

func Test_supportArchiveClient_AddFinalizer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "mySupportArchive", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/mySupportArchive", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "mySupportArchive", createdSupportArchive.Name)
			assert.Len(t, createdSupportArchive.Finalizers, 1)
			assert.Equal(t, "myFinalizer", createdSupportArchive.Finalizers[0])

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.AddFinalizer(testCtx, supportArchive, "myFinalizer")

		// then
		require.NoError(t, err)
	})

	t.Run("should fail to set finalizer on client error", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "mySupportArchive", Namespace: "test"}}

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/mySupportArchive", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "mySupportArchive", createdSupportArchive.Name)
			assert.Len(t, createdSupportArchive.Finalizers, 1)
			assert.Equal(t, "myFinalizer", createdSupportArchive.Finalizers[0])

			writer.WriteHeader(500)
			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.AddFinalizer(testCtx, supportArchive, "myFinalizer")

		// then
		require.Error(t, err)
		assert.ErrorContains(t, err, "failed to add finalizer myFinalizer to supportArchive:")
	})
}

func Test_supportArchiveClient_RemoveFinalizer(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "mySupportArchive", Namespace: "test"}}
		controllerutil.AddFinalizer(supportArchive, "finalizer1")
		controllerutil.AddFinalizer(supportArchive, "finalizer2")

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/mySupportArchive", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "mySupportArchive", createdSupportArchive.Name)
			assert.Len(t, createdSupportArchive.Finalizers, 1)
			assert.Equal(t, "finalizer2", createdSupportArchive.Finalizers[0])

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.RemoveFinalizer(testCtx, supportArchive, "finalizer1")

		// then
		require.NoError(t, err)
	})

	t.Run("should fail to set finalizer on client error", func(t *testing.T) {
		// given
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "mySupportArchive", Namespace: "test"}}
		controllerutil.AddFinalizer(supportArchive, "finalizer1")
		controllerutil.AddFinalizer(supportArchive, "finalizer2")

		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, http.MethodPut, request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/mySupportArchive", request.URL.Path)

			bytes, err := io.ReadAll(request.Body)
			require.NoError(t, err)

			createdSupportArchive := &v1.SupportArchive{}
			require.NoError(t, json.Unmarshal(bytes, createdSupportArchive))
			assert.Equal(t, "mySupportArchive", createdSupportArchive.Name)
			assert.Len(t, createdSupportArchive.Finalizers, 1)
			assert.Equal(t, "finalizer1", createdSupportArchive.Finalizers[0])

			writer.WriteHeader(500)
			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(bytes)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.RemoveFinalizer(testCtx, supportArchive, "finalizer2")

		// then
		require.Error(t, err)
		assert.ErrorContains(t, err, "failed to remove finalizer finalizer2 from supportArchive")
	})
}

func Test_supportArchiveClient_UpdateStatusWithRetry(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// given
		downloadURL := "url"
		supportArchive := &v1.SupportArchive{ObjectMeta: metav1.ObjectMeta{Name: "mySupportArchive", Namespace: "test"}}
		modifyFunc := func(status v1.SupportArchiveStatus) v1.SupportArchiveStatus {
			meta.SetStatusCondition(&status.Conditions, metav1.Condition{
				Type:               v1.ConditionSupportArchiveCreated,
				Status:             metav1.ConditionTrue,
				LastTransitionTime: metav1.Time{Time: time.Now()},
				Reason:             "AllCollectorsExecuted",
				Message:            fmt.Sprintf("It is available for download under following url: %s", downloadURL),
			})
			status.DownloadPath = downloadURL
			return status
		}
		server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			assert.Equal(t, "PUT", request.Method)
			assert.Equal(t, "/apis/k8s.cloudogu.com/v1/namespaces/test/supportarchives/mySupportArchive/status", request.URL.Path)

			result, err := json.Marshal(v1.SupportArchiveStatus{})
			require.NoError(t, err)

			writer.Header().Add("content-type", "application/json")
			_, err = writer.Write(result)
			require.NoError(t, err)
		}))

		config := rest.Config{
			Host: server.URL,
		}
		client, err := NewForConfig(&config)
		require.NoError(t, err)
		sClient := client.SupportArchives("test")

		// when
		_, err = sClient.UpdateStatusWithRetry(testCtx, supportArchive, modifyFunc, metav1.UpdateOptions{})

		// then
		require.NoError(t, err)
	})
}

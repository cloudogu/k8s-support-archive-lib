package v1

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	v1 "github.com/cloudogu/k8s-support-archive-lib/api/v1"
	"github.com/cloudogu/retry-lib/retry"
)

type supportArchiveClient struct {
	client rest.Interface
	ns     string
}

// UpdateStatusCreating sets the status of the supportArchive to "creating".
func (client *supportArchiveClient) UpdateStatusCreating(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error) {
	return client.updateStatusPhaseWithRetry(ctx, supportArchive, v1.StatusPhaseCreating)
}

// UpdateStatusCreated sets the status of the supportArchive to "created".
func (client *supportArchiveClient) UpdateStatusCreated(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error) {
	return client.updateStatusPhaseWithRetry(ctx, supportArchive, v1.StatusPhaseCreated)
}

// UpdateStatusDeleting sets the status of the supportArchive to "deleting".
func (client *supportArchiveClient) UpdateStatusDeleting(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error) {
	return client.updateStatusPhaseWithRetry(ctx, supportArchive, v1.StatusPhaseDeleting)
}

// UpdateStatusFailed sets the status of the supportArchive to "failed".
func (client *supportArchiveClient) UpdateStatusFailed(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error) {
	return client.updateStatusPhaseWithRetry(ctx, supportArchive, v1.StatusPhaseFailed)
}

func (client *supportArchiveClient) updateStatusPhaseWithRetry(ctx context.Context, supportArchive *v1.SupportArchive, targetStatus v1.StatusPhase) (*v1.SupportArchive, error) {
	var resultSupportArchive *v1.SupportArchive
	err := retry.OnConflict(func() error {
		updatedSupportArchive, err := client.Get(ctx, supportArchive.GetName(), metav1.GetOptions{})
		if err != nil {
			return err
		}

		// do not overwrite the whole status, so we do not lose other values from the Status object
		// esp. a potentially set requeue time
		updatedSupportArchive.Status.Phase = targetStatus
		resultSupportArchive, err = client.UpdateStatus(ctx, updatedSupportArchive, metav1.UpdateOptions{})
		return err
	})

	return resultSupportArchive, err
}

// UpdateStatusWithRetry updates the status of the resource, retrying if a conflict error arises.
func (client *supportArchiveClient) UpdateStatusWithRetry(ctx context.Context, cr *v1.SupportArchive, modifyStatusFn func(v1.SupportArchiveStatus) v1.SupportArchiveStatus, opts metav1.UpdateOptions) (result *v1.SupportArchive, err error) {
	firstTry := true

	var currentObj *v1.SupportArchive
	err = retry.OnConflict(func() error {
		if firstTry {
			firstTry = false
			currentObj = cr.DeepCopy()
		} else {
			currentObj, err = client.Get(ctx, cr.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
		}

		currentObj.Status = modifyStatusFn(currentObj.Status)
		currentObj, err = client.UpdateStatus(ctx, currentObj, opts)
		return err
	})
	if err != nil {
		return nil, err
	}

	return currentObj, nil
}

// AddFinalizer adds the given finalizer to the supportArchive.
func (client *supportArchiveClient) AddFinalizer(ctx context.Context, supportArchive *v1.SupportArchive, finalizer string) (*v1.SupportArchive, error) {
	controllerutil.AddFinalizer(supportArchive, finalizer)
	result, err := client.Update(ctx, supportArchive, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to add finalizer %s to supportArchive: %w", finalizer, err)
	}

	return result, nil
}

// RemoveFinalizer removes the given finalizer to the supportArchive.
func (client *supportArchiveClient) RemoveFinalizer(ctx context.Context, supportArchive *v1.SupportArchive, finalizer string) (*v1.SupportArchive, error) {
	controllerutil.RemoveFinalizer(supportArchive, finalizer)
	result, err := client.Update(ctx, supportArchive, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to remove finalizer %s from supportArchive: %w", finalizer, err)
	}

	return result, err
}

// Get takes name of the supportArchive, and returns the corresponding supportArchive object, and an error if there is any.
func (client *supportArchiveClient) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SupportArchive, err error) {
	result = &v1.SupportArchive{}
	err = client.client.Get().
		Namespace(client.ns).
		Resource("supportArchives").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of supportArchives that match those selectors.
func (client *supportArchiveClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SupportArchiveList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SupportArchiveList{}
	err = client.client.Get().
		Namespace(client.ns).
		Resource("supportArchives").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested supportArchives.
func (client *supportArchiveClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return client.client.Get().
		Namespace(client.ns).
		Resource("supportArchives").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a supportArchive and creates it.  Returns the server's representation of the supportArchive, and an error, if there is any.
func (client *supportArchiveClient) Create(ctx context.Context, supportArchive *v1.SupportArchive, opts metav1.CreateOptions) (result *v1.SupportArchive, err error) {
	result = &v1.SupportArchive{}
	err = client.client.Post().
		Namespace(client.ns).
		Resource("supportArchives").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(supportArchive).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a supportArchive and updates it. Returns the server's representation of the supportArchive, and an error, if there is any.
func (client *supportArchiveClient) Update(ctx context.Context, supportArchive *v1.SupportArchive, opts metav1.UpdateOptions) (result *v1.SupportArchive, err error) {
	result = &v1.SupportArchive{}
	err = client.client.Put().
		Namespace(client.ns).
		Resource("supportArchives").
		Name(supportArchive.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(supportArchive).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (client *supportArchiveClient) UpdateStatus(ctx context.Context, supportArchive *v1.SupportArchive, opts metav1.UpdateOptions) (result *v1.SupportArchive, err error) {
	result = &v1.SupportArchive{}
	err = client.client.Put().
		Namespace(client.ns).
		Resource("supportArchives").
		Name(supportArchive.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(supportArchive).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the supportArchive and deletes it. Returns an error if one occurs.
func (client *supportArchiveClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return client.client.Delete().
		Namespace(client.ns).
		Resource("supportArchives").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (client *supportArchiveClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return client.client.Delete().
		Namespace(client.ns).
		Resource("supportArchives").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched supportArchive.
func (client *supportArchiveClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SupportArchive, err error) {
	result = &v1.SupportArchive{}
	err = client.client.Patch(pt).
		Namespace(client.ns).
		Resource("supportArchives").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

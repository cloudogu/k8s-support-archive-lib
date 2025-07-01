package v1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"

	v1 "github.com/cloudogu/k8s-support-archive-lib/api/v1"
)

type SupportArchiveV1Interface interface {
	SupportArchives(namespace string) SupportArchiveInterface
}

type SupportArchiveInterface interface {
	// Create takes the representation of a supportArchive and creates it.  Returns the server's representation of the supportArchive, and an error, if there is any.
	Create(ctx context.Context, supportArchive *v1.SupportArchive, opts metav1.CreateOptions) (*v1.SupportArchive, error)
	// Update takes the representation of a supportArchive and updates it. Returns the server's representation of the supportArchive, and an error, if there is any.
	Update(ctx context.Context, supportArchive *v1.SupportArchive, opts metav1.UpdateOptions) (*v1.SupportArchive, error)
	// UpdateStatus was generated because the type contains a Status member.
	UpdateStatus(ctx context.Context, supportArchive *v1.SupportArchive, opts metav1.UpdateOptions) (*v1.SupportArchive, error)
	// UpdateStatusWithRetry updates the status according to modifyStatusFn and if a conflict error occurs, the method will refetch the resource and retry the status update.
	UpdateStatusWithRetry(ctx context.Context, cr *v1.SupportArchive, modifyStatusFn func(v1.SupportArchiveStatus) v1.SupportArchiveStatus, opts metav1.UpdateOptions) (*v1.SupportArchive, error)
	// UpdateStatusCreating sets the status of the supportArchive to "creating".
	UpdateStatusCreating(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error)
	// UpdateStatusCreated sets the status of the supportArchive to "created".
	UpdateStatusCreated(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error)
	// UpdateStatusDeleting sets the status of the supportArchive to "deleting".
	UpdateStatusDeleting(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error)
	// UpdateStatusFailed sets the status of the supportArchive to "failed".
	UpdateStatusFailed(ctx context.Context, supportArchive *v1.SupportArchive) (*v1.SupportArchive, error)
	// Delete takes name of the supportArchive and deletes it. Returns an error if one occurs.
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	// DeleteCollection deletes a collection of objects.
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	// Get takes name of the supportArchive, and returns the corresponding supportArchive object, and an error if there is any.
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SupportArchive, error)
	// List takes label and field selectors, and returns the list of supportArchives that match those selectors.
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SupportArchiveList, error)
	// Watch returns a watch.Interface that watches the requested supportArchives.
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	// Patch applies the patch and returns the patched supportArchive.
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SupportArchive, err error)
	// AddFinalizer adds the given finalizer to the supportArchive.
	AddFinalizer(ctx context.Context, supportArchive *v1.SupportArchive, finalizer string) (*v1.SupportArchive, error)
	// RemoveFinalizer removes the given finalizer to the supportArchive.
	RemoveFinalizer(ctx context.Context, supportArchive *v1.SupportArchive, finalizer string) (*v1.SupportArchive, error)
}

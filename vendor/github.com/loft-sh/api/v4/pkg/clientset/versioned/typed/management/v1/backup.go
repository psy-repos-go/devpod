// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BackupsGetter has a method to return a BackupInterface.
// A group's client should implement this interface.
type BackupsGetter interface {
	Backups() BackupInterface
}

// BackupInterface has methods to work with Backup resources.
type BackupInterface interface {
	Create(ctx context.Context, backup *v1.Backup, opts metav1.CreateOptions) (*v1.Backup, error)
	Update(ctx context.Context, backup *v1.Backup, opts metav1.UpdateOptions) (*v1.Backup, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, backup *v1.Backup, opts metav1.UpdateOptions) (*v1.Backup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Backup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BackupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Backup, err error)
	BackupExpansion
}

// backups implements BackupInterface
type backups struct {
	*gentype.ClientWithList[*v1.Backup, *v1.BackupList]
}

// newBackups returns a Backups
func newBackups(c *ManagementV1Client) *backups {
	return &backups{
		gentype.NewClientWithList[*v1.Backup, *v1.BackupList](
			"backups",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.Backup { return &v1.Backup{} },
			func() *v1.BackupList { return &v1.BackupList{} }),
	}
}

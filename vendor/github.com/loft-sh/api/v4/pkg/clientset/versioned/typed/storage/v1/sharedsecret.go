// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/loft-sh/api/v4/pkg/apis/storage/v1"
	scheme "github.com/loft-sh/api/v4/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// SharedSecretsGetter has a method to return a SharedSecretInterface.
// A group's client should implement this interface.
type SharedSecretsGetter interface {
	SharedSecrets(namespace string) SharedSecretInterface
}

// SharedSecretInterface has methods to work with SharedSecret resources.
type SharedSecretInterface interface {
	Create(ctx context.Context, sharedSecret *v1.SharedSecret, opts metav1.CreateOptions) (*v1.SharedSecret, error)
	Update(ctx context.Context, sharedSecret *v1.SharedSecret, opts metav1.UpdateOptions) (*v1.SharedSecret, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, sharedSecret *v1.SharedSecret, opts metav1.UpdateOptions) (*v1.SharedSecret, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SharedSecret, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SharedSecretList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SharedSecret, err error)
	SharedSecretExpansion
}

// sharedSecrets implements SharedSecretInterface
type sharedSecrets struct {
	*gentype.ClientWithList[*v1.SharedSecret, *v1.SharedSecretList]
}

// newSharedSecrets returns a SharedSecrets
func newSharedSecrets(c *StorageV1Client, namespace string) *sharedSecrets {
	return &sharedSecrets{
		gentype.NewClientWithList[*v1.SharedSecret, *v1.SharedSecretList](
			"sharedsecrets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.SharedSecret { return &v1.SharedSecret{} },
			func() *v1.SharedSecretList { return &v1.SharedSecretList{} }),
	}
}

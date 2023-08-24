// Code generated by ifacemaker. DO NOT EDIT.

package v1

import (
	"context"

	mpc_keyspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"github.com/googleapis/gax-go/v2"
)

// ClientWrappedAddDeviceOperation is a generated interface for mpckey protoc-gen-go_gapic client WrappedAddDeviceOperation.
type ClientWrappedAddDeviceOperation interface {
	// PathPrefix returns the path prefix for the operation to be used in HTTP requests.
	// E.g. if the path prefix is `/waas/mpc_keys`, then the request path would be:
	// `/waas/mpc_keys/v1/operations/<operationId>`.
	PathPrefix() string
	// Wait delegates to the wrapped longrunning AddDeviceOperation with the
	// override LRO options and handling unwrapping client errors.
	Wait(ctx context.Context, opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error)
	// Poll delegates to the wrapped longrunning AddDeviceOperation with the
	// override LRO options and handling unwrapping client errors.
	Poll(ctx context.Context, opts ...gax.CallOption) (*mpc_keyspb.DeviceGroup, error)
	// Metadata delegates to the wrapped longrunning AddDeviceOperation.
	Metadata() (*mpc_keyspb.AddDeviceMetadata, error)
	// Done delegates to the wrapped longrunning AddDeviceOperation.
	Done() bool
	// Name delegates to the wrapped longrunning AddDeviceOperation.
	Name() string
}

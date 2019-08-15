package obc

import (
	"github.com/noobaa/noobaa-operator/pkg/obc"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Add creates a Controller and adds it to the Manager.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {

	// run the noobaa bucket provisioner:
	err := obc.RunProvisioner(
		mgr.GetClient(),
		mgr.GetScheme(),
		mgr.GetRecorder("noobaa-operator"),
	)
	if err != nil {
		return err
	}

	return nil
}
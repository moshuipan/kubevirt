// Code generated by swagger-doc. DO NOT EDIT.

package v1

func (VirtualMachineInstance) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineInstance is *the* VirtualMachineInstance Definition. It represents a virtual machine in the runtime environment of kubernetes.",
		"spec":   "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
		"status": "Status is the high level overview of how the VirtualMachineInstance is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineInstanceList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceList is a list of VirtualMachines",
	}
}

func (VirtualMachineInstanceSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                              "VirtualMachineInstanceSpec is a description of a VirtualMachineInstance.",
		"domain":                        "Specification of the desired behavior of the VirtualMachineInstance on the host.",
		"nodeSelector":                  "NodeSelector is a selector which must be true for the vmi to fit on a node.\nSelector which must match a node's labels for the vmi to be scheduled on that node.\nMore info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/\n+optional",
		"affinity":                      "If affinity is specifies, obey all the affinity rules",
		"tolerations":                   "If toleration is specified, obey all the toleration rules.",
		"terminationGracePeriodSeconds": "Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated.",
		"volumes":                       "List of volumes that can be mounted by disks belonging to the vmi.",
		"livenessProbe":                 "Periodic probe of VirtualMachineInstance liveness.\nVirtualmachineInstances will be stopped if the probe fails.\nCannot be updated.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"readinessProbe":                "Periodic probe of VirtualMachineInstance service readiness.\nVirtualmachineInstances will be removed from service endpoints if the probe fails.\nCannot be updated.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"hostname":                      "Specifies the hostname of the vmi\nIf not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly.\n+optional",
		"subdomain":                     "If specified, the fully qualified vmi hostname will be \"<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>\".\nIf not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi,\nno matter if the vmi itself can pick up a hostname.\n+optional",
		"networks":                      "List of networks that can be attached to a vm's virtual interface.",
		"dnsPolicy":                     "Set DNS policy for the pod.\nDefaults to \"ClusterFirst\".\nValid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.\nDNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.\nTo have DNS options set along with hostNetwork, you have to specify DNS policy\nexplicitly to 'ClusterFirstWithHostNet'.\n+optional",
		"dnsConfig":                     "Specifies the DNS parameters of a pod.\nParameters specified here will be merged to the generated DNS\nconfiguration based on DNSPolicy.\n+optional",
	}
}

func (VirtualMachineInstanceStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "VirtualMachineInstanceStatus represents information about the status of a VirtualMachineInstance. Status may trail the actual\nstate of a system.",
		"nodeName":        "NodeName is the name where the VirtualMachineInstance is currently running.",
		"reason":          "A brief CamelCase message indicating details about why the VMI is in this state. e.g. 'NodeUnresponsive'\n+optional",
		"conditions":      "Conditions are specific points in VirtualMachineInstance's pod runtime.",
		"phase":           "Phase is the status of the VirtualMachineInstance in kubernetes world. It is not the VirtualMachineInstance status, but partially correlates to it.",
		"interfaces":      "Interfaces represent the details of available network interfaces.",
		"migrationState":  "Represents the status of a live migration",
		"migrationMethod": "Represents the method using which the vmi can be migrated: live migration or block migration",
	}
}

func (VirtualMachineInstanceCondition) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VirtualMachineInstanceNetworkInterface) SwaggerDoc() map[string]string {
	return map[string]string{
		"ipAddress":     "IP address of a Virtual Machine interface",
		"mac":           "Hardware address of a Virtual Machine interface",
		"name":          "Name of the interface, corresponds to name of the network assigned to the interface",
		"ipAddresses":   "List of all IP addresses of a Virtual Machine interface",
		"interfaceName": "The interface name inside the Virtual Machine",
	}
}

func (VirtualMachineInstanceMigrationState) SwaggerDoc() map[string]string {
	return map[string]string{
		"startTimestamp":                 "The time the migration action began",
		"endTimestamp":                   "The time the migration action ended",
		"targetNodeDomainDetected":       "The Target Node has seen the Domain Start Event",
		"targetNodeAddress":              "The address of the target node to use for the migration",
		"targetDirectMigrationNodePorts": "The list of ports opened for live migration on the destination node",
		"targetNode":                     "The target node that the VMI is moving to",
		"targetPod":                      "The target pod that the VMI is moving to",
		"sourceNode":                     "The source node that the VMI originated on",
		"completed":                      "Indicates the migration completed",
		"failed":                         "Indicates that the migration failed",
		"migrationUid":                   "The VirtualMachineInstanceMigration object associated with this migration",
	}
}

func (MigrationConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"completionTimeoutPerGiB": "The time for GiB of data to wait for the migration to be completed before aborting it",
		"progressTimeout":         "The time to wait for live migration to make progress in transferring data.",
	}
}

func (VMISelector) SwaggerDoc() map[string]string {
	return map[string]string{
		"name": "Name of the VirtualMachineInstance to migrate",
	}
}

func (VirtualMachineInstanceReplicaSet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineInstance is *the* VirtualMachineInstance Definition. It represents a virtual machine in the runtime environment of kubernetes.",
		"spec":   "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
		"status": "Status is the high level overview of how the VirtualMachineInstance is doing. It contains information available to controllers and users.",
	}
}

func (VirtualMachineInstanceReplicaSetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VMIList is a list of VMIs",
	}
}

func (VirtualMachineInstanceReplicaSetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"replicas": "Number of desired pods. This is a pointer to distinguish between explicit\nzero and not specified. Defaults to 1.\n+optional",
		"selector": "Label selector for pods. Existing ReplicaSets whose pods are\nselected by this will be the ones affected by this deployment.",
		"template": "Template describes the pods that will be created.",
		"paused":   "Indicates that the replica set is paused.\n+optional",
	}
}

func (VirtualMachineInstanceReplicaSetStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"replicas":      "Total number of non-terminated pods targeted by this deployment (their labels match the selector).\n+optional",
		"readyReplicas": "The number of ready replicas for this replica set.\n+optional",
		"labelSelector": "Canonical form of the label selector for HPA which consumes it through the scale subresource.",
	}
}

func (VirtualMachineInstanceReplicaSetCondition) SwaggerDoc() map[string]string {
	return map[string]string{}
}

func (VirtualMachineInstanceTemplateSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"spec": "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
	}
}

func (VirtualMachineInstanceMigration) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceMigration represents the object tracking a VMI's migration\nto another host in the cluster",
	}
}

func (VirtualMachineInstanceMigrationList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceMigrationList is a list of VirtualMachineMigrations",
	}
}

func (VirtualMachineInstanceMigrationSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"vmiName": "The name of the VMI to perform the migration on. VMI must exist in the migration objects namespace",
	}
}

func (VirtualMachineInstanceMigrationStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstanceMigration reprents information pertaining to a VMI's migration.",
	}
}

func (VirtualMachineInstancePreset) SwaggerDoc() map[string]string {
	return map[string]string{
		"spec": "VirtualMachineInstance Spec contains the VirtualMachineInstance specification.",
	}
}

func (VirtualMachineInstancePresetList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineInstancePresetList is a list of VirtualMachinePresets",
	}
}

func (VirtualMachineInstancePresetSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"selector": "Selector is a label query over a set of VMIs.\nRequired.",
		"domain":   "Domain is the same object type as contained in VirtualMachineInstanceSpec",
	}
}

func (VirtualMachine) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachine handles the VirtualMachines that are not running\nor are in a stopped state\nThe VirtualMachine contains the template to create the\nVirtualMachineInstance. It also mirrors the running state of the created\nVirtualMachineInstance in its status.",
		"spec":   "Spec contains the specification of VirtualMachineInstance created",
		"status": "Status holds the current state of the controller and brief information\nabout its associated VirtualMachineInstance",
	}
}

func (VirtualMachineList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachineList is a list of virtualmachines",
		"items": "Items is a list of VirtualMachines",
	}
}

func (VirtualMachineSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "VirtualMachineSpec describes how the proper VirtualMachine\nshould look like",
		"running":             "Running controls whether the associatied VirtualMachineInstance is created or not",
		"template":            "Template is the direct specification of VirtualMachineInstance",
		"dataVolumeTemplates": "dataVolumeTemplates is a list of dataVolumes that the VirtualMachineInstance template can reference.\nDataVolumes in this list are dynamically created for the VirtualMachine and are tied to the VirtualMachine's life-cycle.",
	}
}

func (VirtualMachineStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "VirtualMachineStatus represents the status returned by the\ncontroller to describe how the VirtualMachine is doing",
		"created":    "Created indicates if the virtual machine is created in the cluster",
		"ready":      "Ready indicates if the virtual machine is running and ready",
		"conditions": "Hold the state information of the VirtualMachine and its VirtualMachineInstance",
	}
}

func (VirtualMachineCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "VirtualMachineCondition represents the state of VirtualMachine",
	}
}

func (Handler) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "Handler defines a specific action that should be taken",
		"httpGet":   "HTTPGet specifies the http request to perform.\n+optional",
		"tcpSocket": "TCPSocket specifies an action involving a TCP port.\nTCP hooks not yet supported\n+optional",
	}
}

func (Probe) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "Probe describes a health check to be performed against a VirtualMachineInstance to determine whether it is\nalive or ready to receive traffic.",
		"initialDelaySeconds": "Number of seconds after the VirtualMachineInstance has started before liveness probes are initiated.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"timeoutSeconds":      "Number of seconds after which the probe times out.\nDefaults to 1 second. Minimum value is 1.\nMore info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes\n+optional",
		"periodSeconds":       "How often (in seconds) to perform the probe.\nDefault to 10 seconds. Minimum value is 1.\n+optional",
		"successThreshold":    "Minimum consecutive successes for the probe to be considered successful after having failed.\nDefaults to 1. Must be 1 for liveness. Minimum value is 1.\n+optional",
		"failureThreshold":    "Minimum consecutive failures for the probe to be considered failed after having succeeded.\nDefaults to 3. Minimum value is 1.\n+optional",
	}
}

func (KubeVirt) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirt represents the object deploying all KubeVirt resources",
	}
}

func (KubeVirtList) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirtList is a list of KubeVirts",
	}
}

func (KubeVirtSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"imagePullPolicy": "The ImagePullPolicy to use.",
	}
}

func (KubeVirtStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirtStatus represents information pertaining to a KubeVirt deployment.",
	}
}

func (KubeVirtCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "KubeVirtCondition represents a condition of a KubeVirt deployment",
	}
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceHealthScriptAssignment Contains properties used to assign a device management script to a group.
type DeviceHealthScriptAssignment struct {
	// Entity is the base model of DeviceHealthScriptAssignment
	Entity
	// Target The Azure Active Directory group we are targeting the script to
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
	// RunRemediationScript Determine whether we want to run detection script only or run both detection script and remediation script
	RunRemediationScript *bool `json:"runRemediationScript,omitempty"`
	// RunSchedule Script run schedule for the target group
	RunSchedule *RunSchedule `json:"runSchedule,omitempty"`
}
// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CalendarSharingMessageAction undocumented
type CalendarSharingMessageAction struct {
	// Object is the base model of CalendarSharingMessageAction
	Object
	// Importance undocumented
	Importance *CalendarSharingActionImportance `json:"importance,omitempty"`
	// ActionType undocumented
	ActionType *CalendarSharingActionType `json:"actionType,omitempty"`
	// Action undocumented
	Action *CalendarSharingAction `json:"action,omitempty"`
}
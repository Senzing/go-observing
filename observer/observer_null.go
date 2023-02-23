package observer

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ObserverNull is a simple example of the Subject interface.
// It is mainly used for testing.
type ObserverNull struct {
	Id       string
	IsSilent bool
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetObserverId method returns the unique identifier of the observer.
Use by the subject to manage the list of Observers.

Input
  - ctx: A context to control lifecycle.
*/
func (observer *ObserverNull) GetObserverId(ctx context.Context) string {
	return observer.Id
}

/*
The UpdateObserver method processes the message sent by the Subject.
The subject invokes UpdateObserver as a goroutine.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (observer *ObserverNull) UpdateObserver(ctx context.Context, message string) {
	if !observer.IsSilent {
		fmt.Printf("Observer: %s;  Message: %s\n", observer.Id, message)
	}
}

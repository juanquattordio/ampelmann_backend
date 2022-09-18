package errors

import "fmt"

type Message string
type Parameters map[string]interface{}

const (
	Generic        Message = "Error trying to %v %v."
	Connecting     Message = "Error trying to connect to resource."
	Getting        Message = "Error trying to get resource."
	Searching      Message = "Error trying to search resource."
	Saving         Message = "Error trying to save resource."
	Deleting       Message = "Error trying to delete resource."
	Locking        Message = "Error trying to lock a resource."
	Unlocking      Message = "Error trying to unlock a resource."
	Publishing     Message = "Error trying to publish notification to BigQ topic."
	BindingRequest Message = "Error binding request."
	Unmarshalling  Message = "Error unmarshalling response."
	Unavailable    Message = "Action not available."
	Status         Message = "Impossible change %v %v status from %v to %v."
	InterestRate   Message = "The interest rate must be higher than %v and lower than %v."
	AlreadyLocked  Message = "Resource is already locked and processed."
	DateRange      Message = "The period between first_due_date and start_date must be higher than %v days and lower than %v days."
)

func (message Message) String() string {
	return string(message)
}

func (message Message) Fmt(v ...interface{}) string {
	return fmt.Sprintf(message.String(), v...)
}

func (message Message) WithParams(values ...interface{}) string {
	msg := message.String()

	for i := 0; i < len(values); i += 2 {
		msg += fmt.Sprintf(" %v:%v.", values[i], values[i+1])
	}

	return msg
}

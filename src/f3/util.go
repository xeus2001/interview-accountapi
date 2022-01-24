package f3

import uuid "github.com/nu7hatch/gouuid"

// NewUuid creates a new random UUID. This can be used as identifier for accounts or organizations.
func NewUuid() (*uuid.UUID, Err) {
	v4, e := uuid.NewV4()
	if e != nil {
		return nil, err{ErrUuidCreationFailed, msgUuidCreationFailed, e}
	}
	return v4, nil
}

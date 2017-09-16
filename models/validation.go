package models

type Validate interface {
	Validate() (bool, error)
}

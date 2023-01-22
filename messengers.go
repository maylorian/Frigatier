package Frigatier

type Messenger interface {
	Name() string
	Notify(*Detection, string) error
	IsEnabled() bool
}

package repositories

type IMessageBus interface {
	Return(queue string, msg []byte) error
}

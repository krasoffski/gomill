package crawler

type task interface {
	process()
	output()
}

type factory interface {
	create(lint string) task
}

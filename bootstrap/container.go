package bootstrap

type Container struct {
	name string
	instance interface{}
}

type ContainerInterface interface {
	Set()
	Get()
}

var instances map[string]interface{}

func (container Container) Set() {
	instances = make(map[string]interface{})
	instances[container.name] = container.instance
}

func Get(name string) interface{} {
	return instances[name]
}

package main

type getConnection struct {
	instance ApplicationInstance
	inbox    chan *Connection
}

var StoreConnection = make(chan *Connection)

var getConnectionChan = make(chan getConnection)

func ConnectionStore() {
	connectionMap := make(map[ApplicationInstance]*Connection)
	select {
	case c := <-StoreConnection:
		connectionMap[c.instance] = c
	case g := <-getConnectionChan:
		instance := g.instance
		conn := connectionMap[instance]
		g.inbox <- conn
	}
}

func GetConnection(instance *ApplicationInstance) *Connection {
	inbox := make(chan *Connection)
	getConnectionChan <- getConnection{instance: *instance, inbox: inbox}
	conn := <-inbox
	return conn
}

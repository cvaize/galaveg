package bootstrap

import "galaveg/connections"

func CmdCloseConnections() {
	connections.DB.Close()
}

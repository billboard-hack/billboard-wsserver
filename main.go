// Author: "Shun Yokota"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"log"

	"billboard-wsserver/server"
)

func main() {
	log.Fatalln(server.Start())
}

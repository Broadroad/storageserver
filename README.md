<h1 align=center>
<img src="logo/512 horizontal.svg" width=60%>
</h1>

# The StorageServer in go

## Description
Storageserve is an easy-to-use, high performance storage. However, object - where-to-read map should store by user themselves.

As go is good at developing network and distributed system, so storageserver is writen in go.

## Features
* **Scalable**<br>
  Each storageserver only responsible for only onle node's read/write. So you can deploy as many as you can.

* **Dynamic Configurable**<br>
  Storageserver can have many config, and you can change them in config file and restart. Also, you can dynamic change these config without restart, however, this change will disappear after restart.

* **Dynamic Add or Remove Partition**<br>
  Storageserver can know partitions status periodically.

## License

Open source licensed under the MIT license (see _LICENSE_ file for details).

## Go version
go1.7+

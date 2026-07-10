package model

type Object string

const (
	ObjectCluster Object = "cluster"
	ObjectHost    Object = "host"
	ObjectVM      Object = "vm"
	ObjectPool    Object = "pool"
	ObjectOSD     Object = "osd"
)

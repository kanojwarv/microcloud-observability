package model

type Object string

const (
	ObjectVM      Object = "vm"
	ObjectHost    Object = "host"
	ObjectCluster Object = "cluster"
	ObjectPool    Object = "pool"
	ObjectOSD     Object = "osd"
)

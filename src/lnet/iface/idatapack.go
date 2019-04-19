package iface

/*
	封包数据和拆包数据
	直接面向TCP连接中的数据流,为传输数据添加头部信息，用于处理TCP粘包问题。
*/
type IDataPack interface {
	GetHeadLen() uint32                       //获取包头长度方法
	Pack(msg IMessagePackage) ([]byte, error) //封包方法
	Unpack([]byte) (IMessagePackage, error)   //拆包方法
}
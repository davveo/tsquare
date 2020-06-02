package idworker

import (
	"errors"
	"fmt"
)

func InitNode(nodeid int64) error {

	if nodeid > -1^(-1<<NodeIdBits) || nodeid < 0 {
		return errors.New(fmt.Sprintf("Node ID不合法，只能为0,1,2,且与其他机器不重复"))
	}
	NodeID = nodeid
	GetIdWorker()
	return nil
}

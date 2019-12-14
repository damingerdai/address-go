package utils

import sf "github.com/bwmarrin/snowflake"

// https://github.com/bwmarrin/snowflake
func CreateID() (id int64, err error) {
	node, err := sf.NewNode(1)
	if err != nil {
		id = -1
		return
	}
	id = node.Generate().Int64()
	err = nil
	return
}

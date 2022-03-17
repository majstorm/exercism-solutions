package tree

import (
	"errors"
	"sort"
)

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if records == nil {
		return nil, errors.New("No records provided")
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	nodes := make(map[int]*Node, len(records))

	for loc, record := range records {
		if (record.ID > 0 && record.ID == record.Parent) || record.ID < record.Parent || record.ID != loc {
			return nil, errors.New("Record is invalid")
		}

		nodes[record.ID] = &Node{ID: record.ID}
		if record.ID > 0 {
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
		}
	}

	return nodes[0], nil
}

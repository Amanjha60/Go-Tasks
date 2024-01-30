package main

import (
	"fmt"
)

// ProductMap represents the map of product families with associated node types
type ProductMap map[string][]string

// PrintMap prints the product map
func (pm ProductMap) PrintMap() {
	fmt.Println("Product Family -> Nodes:")
	for family, nodes := range pm {
		fmt.Printf("%s -> %v\n", family, nodes)
	}
	fmt.Println()
}

// AddFamily adds a new family to the product map
func (pm ProductMap) AddFamily(family string, nodes []string) {
	pm[family] = nodes
	fmt.Printf("Family %s added successfully.\n", family)
}

// AddNodeToFamily adds a node to the specified family if it doesn't already exist
func (pm ProductMap) AddNodeToFamily(family, node string) {
	nodes, ok := pm[family]
	if !ok {
		fmt.Printf("Family %s does not exist.\n", family)
		return
	}

	for _, n := range nodes {
		if n == node {
			fmt.Printf("Node %s already exists in family %s.\n", node, family)
			return
		}
	}

	pm[family] = append(nodes, node)
	fmt.Printf("Node %s added to family %s.\n", node, family)
}

// DeleteNode deletes a node from all families if it exists
func (pm ProductMap) DeleteNode(node string) {
	for family, nodes := range pm {
		for i, n := range nodes {
			if n == node {
				pm[family] = append(nodes[:i], nodes[i+1:]...)
				fmt.Printf("Node %s deleted from family %s.\n", node, family)
				break
			}
		}
	}
}

func main() {
	// Initialize the product map
	productMap := make(ProductMap)

	// Adding initial families and nodes
	productMap.AddFamily("SR", []string{"7750", "7450", "7950"})

	// Print initial map
	productMap.PrintMap()

	// Add new family and node
	productMap.AddFamily("XR", []string{"8200", "8400"})
	productMap.AddNodeToFamily("SR", "7250")

	// Print updated map
	productMap.PrintMap()

	// Delete deprecated node
	productMap.DeleteNode("7450")

	// Print final map
	productMap.PrintMap()
}

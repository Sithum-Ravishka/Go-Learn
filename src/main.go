package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MerkleNode represents a node in the Merkle tree.
type MerkleNode struct {
	Hash  string
	Left  *MerkleNode
	Right *MerkleNode
}

// calculateHash computes the SHA-256 hash of a given data string.
func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// buildMerkleTree recursively builds a Merkle tree from a list of data.
func buildMerkleTree(data []string) *MerkleNode {
	if len(data) == 1 {
		// If there's only one element in data, create a leaf node.
		return &MerkleNode{Hash: calculateHash(data[0]), Left: nil, Right: nil}
	}

	mid := len(data) / 2
	// Split the data into two halves.
	left := buildMerkleTree(data[:mid])
	// Recursively build the left subtree.
	right := buildMerkleTree(data[mid:])
	// Recursively build the right subtree.
	return &MerkleNode{Hash: calculateHash(left.Hash + right.Hash), Left: left, Right: right}
	// Create a new node with the combined hash of left and right, and set left and right child nodes.
}

// printTree recursively prints the Merkle tree.
func printTree(node *MerkleNode, indent string) {
	if node != nil {
		fmt.Println(indent+"Hash:", node.Hash)
		// Print the hash of the current node with indentation.
		if node.Left != nil {
			printTree(node.Left, indent+"  ")
			// Recursively print the left subtree.
		}
		if node.Right != nil {
			printTree(node.Right, indent+"  ")
			// Recursively print the right subtree.
		}
	}
}

func main() {
	data := []string{"2", "5", "5"}
	mid := len(data) / 2
	fmt.Println(mid)
	fmt.Println(data[:mid])
	fmt.Println(data[mid:])
	fmt.Println("Root Hash:", len(data))
	// Sample data for the Merkle tree.
	root := buildMerkleTree(data)
	// Build the Merkle tree from the data.
	printTree(root, "")
	// Print the tree structure.
	fmt.Println("Root Hash:", root.Hash)
	// Print the root hash of the Merkle tree.
}

// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//sourced from https://github.com/ethereum/go-ethereum/blob/master/trie/trie.go

package bidadjustment

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
)

// proofToPath converts a merkle proof to trie node path. The main purpose of
// this function is recovering a node path from the merkle proof stream. All
// necessary nodes will be resolved and leave the remaining as hashnode.
//
// The given edge proof is allowed to be an existent or non-existent proof.
func proofToPath(rootHash common.Hash, root node, key []byte, proofDb ethdb.KeyValueReader, allowNonExistent bool) (node, []byte, error) {
	// resolveNode retrieves and resolves trie node from merkle proof stream
	resolveNode := func(hash common.Hash) (node, error) {
		buf, _ := proofDb.Get(hash[:])
		if buf == nil {
			return nil, fmt.Errorf("proof node (hash %064x) missing", hash)
		}
		n, err := decodeNode(hash[:], buf)
		if err != nil {
			return nil, fmt.Errorf("bad proof node %v", err)
		}
		return n, err
	}
	// If the root node is empty, resolve it first.
	// Root node must be included in the proof.
	if root == nil {
		n, err := resolveNode(rootHash)
		if err != nil {
			return nil, nil, err
		}
		root = n
	}
	var (
		err           error
		child, parent node
		keyrest       []byte
		valnode       []byte
	)
	key, parent = keybytesToHex(key), root
	for {
		keyrest, child = get(parent, key, false)
		switch cld := child.(type) {
		case nil:
			// The trie doesn't contain the key. It's possible
			// the proof is a non-existing proof, but at least
			// we can prove all resolved nodes are correct, it's
			// enough for us to prove range.
			if allowNonExistent {
				return root, nil, nil
			}
			return nil, nil, errors.New("the node is not contained in trie")
		case *shortNode:
			key, parent = keyrest, child // Already resolved
			continue
		case *fullNode:
			key, parent = keyrest, child // Already resolved
			continue
		case hashNode:
			child, err = resolveNode(common.BytesToHash(cld))
			if err != nil {
				return nil, nil, err
			}
		case valueNode:
			valnode = cld
		}
		// Link the parent and child.
		switch pnode := parent.(type) {
		case *shortNode:
			pnode.Val = child
		case *fullNode:
			pnode.Children[key[0]] = child
		default:
			panic(fmt.Sprintf("%T: invalid node: %v", pnode, pnode))
		}
		if len(valnode) > 0 {
			return root, valnode, nil // The whole path is resolved
		}
		key, parent = keyrest, child
	}
}

func (n *fullNode) encode(w rlp.EncoderBuffer) {
	offset := w.List()
	for _, c := range n.Children {
		if c != nil {
			c.encode(w)
		} else {
			w.Write(rlp.EmptyString) //nolint:errcheck
		}
	}
	w.ListEnd(offset)
}

func (n *shortNode) encode(w rlp.EncoderBuffer) {
	offset := w.List()
	w.WriteBytes(n.Key)
	if n.Val != nil {
		n.Val.encode(w)
	} else {
		w.Write(rlp.EmptyString) //nolint:errcheck
	}
	w.ListEnd(offset)
}

func (n hashNode) encode(w rlp.EncoderBuffer) {
	w.WriteBytes(n)
}

func (n valueNode) encode(w rlp.EncoderBuffer) {
	w.WriteBytes(n)
}

//nolint:unused
func (n rawNode) encode(w rlp.EncoderBuffer) {
	w.Write(n) //nolint:errcheck
}

// get returns the child of the given node. Return nil if the
// node with specified key doesn't exist at all.
//
// There is an additional flag `skipResolved`. If it's set then
// all resolved nodes won't be returned.
func get(tn node, key []byte, skipResolved bool) ([]byte, node) {
	for {
		switch n := tn.(type) {
		case *shortNode:
			if len(key) < len(n.Key) || !bytes.Equal(n.Key, key[:len(n.Key)]) {
				return nil, nil
			}
			tn = n.Val
			key = key[len(n.Key):]
			if !skipResolved {
				return key, tn
			}
		case *fullNode:
			tn = n.Children[key[0]]
			key = key[1:]
			if !skipResolved {
				return key, tn
			}
		case hashNode:
			return key, n
		case nil:
			return key, nil
		case valueNode:
			return nil, n
		default:
			panic(fmt.Sprintf("%T: invalid node: %v", tn, tn))
		}
	}
}

//Blxr Custom Code

type RootNode struct {
	Node node
}

func (r RootNode) updateNode(n node, prefix, key []byte, value node) (bool, node, error) {
	if len(key) == 0 {
		if v, ok := n.(valueNode); ok {
			return !bytes.Equal(v, value.(valueNode)), value, nil
		}
		return true, value, nil
	}
	switch n := n.(type) {
	case *shortNode:
		matchlen := prefixLen(key, n.Key)
		// If the whole key matches, keep this short node as is
		// and only update the value.
		if matchlen == len(n.Key) {
			dirty, nn, err := r.updateNode(n.Val, append(prefix, key[:matchlen]...), key[matchlen:], value)
			if !dirty || err != nil {
				return false, n, err
			}
			return true, &shortNode{n.Key, nn, nodeFlag{dirty: true}}, nil
		}
		// Otherwise branch out at the index where they differ.
		branch := &fullNode{flags: nodeFlag{dirty: true}}
		var err error
		_, branch.Children[n.Key[matchlen]], err = r.updateNode(nil, append(prefix, n.Key[:matchlen+1]...), n.Key[matchlen+1:], n.Val)
		if err != nil {
			return false, nil, err
		}
		_, branch.Children[key[matchlen]], err = r.updateNode(nil, append(prefix, key[:matchlen+1]...), key[matchlen+1:], value)
		if err != nil {
			return false, nil, err
		}
		// Replace this shortNode with the branch if it occurs at index 0.
		if matchlen == 0 {
			return true, branch, nil
		}
		// New branch node is created as a child of the original short node.
		// Track the newly inserted node in the tracer. The node identifier
		// passed is the path from the root node.
		// t.tracer.onInsert(append(prefix, key[:matchlen]...))

		// Replace it with a short node leading up to the branch.
		return true, &shortNode{key[:matchlen], branch, nodeFlag{dirty: true}}, nil

	case *fullNode:
		dirty, nn, err := r.updateNode(n.Children[key[0]], append(prefix, key[0]), key[1:], value)
		if !dirty || err != nil {
			return false, n, err
		}
		n = n.copy()
		n.flags = nodeFlag{dirty: true}
		n.Children[key[0]] = nn
		return true, n, nil

	case nil:
		// New short node is created and track it in the tracer. The node identifier
		// passed is the path from the root node. Note the valueNode won't be tracked
		// since it's always embedded in its parent.
		// t.tracer.onInsert(prefix)

		return true, &shortNode{key, value, nodeFlag{dirty: true}}, nil

	// case hashNode:
	// We've hit a part of the trie that isn't loaded yet. Load
	// the node and insert into it. This leaves all child nodes on
	// the path to the value in the trie.
	// rn, err := t.resolveAndTrack(n, prefix)
	// if err != nil {
	// 	return false, nil, err
	// }
	// dirty, nn, err := t.insert(rn, prefix, key, value)
	// if !dirty || err != nil {
	// 	return false, rn, err
	// }
	// return true, nn, nil

	default:
		panic(fmt.Sprintf("%T: invalid node: %v", n, n))
	}
}

func (root RootNode) Hash() (node, node) {
	// func (t *Trie) hashRoot() (node, node) {
	if root.Node == nil {
		return hashNode(types.EmptyRootHash.Bytes()), nil
	}
	// If the number of changes is below 100, we let one thread handle it
	h := newHasher(false)
	defer func() {
		returnHasherToPool(h)
	}()
	hashed, updatedNode := h.hash(root.Node, true)
	return hashed, updatedNode
}

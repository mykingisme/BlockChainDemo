package main

type BlockChain struct {
	blocks []*Block
}

func CreateBlockChain() *BlockChain {
	return &BlockChain{[]*Block{CreateGenesisBlock()}}
}
func (bc *BlockChain) AddBlock(data string) {
	bc.blocks = append(bc.blocks, CreateBlock(data, bc.blocks[len(bc.blocks)-1].Hash))
}

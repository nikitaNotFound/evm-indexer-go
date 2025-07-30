package networks

const (
	ETHUniswapV2FactoryAddress      = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	ArbitrumUniswapV2FactoryAddress = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	OptimismUniswapV2FactoryAddress = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	BaseUniswapV2FactoryAddress     = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	BNBUniswapV2FactoryAddress      = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
)

func GetUniswapV2FactoryAddress(network Network) string {
	switch network {
	case ETH:
		return ETHUniswapV2FactoryAddress
	case Arbitrum:
		return ArbitrumUniswapV2FactoryAddress
	case Optimism:
		return OptimismUniswapV2FactoryAddress
	case Base:
		return BaseUniswapV2FactoryAddress
	case BNB:
		return BNBUniswapV2FactoryAddress
	default:
		return ""
	}
}

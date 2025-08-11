package networks

const (
	ETHUniswapV2FactoryAddress      = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
	ArbitrumUniswapV2FactoryAddress = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	OptimismUniswapV2FactoryAddress = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	BaseUniswapV2FactoryAddress     = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	BNBUniswapV2FactoryAddress      = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"

	ETHUniswapV3FactoryAddress      = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	ArbitrumUniswapV3FactoryAddress = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	OptimismUniswapV3FactoryAddress = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	BaseUniswapV3FactoryAddress     = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
	BNBUniswapV3FactoryAddress      = "0x1F98431c8aD98523631AE4a59f267346ea31F984"
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

func GetUniswapV3FactoryAddress(network Network) string {
	switch network {
	case ETH:
		return ETHUniswapV3FactoryAddress
	case Arbitrum:
		return ArbitrumUniswapV3FactoryAddress
	case Optimism:
		return OptimismUniswapV3FactoryAddress
	case Base:
		return BaseUniswapV3FactoryAddress
	case BNB:
		return BNBUniswapV3FactoryAddress
	default:
		return ""
	}
}

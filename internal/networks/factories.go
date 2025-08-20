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

	ETHPancakeV2FactoryAddress      = "0x1097053Fd2ea711dad45caCcc45EfF7548fCB362"
	ArbitrumPancakeV2FactoryAddress = "0x02a84c1b3BBD7401a5f7fa98a384EBC70bB5749E"
	OptimismPancakeV2FactoryAddress = "0x02a84c1b3BBD7401a5f7fa98a384EBC70bB5749E"
	BasePancakeV2FactoryAddress     = "0x02a84c1b3BBD7401a5f7fa98a384EBC70bB5749E"
	BNBPancakeV2FactoryAddress      = "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73"

	ETHPancakeV3FactoryAddress      = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	ArbitrumPancakeV3FactoryAddress = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	OptimismPancakeV3FactoryAddress = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	BasePancakeV3FactoryAddress     = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"
	BNBPancakeV3FactoryAddress      = "0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865"

	// SushiSwap V2 Factory addresses
	ETHSushiSwapV2FactoryAddress      = "0xc0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac"
	ArbitrumSushiSwapV2FactoryAddress = "0xc35dadb65012ec5796536bd9864ed8773abc74c4"
	OptimismSushiSwapV2FactoryAddress = "0xc35dadb65012ec5796536bd9864ed8773abc74c4"
	BaseSushiSwapV2FactoryAddress     = "0xc35dadb65012ec5796536bd9864ed8773abc74c4"
	BNBSushiSwapV2FactoryAddress      = "0xc35dadb65012ec5796536bd9864ed8773abc74c4"

	// SushiSwap V3 Factory addresses
	ETHSushiSwapV3FactoryAddress      = "0xbaceb8ec6b9355dfc0269c18bac9d6e2bdc29c4f"
	ArbitrumSushiSwapV3FactoryAddress = "0x1af415a1eba07a4986a52b6f2e7de7003d82231e"
	OptimismSushiSwapV3FactoryAddress = "0x93395129bd3fcf49d95730d3c2737c17990ff328"
	BaseSushiSwapV3FactoryAddress     = "0xc35dadb65012ec5796536bd9864ed8773abc74c4"
	BNBSushiSwapV3FactoryAddress      = "0x126555dd55a39328f69400d6ae4f782bd4c34abb"
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

func GetPancakeV2FactoryAddress(network Network) string {
	switch network {
	case ETH:
		return ETHPancakeV2FactoryAddress
	case Arbitrum:
		return ArbitrumPancakeV2FactoryAddress
	case Optimism:
		return OptimismPancakeV2FactoryAddress
	case Base:
		return BasePancakeV2FactoryAddress
	case BNB:
		return BNBPancakeV2FactoryAddress
	default:
		return ""
	}
}

func GetPancakeV3FactoryAddress(network Network) string {
	switch network {
	case ETH:
		return ETHPancakeV3FactoryAddress
	case Arbitrum:
		return ArbitrumPancakeV3FactoryAddress
	case Optimism:
		return OptimismPancakeV3FactoryAddress
	case Base:
		return BasePancakeV3FactoryAddress
	case BNB:
		return BNBPancakeV3FactoryAddress
	default:
		return ""
	}
}

func GetSushiSwapV2FactoryAddress(network Network) string {
	switch network {
	case ETH:
		return ETHSushiSwapV2FactoryAddress
	case Arbitrum:
		return ArbitrumSushiSwapV2FactoryAddress
	case Optimism:
		return OptimismSushiSwapV2FactoryAddress
	case Base:
		return BaseSushiSwapV2FactoryAddress
	case BNB:
		return BNBSushiSwapV2FactoryAddress
	default:
		return ""
	}
}

func GetSushiSwapV3FactoryAddress(network Network) string {
	switch network {
	case ETH:
		return ETHSushiSwapV3FactoryAddress
	case Arbitrum:
		return ArbitrumSushiSwapV3FactoryAddress
	case Optimism:
		return OptimismSushiSwapV3FactoryAddress
	case Base:
		return BaseSushiSwapV3FactoryAddress
	case BNB:
		return BNBSushiSwapV3FactoryAddress
	default:
		return ""
	}
}

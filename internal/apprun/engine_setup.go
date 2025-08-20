package apprun

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nikitaNotFound/evm-indexer-go/internal/config"
	"github.com/nikitaNotFound/evm-indexer-go/internal/engine"
	"github.com/nikitaNotFound/evm-indexer-go/internal/indexers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/producers"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres"
	"github.com/rs/zerolog/log"
)

func setupEngine(cfg *config.Config, ethClient *ethclient.Client, pgStorage *postgres.Storage) *engine.Engine {
	blocksProducer := producers.NewBlocksProducer(ethClient)
	blocksIndexer := indexers.NewBlocksIndexer(pgStorage)
	rawTxsIndexer := indexers.NewRawTxsIndexer(pgStorage)

	uniswapV2Producer := producers.NewUniswapV2PoolsProducer(ethClient, cfg)
	uniswapV2Indexer := indexers.NewUniswapV2Indexer(pgStorage)

	uniswapV3Producer := producers.NewUniswapV3PoolsProducer(ethClient, cfg)
	uniswapV3Indexer := indexers.NewUniswapV3Indexer(pgStorage)

	pancakeV2Producer := producers.NewPancakeV2PoolsProducer(ethClient, cfg)
	pancakeV2Indexer := indexers.NewPancakeV2Indexer(pgStorage)

	pancakeV3Producer := producers.NewPancakeV3PoolsProducer(ethClient, cfg)
	pancakeV3Indexer := indexers.NewPancakeV3Indexer(pgStorage)

	sushiSwapV2Producer := producers.NewSushiSwapV2PoolsProducer(ethClient, cfg)
	sushiSwapV2Indexer := indexers.NewSushiSwapV2Indexer(pgStorage)

	sushiSwapV3Producer := producers.NewSushiSwapV3PoolsProducer(ethClient, cfg)
	sushiSwapV3Indexer := indexers.NewSushiSwapV3Indexer(pgStorage)

	engine := engine.CreateEngine(cfg, []engine.DataProducer{
		blocksProducer,
		uniswapV2Producer,
		uniswapV3Producer,
		pancakeV2Producer,
		pancakeV3Producer,
		sushiSwapV2Producer,
		sushiSwapV3Producer,
	})

	if err := engine.IndexersGate().CreateTopic(producers.BlocksTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create blocks topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.BlocksTopicName, blocksIndexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to blocks topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.RawTxsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create raw_txs topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.RawTxsTopicName, rawTxsIndexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to raw_txs topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.UniswapV2PoolsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create uniswap_v2_pools topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.UniswapV2PoolsTopicName, uniswapV2Indexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to uniswap_v2_pools topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.UniswapV3PoolsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create uniswap_v3_pools topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.UniswapV3PoolsTopicName, uniswapV3Indexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to uniswap_v3_pools topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.PancakeV2PoolsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create pancake_v2_pools topic")
	}

	if err := engine.IndexersGate().Subscribe(
		producers.PancakeV2PoolsTopicName,
		pancakeV2Indexer,
	); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to pancake_v2_pools topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.PancakeV3PoolsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create pancake_v3_pools topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.PancakeV3PoolsTopicName, pancakeV3Indexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to pancake_v3_pools topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.SushiSwapV2PoolsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create sushiswap_v2_pools topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.SushiSwapV2PoolsTopicName, sushiSwapV2Indexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to sushiswap_v2_pools topic")
	}

	if err := engine.IndexersGate().CreateTopic(producers.SushiSwapV3PoolsTopicName); err != nil {
		log.Fatal().Err(err).Msg("failed to create sushiswap_v3_pools topic")
	}

	if err := engine.IndexersGate().Subscribe(producers.SushiSwapV3PoolsTopicName, sushiSwapV3Indexer); err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe to sushiswap_v3_pools topic")
	}

	return &engine
}

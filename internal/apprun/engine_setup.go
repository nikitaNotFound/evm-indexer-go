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

	engine := engine.CreateEngine(cfg, []engine.DataProducer{
		blocksProducer,
		uniswapV2Producer,
		uniswapV3Producer,
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

	return &engine
}

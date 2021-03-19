package graphql

import (
	"context"
	"testing"

	"{{cookiecutter.go_module_name}}/app/basic/demo/proto/model/graphql"
)

func TestChainGraph_Block(t *testing.T) {
	var qs graphql.Query

	err := testDao.Chain.Block(context.Background(), &qs)
	t.Logf("graphql query resp: %v, err: %v", qs, err)
}

func TestChainGraph_CreateMarketEvents(t *testing.T) {
	var qs graphql.QueryCreateMarketEvents

	err := testDao.Chain.CreateMarketEvents(context.Background(), &qs)
	t.Logf("graphql query resp: %v, err: %v", qs, err)
}

func TestChainGraph_FinalizeMarketEvents(t *testing.T) {
	var qs graphql.QueryFinalizeMarketEvents

	err := testDao.Chain.FinalizeMarketEvents(context.Background(), &qs)
	t.Logf("graphql query resp: %v, err: %v", qs, err)
}

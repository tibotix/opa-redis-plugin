package internal

import (
	"github.com/tibotix/opa-redis-plugin/utils"

	"github.com/go-redis/redis/v8"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

func registerDO(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.do",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Any{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []interface{}
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val, err := rdb.Do(p.redisContext, utils.Conva(v0)...).Text()
			switch {
			case err == redis.Nil:
				return ast.NullTerm(), nil
			case err != nil:
				return nil, err
			default:
				return ast.StringTerm(val), nil
			}
		},
	)
}

func (p *redisPlugin) registerManualCommands() {
	registerDO(p)
}

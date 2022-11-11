package internal

import (
	"github.com/go-redis/redis/v8"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// GENERIC COMMANDS
func (p *redisPlugin) registerGET() {
	rego.RegisterBuiltin1(
		&rego.Function{
			Name:             "redis.get",
			Decl:             types.NewFunction(types.Args(types.S), types.S),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var key string
			if err := ast.As(a.Value, &key); err != nil {
				return nil, err
			}

			val, err := rdb.Get(p.redisContext, key).Result()
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

func (p *redisPlugin) registerEXISTS() {
	rego.RegisterBuiltin1(
		&rego.Function{
			Name:             "redis.exists",
			Decl:             types.NewFunction(types.Args(types.S), types.N),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var key string
			if err := ast.As(a.Value, &key); err != nil {
				return nil, err
			}

			// TODO: support List of keys
			val, err := rdb.Exists(p.redisContext, key).Result()
			switch {
			case err == redis.Nil:
				return ast.NullTerm(), nil
			case err != nil:
				return nil, err
			default:
				return ast.IntNumberTerm(int(val)), nil
			}
		},
	)
}

// SET COMMANDS
func (p *redisPlugin) registerSISMEMBER() {
	rego.RegisterBuiltin2(
		&rego.Function{
			Name:             "redis.sismember",
			Decl:             types.NewFunction(types.Args(types.S, types.S), types.B),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, a, b *ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var key, member string
			if err := ast.As(a.Value, &key); err != nil {
				return nil, err
			}
			if err := ast.As(b.Value, &member); err != nil {
				return nil, err
			}

			val, err := rdb.SIsMember(p.redisContext, key, member).Result()
			switch {
			case err == redis.Nil:
				return ast.NullTerm(), nil
			case err != nil:
				return nil, err
			default:
				return ast.BooleanTerm(val), nil
			}
		},
	)
}

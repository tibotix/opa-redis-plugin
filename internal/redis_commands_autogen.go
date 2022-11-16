// Code generated! DO NOT EDIT
package internal

import (
	"github.com/tibotix/opa-redis-plugin/utils"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

func registerCLIENTGETNAME(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientgetname",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClientGetName(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerECHO(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.echo",
			Decl:             types.NewFunction(types.Args(types.Any{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 interface{}
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Echo(p.redisContext, utils.Conv(v0))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPING(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.ping",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Ping(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerQUIT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.quit",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Quit(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDEL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.del",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Del(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerUNLINK(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.unlink",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Unlink(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDUMP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.dump",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Dump(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXISTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.exists",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Exists(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIRE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expire",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Expire(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIREAT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expireat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Time
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireAt(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIRENX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expirenx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireNX(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIREXX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expirexx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireXX(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIREGT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expiregt",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireGT(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEXPIRELT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.expirelt",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ExpireLT(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerKEYS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.keys",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Keys(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp07421 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp07421 = append(tmp07421, term)
				}
				term := ast.ArrayTerm(tmp07421...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerMIGRATE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.migrate",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			var v4 time.Duration
			if err := ast.As(terms[4].Value, &v4); err != nil {
				return nil, err
			}

			val := rdb.Migrate(p.redisContext, v0, v1, v2, v3, v4)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMOVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.move",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Move(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerOBJECTREFCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.objectrefcount",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ObjectRefCount(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerOBJECTENCODING(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.objectencoding",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ObjectEncoding(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerOBJECTIDLETIME(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.objectidletime",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ObjectIdleTime(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.Seconds()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPERSIST(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.persist",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Persist(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPEXPIRE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pexpire",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PExpire(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPEXPIREAT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pexpireat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Time
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PExpireAt(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPTTL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pttl",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PTTL(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.Milliseconds()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRANDOMKEY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.randomkey",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.RandomKey(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRENAME(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rename",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Rename(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRENAMENX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.renamenx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RenameNX(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRESTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.restore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Restore(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRESTOREREPLACE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.restorereplace",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.RestoreReplace(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSORT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sort",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Order", types.String{}), types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Sort
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Sort(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp99986 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp99986 = append(tmp99986, term)
				}
				term := ast.ArrayTerm(tmp99986...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSORTSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sortstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Order", types.String{}), types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.Sort
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SortStore(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSORTINTERFACES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sortinterfaces",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Order", types.String{}), types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Sort
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SortInterfaces(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp50135 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp50135 = append(tmp50135, term)
				}
				term := ast.ArrayTerm(tmp50135...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerTOUCH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.touch",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Touch(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerTTL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.ttl",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.TTL(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.Seconds()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerTYPE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.type",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Type(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerAPPEND(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.append",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Append(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDECR(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.decr",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Decr(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDECRBY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.decrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.DecrBy(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.get",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Get(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GetRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GetSet(p.redisContext, v0, utils.Conv(v1))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETEX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 time.Duration
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GetEx(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETDEL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getdel",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.GetDel(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINCR(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.incr",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Incr(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINCRBY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.incrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.IncrBy(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINCRBYFLOAT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.incrbyfloat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.IncrByFloat(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMGET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.mget",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.MGet(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp11677 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp11677 = append(tmp11677, term)
				}
				term := ast.ArrayTerm(tmp11677...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerMSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.mset",
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

			val := rdb.MSet(p.redisContext, utils.Conva(v0)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMSETNX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.msetnx",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Any{})), types.Boolean{}),
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

			val := rdb.MSetNX(p.redisContext, utils.Conva(v0)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.set",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Set(p.redisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETARGS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setargs",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Mode", types.String{}), types.NewStaticProperty("TTL", types.Number{}), types.NewStaticProperty("ExpireAt", types.Number{}), types.NewStaticProperty("Get", types.Boolean{}), types.NewStaticProperty("KeepTTL", types.Boolean{})}, nil)), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 redis.SetArgs
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetArgs(p.redisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETEX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetEX(p.redisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETNX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetNX(p.redisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETXX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setxx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetXX(p.redisContext, v0, utils.Conv(v1), v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSTRLEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.strlen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.StrLen(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCOPY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.copy",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}, types.Boolean{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 bool
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.Copy(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGETBIT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.getbit",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GetBit(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSETBIT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.setbit",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SetBit(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Start", types.Number{}), types.NewStaticProperty("End", types.Number{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.BitCount
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitCount(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPAND(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopand",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpAnd(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPOR(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopor",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpOr(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPXOR(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopxor",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpXor(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITOPNOT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitopnot",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitOpNot(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITPOS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitpos",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.NewArray([]types.Type{}, types.Number{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.BitPos(p.redisContext, v0, v1, v2...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBITFIELD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bitfield",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.NewArray([]types.Type{}, types.Number{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BitField(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp71571 []*ast.Term
				for _, v := range r0 {
					term := ast.IntNumberTerm(int(v))
					tmp71571 = append(tmp71571, term)
				}
				term := ast.ArrayTerm(tmp71571...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSCAN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scan",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 uint64
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Scan(p.redisContext, v0, v1, v2)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp15560 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp15560 = append(tmp15560, term)
				}
				tr0 := ast.ArrayTerm(tmp15560...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerSCANTYPE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scantype",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.String{}, types.Number{}, types.String{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 uint64
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.ScanType(p.redisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp72799 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp72799 = append(tmp72799, term)
				}
				tr0 := ast.ArrayTerm(tmp72799...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerSSCAN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sscan",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 uint64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.SScan(p.redisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp58212 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp58212 = append(tmp58212, term)
				}
				tr0 := ast.ArrayTerm(tmp58212...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerHSCAN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hscan",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 uint64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.HScan(p.redisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp13495 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp13495 = append(tmp13495, term)
				}
				tr0 := ast.ArrayTerm(tmp13495...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerZSCAN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zscan",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}, types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.Number{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 uint64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.ZScan(p.redisContext, v0, v1, v2, v3)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp27284 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp27284 = append(tmp27284, term)
				}
				tr0 := ast.ArrayTerm(tmp27284...)
				tr1 := ast.UIntNumberTerm(uint64(r1))
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerHDEL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hdel",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HDel(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHEXISTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hexists",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HExists(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHGET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hget",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HGet(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHGETALL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hgetall",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.String{}))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HGetAll(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp46110 [][2]*ast.Term
				for key, value := range r0 {
					k := ast.StringTerm(key)
					v := ast.StringTerm(value)
					tmp46110 = append(tmp46110, [2]*ast.Term{k, v})
				}
				term := ast.ObjectTerm(tmp46110...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHINCRBY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hincrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HIncrBy(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHINCRBYFLOAT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hincrbyfloat",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 float64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HIncrByFloat(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHKEYS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hkeys",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HKeys(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp68926 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp68926 = append(tmp68926, term)
				}
				term := ast.ArrayTerm(tmp68926...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHLEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hlen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HLen(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHMGET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hmget",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HMGet(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp43354 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp43354 = append(tmp43354, term)
				}
				term := ast.ArrayTerm(tmp43354...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HSet(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHMSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hmset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.HMSet(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHSETNX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hsetnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Any{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HSetNX(p.redisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerHVALS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hvals",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.HVals(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp92667 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp92667 = append(tmp92667, term)
				}
				term := ast.ArrayTerm(tmp92667...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerHRANDFIELD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.hrandfield",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Boolean{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 bool
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.HRandField(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp16482 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp16482 = append(tmp16482, term)
				}
				term := ast.ArrayTerm(tmp16482...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBLPOP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.blpop",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BLPop(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp33807 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp33807 = append(tmp33807, term)
				}
				term := ast.ArrayTerm(tmp33807...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBRPOP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.brpop",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BRPop(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp96814 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp96814 = append(tmp96814, term)
				}
				term := ast.ArrayTerm(tmp96814...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBRPOPLPUSH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.brpoplpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 time.Duration
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.BRPopLPush(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINDEX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lindex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LIndex(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINSERT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.linsert",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Any{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 interface{}
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.LInsert(p.redisContext, v0, v1, utils.Conv(v2), utils.Conv(v3))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINSERTBEFORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.linsertbefore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LInsertBefore(p.redisContext, v0, utils.Conv(v1), utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLINSERTAFTER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.linsertafter",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LInsertAfter(p.redisContext, v0, utils.Conv(v1), utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLLEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.llen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.LLen(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPOP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpop",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.LPop(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPOPCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpopcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LPopCount(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp47838 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp47838 = append(tmp47838, term)
				}
				term := ast.ArrayTerm(tmp47838...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerLPOS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpos",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Rank", types.Number{}), types.NewStaticProperty("MaxLen", types.Number{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 redis.LPosArgs
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LPos(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPOSCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lposcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Rank", types.Number{}), types.NewStaticProperty("MaxLen", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.Number{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 redis.LPosArgs
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.LPosCount(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp90547 []*ast.Term
				for _, v := range r0 {
					term := ast.IntNumberTerm(int(v))
					tmp90547 = append(tmp90547, term)
				}
				term := ast.ArrayTerm(tmp90547...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerLPUSH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LPush(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLPUSHX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lpushx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.LPushX(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp65802 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp65802 = append(tmp65802, term)
				}
				term := ast.ArrayTerm(tmp65802...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerLREM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lrem",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LRem(p.redisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Any{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LSet(p.redisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLTRIM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.ltrim",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.LTrim(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPOP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpop",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.RPop(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPOPCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpopcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPopCount(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp08693 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp08693 = append(tmp08693, term)
				}
				term := ast.ArrayTerm(tmp08693...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerRPOPLPUSH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpoplpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPopLPush(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPUSH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpush",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPush(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerRPUSHX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.rpushx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.RPushX(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLMOVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lmove",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.LMove(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBLMOVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.blmove",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.String{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			var v4 time.Duration
			if err := ast.As(terms[4].Value, &v4); err != nil {
				return nil, err
			}

			val := rdb.BLMove(p.redisContext, v0, v1, v2, v3, v4)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSADD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SAdd(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCARD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scard",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SCard(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSDIFF(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sdiff",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SDiff(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp30167 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp30167 = append(tmp30167, term)
				}
				term := ast.ArrayTerm(tmp30167...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSDIFFSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sdiffstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SDiffStore(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSINTER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sinter",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SInter(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp92740 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp92740 = append(tmp92740, term)
				}
				term := ast.ArrayTerm(tmp92740...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSINTERSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sinterstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SInterStore(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSISMEMBER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sismember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SIsMember(p.redisContext, v0, utils.Conv(v1))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSMISMEMBER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smismember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.NewArray([]types.Type{}, types.Boolean{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SMIsMember(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp44620 []*ast.Term
				for _, v := range r0 {
					term := ast.BooleanTerm(v)
					tmp44620 = append(tmp44620, term)
				}
				term := ast.ArrayTerm(tmp44620...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSMEMBERS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smembers",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SMembers(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp87739 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp87739 = append(tmp87739, term)
				}
				term := ast.ArrayTerm(tmp87739...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSMEMBERSMAP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smembersmap",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.NewObject([]*types.StaticProperty{}, nil)))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SMembersMap(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp98252 [][2]*ast.Term
				for key, _ := range r0 {
					k := ast.StringTerm(key)

					v := ast.ObjectTerm()

					tmp98252 = append(tmp98252, [2]*ast.Term{k, v})
				}
				term := ast.ObjectTerm(tmp98252...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSMOVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.smove",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Any{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.SMove(p.redisContext, v0, v1, utils.Conv(v2))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSPOP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.spop",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SPop(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSPOPN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.spopn",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SPopN(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp66882 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp66882 = append(tmp66882, term)
				}
				term := ast.ArrayTerm(tmp66882...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSRANDMEMBER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.srandmember",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SRandMember(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSRANDMEMBERN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.srandmembern",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SRandMemberN(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp19968 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp19968 = append(tmp19968, term)
				}
				term := ast.ArrayTerm(tmp19968...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSREM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.srem",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SRem(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSUNION(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sunion",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.SUnion(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp73837 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp73837 = append(tmp73837, term)
				}
				term := ast.ArrayTerm(tmp73837...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSUNIONSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.sunionstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SUnionStore(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXADD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xadd",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("NoMkStream", types.Boolean{}), types.NewStaticProperty("MaxLen", types.Number{}), types.NewStaticProperty("MaxLenApprox", types.Number{}), types.NewStaticProperty("MinID", types.String{}), types.NewStaticProperty("Approx", types.Boolean{}), types.NewStaticProperty("Limit", types.Number{}), types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.Any{})}, nil)), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XAddArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XAdd(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXDEL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xdel",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XDel(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXLEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xlen",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XLen(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp67339 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp51809 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp51809 = append(tmp51809, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp51809...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp67339 = append(tmp67339, term)
				}
				term := ast.ArrayTerm(tmp67339...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXRANGEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrangen",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.XRangeN(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp81656 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp41300 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp41300 = append(tmp41300, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp41300...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp81656 = append(tmp81656, term)
				}
				term := ast.ArrayTerm(tmp81656...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREVRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrevrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XRevRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp97434 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp79153 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp79153 = append(tmp79153, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp79153...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp97434 = append(tmp97434, term)
				}
				term := ast.ArrayTerm(tmp97434...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREVRANGEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xrevrangen",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 int64
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.XRevRangeN(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp95015 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp73064 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp73064 = append(tmp73064, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp73064...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp95015 = append(tmp95015, term)
				}
				term := ast.ArrayTerm(tmp95015...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREAD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xread",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Streams", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Block", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XReadArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XRead(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp66155 []*ast.Term
				for _, v := range r0 {

					termStream := ast.StringTerm(v.Stream)

					var tmp67269 []*ast.Term
					for _, v := range v.Messages {

						termID := ast.StringTerm(v.ID)

						var tmp85885 [][2]*ast.Term
						for key, value := range v.Values {
							k := ast.StringTerm(key)

							v := ast.NullTerm()
							if s, ok := value.(string); ok {
								v = ast.StringTerm(s)
							}

							tmp85885 = append(tmp85885, [2]*ast.Term{k, v})
						}
						termValues := ast.ObjectTerm(tmp85885...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

						tmp67269 = append(tmp67269, term)
					}
					termMessages := ast.ArrayTerm(tmp67269...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Stream"), termStream}, [2]*ast.Term{ast.StringTerm("Messages"), termMessages})

					tmp66155 = append(tmp66155, term)
				}
				term := ast.ArrayTerm(tmp66155...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXREADSTREAMS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xreadstreams",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XReadStreams(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp52548 []*ast.Term
				for _, v := range r0 {

					termStream := ast.StringTerm(v.Stream)

					var tmp27304 []*ast.Term
					for _, v := range v.Messages {

						termID := ast.StringTerm(v.ID)

						var tmp53913 [][2]*ast.Term
						for key, value := range v.Values {
							k := ast.StringTerm(key)

							v := ast.NullTerm()
							if s, ok := value.(string); ok {
								v = ast.StringTerm(s)
							}

							tmp53913 = append(tmp53913, [2]*ast.Term{k, v})
						}
						termValues := ast.ObjectTerm(tmp53913...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

						tmp27304 = append(tmp27304, term)
					}
					termMessages := ast.ArrayTerm(tmp27304...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Stream"), termStream}, [2]*ast.Term{ast.StringTerm("Messages"), termMessages})

					tmp52548 = append(tmp52548, term)
				}
				term := ast.ArrayTerm(tmp52548...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPCREATE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupcreate",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupCreate(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPCREATEMKSTREAM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupcreatemkstream",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupCreateMkStream(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPSETID(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupsetid",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupSetID(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPDESTROY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupdestroy",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XGroupDestroy(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPCREATECONSUMER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupcreateconsumer",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupCreateConsumer(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXGROUPDELCONSUMER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xgroupdelconsumer",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XGroupDelConsumer(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXREADGROUP(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xreadgroup",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("Streams", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Block", types.Number{}), types.NewStaticProperty("NoAck", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XReadGroupArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XReadGroup(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp74768 []*ast.Term
				for _, v := range r0 {

					termStream := ast.StringTerm(v.Stream)

					var tmp42960 []*ast.Term
					for _, v := range v.Messages {

						termID := ast.StringTerm(v.ID)

						var tmp88231 [][2]*ast.Term
						for key, value := range v.Values {
							k := ast.StringTerm(key)

							v := ast.NullTerm()
							if s, ok := value.(string); ok {
								v = ast.StringTerm(s)
							}

							tmp88231 = append(tmp88231, [2]*ast.Term{k, v})
						}
						termValues := ast.ObjectTerm(tmp88231...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

						tmp42960 = append(tmp42960, term)
					}
					termMessages := ast.ArrayTerm(tmp42960...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Stream"), termStream}, [2]*ast.Term{ast.StringTerm("Messages"), termMessages})

					tmp74768 = append(tmp74768, term)
				}
				term := ast.ArrayTerm(tmp74768...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXACK(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xack",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XAck(p.redisContext, v0, v1, v2...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXPENDING(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xpending",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Lower", types.String{}), types.NewStaticProperty("Higher", types.String{}), types.NewStaticProperty("Consumers", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Number{})))}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XPending(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termCount := ast.IntNumberTerm(int(r0.Count))

				termLower := ast.StringTerm(r0.Lower)

				termHigher := ast.StringTerm(r0.Higher)

				var tmp25995 [][2]*ast.Term
				for key, value := range r0.Consumers {
					k := ast.StringTerm(key)
					v := ast.IntNumberTerm(int(value))
					tmp25995 = append(tmp25995, [2]*ast.Term{k, v})
				}
				termConsumers := ast.ObjectTerm(tmp25995...)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Count"), termCount}, [2]*ast.Term{ast.StringTerm("Lower"), termLower}, [2]*ast.Term{ast.StringTerm("Higher"), termHigher}, [2]*ast.Term{ast.StringTerm("Consumers"), termConsumers})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXPENDINGEXT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xpendingext",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Idle", types.Number{}), types.NewStaticProperty("Start", types.String{}), types.NewStaticProperty("End", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("Idle", types.Number{}), types.NewStaticProperty("RetryCount", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XPendingExtArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XPendingExt(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp81039 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					termConsumer := ast.StringTerm(v.Consumer)

					termIdle := ast.IntNumberTerm(int(v.Idle.Milliseconds()))

					termRetryCount := ast.IntNumberTerm(int(v.RetryCount))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Consumer"), termConsumer}, [2]*ast.Term{ast.StringTerm("Idle"), termIdle}, [2]*ast.Term{ast.StringTerm("RetryCount"), termRetryCount})

					tmp81039 = append(tmp81039, term)
				}
				term := ast.ArrayTerm(tmp81039...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXCLAIM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xclaim",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.String{}))}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XClaim(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp62750 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp84422 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp84422 = append(tmp84422, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp84422...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp62750 = append(tmp62750, term)
				}
				term := ast.ArrayTerm(tmp62750...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXCLAIMJUSTID(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xclaimjustid",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.String{}))}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XClaimJustID(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp76638 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp76638 = append(tmp76638, term)
				}
				term := ast.ArrayTerm(tmp76638...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXAUTOCLAIM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xautoclaim",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Start", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)), types.String{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XAutoClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XAutoClaim(p.redisContext, v0)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp50671 []*ast.Term
				for _, v := range r0 {

					termID := ast.StringTerm(v.ID)

					var tmp00608 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp00608 = append(tmp00608, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp00608...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp50671 = append(tmp50671, term)
				}
				tr0 := ast.ArrayTerm(tmp50671...)
				tr1 := ast.StringTerm(r1)
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerXAUTOCLAIMJUSTID(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xautoclaimjustid",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}), types.NewStaticProperty("Group", types.String{}), types.NewStaticProperty("MinIdle", types.Number{}), types.NewStaticProperty("Start", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}), types.String{}}, types.Null{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.XAutoClaimArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XAutoClaimJustID(p.redisContext, v0)
			r0, r1 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp52238 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp52238 = append(tmp52238, term)
				}
				tr0 := ast.ArrayTerm(tmp52238...)
				tr1 := ast.StringTerm(r1)
				return ast.ArrayTerm(tr0, tr1), nil

			default:
				return nil, err
			}
		},
	)
}

func registerXTRIM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrim",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrim(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMAPPROX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimapprox",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrimApprox(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMAXLEN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimmaxlen",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrimMaxLen(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMAXLENAPPROX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimmaxlenapprox",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XTrimMaxLenApprox(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMINID(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimminid",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XTrimMinID(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXTRIMMINIDAPPROX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xtrimminidapprox",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.XTrimMinIDApprox(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerXINFOGROUPS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfogroups",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Consumers", types.Number{}), types.NewStaticProperty("Pending", types.Number{}), types.NewStaticProperty("LastDeliveredID", types.String{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XInfoGroups(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp70406 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termConsumers := ast.IntNumberTerm(int(v.Consumers))

					termPending := ast.IntNumberTerm(int(v.Pending))

					termLastDeliveredID := ast.StringTerm(v.LastDeliveredID)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Consumers"), termConsumers}, [2]*ast.Term{ast.StringTerm("Pending"), termPending}, [2]*ast.Term{ast.StringTerm("LastDeliveredID"), termLastDeliveredID})

					tmp70406 = append(tmp70406, term)
				}
				term := ast.ArrayTerm(tmp70406...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXINFOSTREAM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfostream",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Length", types.Number{}), types.NewStaticProperty("RadixTreeKeys", types.Number{}), types.NewStaticProperty("RadixTreeNodes", types.Number{}), types.NewStaticProperty("Groups", types.Number{}), types.NewStaticProperty("LastGeneratedID", types.String{}), types.NewStaticProperty("FirstEntry", types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil)), types.NewStaticProperty("LastEntry", types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.XInfoStream(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termLength := ast.IntNumberTerm(int(r0.Length))

				termRadixTreeKeys := ast.IntNumberTerm(int(r0.RadixTreeKeys))

				termRadixTreeNodes := ast.IntNumberTerm(int(r0.RadixTreeNodes))

				termGroups := ast.IntNumberTerm(int(r0.Groups))

				termLastGeneratedID := ast.StringTerm(r0.LastGeneratedID)

				termFirstEntryID := ast.StringTerm(r0.FirstEntry.ID)

				var tmp55850 [][2]*ast.Term
				for key, value := range r0.FirstEntry.Values {
					k := ast.StringTerm(key)

					v := ast.NullTerm()
					if s, ok := value.(string); ok {
						v = ast.StringTerm(s)
					}

					tmp55850 = append(tmp55850, [2]*ast.Term{k, v})
				}
				termFirstEntryValues := ast.ObjectTerm(tmp55850...)

				termFirstEntry := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termFirstEntryID}, [2]*ast.Term{ast.StringTerm("Values"), termFirstEntryValues})

				termLastEntryID := ast.StringTerm(r0.LastEntry.ID)

				var tmp03180 [][2]*ast.Term
				for key, value := range r0.LastEntry.Values {
					k := ast.StringTerm(key)

					v := ast.NullTerm()
					if s, ok := value.(string); ok {
						v = ast.StringTerm(s)
					}

					tmp03180 = append(tmp03180, [2]*ast.Term{k, v})
				}
				termLastEntryValues := ast.ObjectTerm(tmp03180...)

				termLastEntry := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termLastEntryID}, [2]*ast.Term{ast.StringTerm("Values"), termLastEntryValues})

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Length"), termLength}, [2]*ast.Term{ast.StringTerm("RadixTreeKeys"), termRadixTreeKeys}, [2]*ast.Term{ast.StringTerm("RadixTreeNodes"), termRadixTreeNodes}, [2]*ast.Term{ast.StringTerm("Groups"), termGroups}, [2]*ast.Term{ast.StringTerm("LastGeneratedID"), termLastGeneratedID}, [2]*ast.Term{ast.StringTerm("FirstEntry"), termFirstEntry}, [2]*ast.Term{ast.StringTerm("LastEntry"), termLastEntry})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXINFOSTREAMFULL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfostreamfull",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Length", types.Number{}), types.NewStaticProperty("RadixTreeKeys", types.Number{}), types.NewStaticProperty("RadixTreeNodes", types.Number{}), types.NewStaticProperty("LastGeneratedID", types.String{}), types.NewStaticProperty("Entries", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Values", types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Any{})))}, nil))), types.NewStaticProperty("Groups", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("LastDeliveredID", types.String{}), types.NewStaticProperty("PelCount", types.Number{}), types.NewStaticProperty("Pending", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Consumer", types.String{}), types.NewStaticProperty("DeliveryTime", types.Number{}), types.NewStaticProperty("DeliveryCount", types.Number{})}, nil))), types.NewStaticProperty("Consumers", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("SeenTime", types.Number{}), types.NewStaticProperty("PelCount", types.Number{}), types.NewStaticProperty("Pending", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("DeliveryTime", types.Number{}), types.NewStaticProperty("DeliveryCount", types.Number{})}, nil)))}, nil)))}, nil)))}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XInfoStreamFull(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termLength := ast.IntNumberTerm(int(r0.Length))

				termRadixTreeKeys := ast.IntNumberTerm(int(r0.RadixTreeKeys))

				termRadixTreeNodes := ast.IntNumberTerm(int(r0.RadixTreeNodes))

				termLastGeneratedID := ast.StringTerm(r0.LastGeneratedID)

				var tmp22800 []*ast.Term
				for _, v := range r0.Entries {

					termID := ast.StringTerm(v.ID)

					var tmp69734 [][2]*ast.Term
					for key, value := range v.Values {
						k := ast.StringTerm(key)

						v := ast.NullTerm()
						if s, ok := value.(string); ok {
							v = ast.StringTerm(s)
						}

						tmp69734 = append(tmp69734, [2]*ast.Term{k, v})
					}
					termValues := ast.ObjectTerm(tmp69734...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Values"), termValues})

					tmp22800 = append(tmp22800, term)
				}
				termEntries := ast.ArrayTerm(tmp22800...)

				var tmp88663 []*ast.Term
				for _, v := range r0.Groups {

					termName := ast.StringTerm(v.Name)

					termLastDeliveredID := ast.StringTerm(v.LastDeliveredID)

					termPelCount := ast.IntNumberTerm(int(v.PelCount))

					var tmp55091 []*ast.Term
					for _, v := range v.Pending {

						termID := ast.StringTerm(v.ID)

						termConsumer := ast.StringTerm(v.Consumer)

						termDeliveryTime := ast.IntNumberTerm(int(v.DeliveryTime.UnixMicro()))

						termDeliveryCount := ast.IntNumberTerm(int(v.DeliveryCount))

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Consumer"), termConsumer}, [2]*ast.Term{ast.StringTerm("DeliveryTime"), termDeliveryTime}, [2]*ast.Term{ast.StringTerm("DeliveryCount"), termDeliveryCount})

						tmp55091 = append(tmp55091, term)
					}
					termPending := ast.ArrayTerm(tmp55091...)

					var tmp42084 []*ast.Term
					for _, v := range v.Consumers {

						termName := ast.StringTerm(v.Name)

						termSeenTime := ast.IntNumberTerm(int(v.SeenTime.UnixMicro()))

						termPelCount := ast.IntNumberTerm(int(v.PelCount))

						var tmp84282 []*ast.Term
						for _, v := range v.Pending {

							termID := ast.StringTerm(v.ID)

							termDeliveryTime := ast.IntNumberTerm(int(v.DeliveryTime.UnixMicro()))

							termDeliveryCount := ast.IntNumberTerm(int(v.DeliveryCount))

							term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("DeliveryTime"), termDeliveryTime}, [2]*ast.Term{ast.StringTerm("DeliveryCount"), termDeliveryCount})

							tmp84282 = append(tmp84282, term)
						}
						termPending := ast.ArrayTerm(tmp84282...)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("SeenTime"), termSeenTime}, [2]*ast.Term{ast.StringTerm("PelCount"), termPelCount}, [2]*ast.Term{ast.StringTerm("Pending"), termPending})

						tmp42084 = append(tmp42084, term)
					}
					termConsumers := ast.ArrayTerm(tmp42084...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("LastDeliveredID"), termLastDeliveredID}, [2]*ast.Term{ast.StringTerm("PelCount"), termPelCount}, [2]*ast.Term{ast.StringTerm("Pending"), termPending}, [2]*ast.Term{ast.StringTerm("Consumers"), termConsumers})

					tmp88663 = append(tmp88663, term)
				}
				termGroups := ast.ArrayTerm(tmp88663...)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Length"), termLength}, [2]*ast.Term{ast.StringTerm("RadixTreeKeys"), termRadixTreeKeys}, [2]*ast.Term{ast.StringTerm("RadixTreeNodes"), termRadixTreeNodes}, [2]*ast.Term{ast.StringTerm("LastGeneratedID"), termLastGeneratedID}, [2]*ast.Term{ast.StringTerm("Entries"), termEntries}, [2]*ast.Term{ast.StringTerm("Groups"), termGroups})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerXINFOCONSUMERS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.xinfoconsumers",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Pending", types.Number{}), types.NewStaticProperty("Idle", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.XInfoConsumers(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp91807 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termPending := ast.IntNumberTerm(int(v.Pending))

					termIdle := ast.IntNumberTerm(int(v.Idle))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Pending"), termPending}, [2]*ast.Term{ast.StringTerm("Idle"), termIdle})

					tmp91807 = append(tmp91807, term)
				}
				term := ast.ArrayTerm(tmp91807...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBZPOPMAX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bzpopmax",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{})}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BZPopMax(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termKey := ast.StringTerm(r0.Key)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Key"), termKey})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerBZPOPMIN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bzpopmin",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{})}, nil)),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.BZPopMin(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				termKey := ast.StringTerm(r0.Key)

				term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Key"), termKey})

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZADD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAdd(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDNX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddNX(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDXX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddxx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddXX(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDCH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddCh(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDNXCH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddnxch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddNXCh(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDXXCH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddxxch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddXXCh(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDARGS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddargs",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("NX", types.Boolean{}), types.NewStaticProperty("XX", types.Boolean{}), types.NewStaticProperty("LT", types.Boolean{}), types.NewStaticProperty("GT", types.Boolean{}), types.NewStaticProperty("Ch", types.Boolean{}), types.NewStaticProperty("Members", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)))}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 redis.ZAddArgs
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddArgs(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZADDARGSINCR(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zaddargsincr",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("NX", types.Boolean{}), types.NewStaticProperty("XX", types.Boolean{}), types.NewStaticProperty("LT", types.Boolean{}), types.NewStaticProperty("GT", types.Boolean{}), types.NewStaticProperty("Ch", types.Boolean{}), types.NewStaticProperty("Members", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)))}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 redis.ZAddArgs
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZAddArgsIncr(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCR(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincr",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZIncr(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCRNX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincrnx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZIncrNX(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCRXX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincrxx",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.Z
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZIncrXX(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZCARD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zcard",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZCard(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZCount(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZLEXCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zlexcount",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZLexCount(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINCRBY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zincrby",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZIncrBy(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZINTER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zinter",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZInter(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp16789 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp16789 = append(tmp16789, term)
				}
				term := ast.ArrayTerm(tmp16789...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZINTERWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zinterwithscores",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 *redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZInterWithScores(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp17315 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp17315 = append(tmp17315, term)
				}
				term := ast.ArrayTerm(tmp17315...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZINTERSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zinterstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZStore
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZInterStore(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZMSCORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zmscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Number{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZMScore(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp80844 []*ast.Term
				for _, v := range r0 {
					term := ast.FloatNumberTerm(float64(v))
					tmp80844 = append(tmp80844, term)
				}
				term := ast.ArrayTerm(tmp80844...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZPOPMAX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zpopmax",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Number{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZPopMax(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp69420 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp69420 = append(tmp69420, term)
				}
				term := ast.ArrayTerm(tmp69420...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZPOPMIN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zpopmin",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Number{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZPopMin(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp35901 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp35901 = append(tmp35901, term)
				}
				term := ast.ArrayTerm(tmp35901...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp87389 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp87389 = append(tmp87389, term)
				}
				term := ast.ArrayTerm(tmp87389...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRangeWithScores(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp28575 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp28575 = append(tmp28575, term)
				}
				term := ast.ArrayTerm(tmp28575...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEBYSCORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangebyscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeByScore(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp33480 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp33480 = append(tmp33480, term)
				}
				term := ast.ArrayTerm(tmp33480...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEBYLEX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangebylex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeByLex(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp31686 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp31686 = append(tmp31686, term)
				}
				term := ast.ArrayTerm(tmp31686...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEBYSCOREWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangebyscorewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeByScoreWithScores(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp83824 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp83824 = append(tmp83824, term)
				}
				term := ast.ArrayTerm(tmp83824...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEARGS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangeargs",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}), types.NewStaticProperty("Start", types.Any{}), types.NewStaticProperty("Stop", types.Any{}), types.NewStaticProperty("ByScore", types.Boolean{}), types.NewStaticProperty("ByLex", types.Boolean{}), types.NewStaticProperty("Rev", types.Boolean{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZRangeArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZRangeArgs(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp89959 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp89959 = append(tmp89959, term)
				}
				term := ast.ArrayTerm(tmp89959...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGEARGSWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangeargswithscores",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}), types.NewStaticProperty("Start", types.Any{}), types.NewStaticProperty("Stop", types.Any{}), types.NewStaticProperty("ByScore", types.Boolean{}), types.NewStaticProperty("ByLex", types.Boolean{}), types.NewStaticProperty("Rev", types.Boolean{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZRangeArgs
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZRangeArgsWithScores(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp39550 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp39550 = append(tmp39550, term)
				}
				term := ast.ArrayTerm(tmp39550...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANGESTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrangestore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}), types.NewStaticProperty("Start", types.Any{}), types.NewStaticProperty("Stop", types.Any{}), types.NewStaticProperty("ByScore", types.Boolean{}), types.NewStaticProperty("ByLex", types.Boolean{}), types.NewStaticProperty("Rev", types.Boolean{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 redis.ZRangeArgs
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRangeStore(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZRANK(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrank",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRank(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREM(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrem",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRem(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREMRANGEBYRANK(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zremrangebyrank",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRemRangeByRank(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREMRANGEBYSCORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zremrangebyscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRemRangeByScore(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREMRANGEBYLEX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zremrangebylex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRemRangeByLex(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrange",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRevRange(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp12917 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp12917 = append(tmp12917, term)
				}
				term := ast.ArrayTerm(tmp12917...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 int64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeWithScores(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp84482 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp84482 = append(tmp84482, term)
				}
				term := ast.ArrayTerm(tmp84482...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEBYSCORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangebyscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeByScore(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp18351 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp18351 = append(tmp18351, term)
				}
				term := ast.ArrayTerm(tmp18351...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEBYLEX(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangebylex",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeByLex(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp60402 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp60402 = append(tmp60402, term)
				}
				term := ast.ArrayTerm(tmp60402...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANGEBYSCOREWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrangebyscorewithscores",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}), types.NewStaticProperty("Max", types.String{}), types.NewStaticProperty("Offset", types.Number{}), types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZRangeBy
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRangeByScoreWithScores(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp78305 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp78305 = append(tmp78305, term)
				}
				term := ast.ArrayTerm(tmp78305...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZREVRANK(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrevrank",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZRevRank(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZSCORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zscore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZScore(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZUNIONSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zunionstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.ZStore
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZUnionStore(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerZUNION(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zunion",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZUnion(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp18108 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp18108 = append(tmp18108, term)
				}
				term := ast.ArrayTerm(tmp18108...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZUNIONWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zunionwithscores",
			Decl:             types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})), types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})), types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 redis.ZStore
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZUnionWithScores(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp11280 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp11280 = append(tmp11280, term)
				}
				term := ast.ArrayTerm(tmp11280...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZRANDMEMBER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zrandmember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Boolean{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 bool
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.ZRandMember(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp73558 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp73558 = append(tmp73558, term)
				}
				term := ast.ArrayTerm(tmp73558...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZDIFF(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zdiff",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZDiff(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp10531 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp10531 = append(tmp10531, term)
				}
				term := ast.ArrayTerm(tmp10531...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZDIFFWITHSCORES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zdiffwithscores",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}), types.NewStaticProperty("Member", types.Any{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ZDiffWithScores(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp63619 []*ast.Term
				for _, v := range r0 {

					termScore := ast.FloatNumberTerm(float64(v.Score))

					termMember := ast.NullTerm()
					if s, ok := v.Member.(string); ok {
						termMember = ast.StringTerm(s)
					}

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), termScore}, [2]*ast.Term{ast.StringTerm("Member"), termMember})

					tmp63619 = append(tmp63619, term)
				}
				term := ast.ArrayTerm(tmp63619...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerZDIFFSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.zdiffstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ZDiffStore(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPFADD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pfadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PFAdd(p.redisContext, v0, utils.Conva(v1)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPFCOUNT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pfcount",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PFCount(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPFMERGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pfmerge",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.PFMerge(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBGREWRITEAOF(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bgrewriteaof",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.BgRewriteAOF(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerBGSAVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.bgsave",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.BgSave(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTKILL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientkill",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClientKill(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTKILLBYFILTER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientkillbyfilter",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClientKillByFilter(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTLIST(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientlist",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClientList(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTPAUSE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientpause",
			Decl:             types.NewFunction(types.Args(types.Number{}), types.Boolean{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 time.Duration
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClientPause(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.BooleanTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLIENTID(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clientid",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClientID(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGGET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configget",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.Any{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ConfigGet(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp47809 []*ast.Term
				for _, v := range r0 {

					term := ast.NullTerm()
					if s, ok := v.(string); ok {
						term = ast.StringTerm(s)
					}

					tmp47809 = append(tmp47809, term)
				}
				term := ast.ArrayTerm(tmp47809...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGRESETSTAT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configresetstat",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ConfigResetStat(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGSET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configset",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ConfigSet(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCONFIGREWRITE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.configrewrite",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ConfigRewrite(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDBSIZE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.dbsize",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.DBSize(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHALL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushall",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushAll(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHALLASYNC(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushallasync",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushAllAsync(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHDB(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushdb",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushDB(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerFLUSHDBASYNC(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.flushdbasync",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.FlushDBAsync(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerINFO(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.info",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.Info(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerLASTSAVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.lastsave",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.LastSave(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSAVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.save",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Save(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSHUTDOWN(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.shutdown",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Shutdown(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSHUTDOWNSAVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.shutdownsave",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ShutdownSave(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSHUTDOWNNOSAVE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.shutdownnosave",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ShutdownNoSave(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSLAVEOF(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.slaveof",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.SlaveOf(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerTIME(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.time",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.Time(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0.UnixMicro()))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerDEBUGOBJECT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.debugobject",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.DebugObject(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerREADONLY(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.readonly",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ReadOnly(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerREADWRITE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.readwrite",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ReadWrite(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerMEMORYUSAGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.memoryusage",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.Number{})), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.MemoryUsage(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEVAL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.eval",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{}), types.NewArray([]types.Type{}, types.Any{})), types.Any{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.Eval(p.redisContext, v0, v1, utils.Conva(v2)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.NullTerm()
				if s, ok := r0.(string); ok {
					term = ast.StringTerm(s)
				}

				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerEVALSHA(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.evalsha",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{}), types.NewArray([]types.Type{}, types.Any{})), types.Any{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 []interface{}
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.EvalSha(p.redisContext, v0, v1, utils.Conva(v2)...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.NullTerm()
				if s, ok := r0.(string); ok {
					term = ast.StringTerm(s)
				}

				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTEXISTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptexists",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Boolean{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ScriptExists(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp76967 []*ast.Term
				for _, v := range r0 {
					term := ast.BooleanTerm(v)
					tmp76967 = append(tmp76967, term)
				}
				term := ast.ArrayTerm(tmp76967...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTFLUSH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptflush",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ScriptFlush(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTKILL(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptkill",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ScriptKill(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerSCRIPTLOAD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.scriptload",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ScriptLoad(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPUBLISH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.publish",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Any{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 interface{}
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.Publish(p.redisContext, v0, utils.Conv(v1))
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerPUBSUBCHANNELS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pubsubchannels",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PubSubChannels(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp93087 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp93087 = append(tmp93087, term)
				}
				term := ast.ArrayTerm(tmp93087...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerPUBSUBNUMSUB(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pubsubnumsub",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{}, types.NewDynamicProperty(types.String{}, types.Number{}))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.PubSubNumSub(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp88613 [][2]*ast.Term
				for key, value := range r0 {
					k := ast.StringTerm(key)
					v := ast.IntNumberTerm(int(value))
					tmp88613 = append(tmp88613, [2]*ast.Term{k, v})
				}
				term := ast.ObjectTerm(tmp88613...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerPUBSUBNUMPAT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.pubsubnumpat",
			Decl:             types.NewFunction(types.Args(), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.PubSubNumPat(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERSLOTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterslots",
			Decl:             types.NewFunction(types.Args(), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Start", types.Number{}), types.NewStaticProperty("End", types.Number{}), types.NewStaticProperty("Nodes", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}), types.NewStaticProperty("Addr", types.String{})}, nil)))}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterSlots(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp62128 []*ast.Term
				for _, v := range r0 {

					termStart := ast.IntNumberTerm(int(v.Start))

					termEnd := ast.IntNumberTerm(int(v.End))

					var tmp89234 []*ast.Term
					for _, v := range v.Nodes {

						termID := ast.StringTerm(v.ID)

						termAddr := ast.StringTerm(v.Addr)

						term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), termID}, [2]*ast.Term{ast.StringTerm("Addr"), termAddr})

						tmp89234 = append(tmp89234, term)
					}
					termNodes := ast.ArrayTerm(tmp89234...)

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Start"), termStart}, [2]*ast.Term{ast.StringTerm("End"), termEnd}, [2]*ast.Term{ast.StringTerm("Nodes"), termNodes})

					tmp62128 = append(tmp62128, term)
				}
				term := ast.ArrayTerm(tmp62128...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERNODES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusternodes",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterNodes(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERMEET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustermeet",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterMeet(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERFORGET(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterforget",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterForget(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERREPLICATE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterreplicate",
			Decl:             types.NewFunction(types.Args(types.String{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterReplicate(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERRESETSOFT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterresetsoft",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterResetSoft(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERRESETHARD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterresethard",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterResetHard(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERINFO(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterinfo",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterInfo(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERKEYSLOT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterkeyslot",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterKeySlot(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERGETKEYSINSLOT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustergetkeysinslot",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.Number{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterGetKeysInSlot(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp87325 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp87325 = append(tmp87325, term)
				}
				term := ast.ArrayTerm(tmp87325...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERCOUNTFAILUREREPORTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustercountfailurereports",
			Decl:             types.NewFunction(types.Args(types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterCountFailureReports(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERCOUNTKEYSINSLOT(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustercountkeysinslot",
			Decl:             types.NewFunction(types.Args(types.Number{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterCountKeysInSlot(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERDELSLOTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterdelslots",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Number{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterDelSlots(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERDELSLOTSRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterdelslotsrange",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterDelSlotsRange(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERSAVECONFIG(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clustersaveconfig",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterSaveConfig(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERSLAVES(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterslaves",
			Decl:             types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterSlaves(p.redisContext, v0)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp96409 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp96409 = append(tmp96409, term)
				}
				term := ast.ArrayTerm(tmp96409...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERFAILOVER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusterfailover",
			Decl:             types.NewFunction(types.Args(), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			val := rdb.ClusterFailover(p.redisContext)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERADDSLOTS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusteraddslots",
			Decl:             types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Number{})), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 []int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			val := rdb.ClusterAddSlots(p.redisContext, v0...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerCLUSTERADDSLOTSRANGE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.clusteraddslotsrange",
			Decl:             types.NewFunction(types.Args(types.Number{}, types.Number{}), types.String{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 int
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 int
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.ClusterAddSlotsRange(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.StringTerm(r0)
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOADD(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geoadd",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []*redis.GeoLocation
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoAdd(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOPOS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geopos",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoPos(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp17311 []*ast.Term
				for _, v := range r0 {
					if v == nil {
						tmp17311 = append(tmp17311, ast.NullTerm())
						continue
					}

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude})

					tmp17311 = append(tmp17311, term)
				}
				term := ast.ArrayTerm(tmp17311...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUS(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadius",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 float64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 *redis.GeoRadiusQuery
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.GeoRadius(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp39626 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					termDist := ast.FloatNumberTerm(float64(v.Dist))

					termGeoHash := ast.IntNumberTerm(int(v.GeoHash))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude}, [2]*ast.Term{ast.StringTerm("Dist"), termDist}, [2]*ast.Term{ast.StringTerm("GeoHash"), termGeoHash})

					tmp39626 = append(tmp39626, term)
				}
				term := ast.ArrayTerm(tmp39626...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUSSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadiusstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.Number{}, types.Number{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 float64
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 float64
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 *redis.GeoRadiusQuery
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.GeoRadiusStore(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUSBYMEMBER(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadiusbymember",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.GeoRadiusQuery
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GeoRadiusByMember(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp99705 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					termDist := ast.FloatNumberTerm(float64(v.Dist))

					termGeoHash := ast.IntNumberTerm(int(v.GeoHash))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude}, [2]*ast.Term{ast.StringTerm("Dist"), termDist}, [2]*ast.Term{ast.StringTerm("GeoHash"), termGeoHash})

					tmp99705 = append(tmp99705, term)
				}
				term := ast.ArrayTerm(tmp99705...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEORADIUSBYMEMBERSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.georadiusbymemberstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("Unit", types.String{}), types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithGeoHash", types.Boolean{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Store", types.String{}), types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.GeoRadiusQuery
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GeoRadiusByMemberStore(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOSEARCH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geosearch",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Member", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Radius", types.Number{}), types.NewStaticProperty("RadiusUnit", types.String{}), types.NewStaticProperty("BoxWidth", types.Number{}), types.NewStaticProperty("BoxHeight", types.Number{}), types.NewStaticProperty("BoxUnit", types.String{}), types.NewStaticProperty("Sort", types.String{}), types.NewStaticProperty("Count", types.Number{}), types.NewStaticProperty("CountAny", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.GeoSearchQuery
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoSearch(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp47393 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp47393 = append(tmp47393, term)
				}
				term := ast.ArrayTerm(tmp47393...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEOSEARCHLOCATION(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geosearchlocation",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("WithCoord", types.Boolean{}), types.NewStaticProperty("WithDist", types.Boolean{}), types.NewStaticProperty("WithHash", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}), types.NewStaticProperty("Longitude", types.Number{}), types.NewStaticProperty("Latitude", types.Number{}), types.NewStaticProperty("Dist", types.Number{}), types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 *redis.GeoSearchLocationQuery
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoSearchLocation(p.redisContext, v0, v1)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp36292 []*ast.Term
				for _, v := range r0 {

					termName := ast.StringTerm(v.Name)

					termLongitude := ast.FloatNumberTerm(float64(v.Longitude))

					termLatitude := ast.FloatNumberTerm(float64(v.Latitude))

					termDist := ast.FloatNumberTerm(float64(v.Dist))

					termGeoHash := ast.IntNumberTerm(int(v.GeoHash))

					term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), termName}, [2]*ast.Term{ast.StringTerm("Longitude"), termLongitude}, [2]*ast.Term{ast.StringTerm("Latitude"), termLatitude}, [2]*ast.Term{ast.StringTerm("Dist"), termDist}, [2]*ast.Term{ast.StringTerm("GeoHash"), termGeoHash})

					tmp36292 = append(tmp36292, term)
				}
				term := ast.ArrayTerm(tmp36292...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func registerGEOSEARCHSTORE(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geosearchstore",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("StoreDist", types.Boolean{})}, nil)), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 *redis.GeoSearchStoreQuery
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			val := rdb.GeoSearchStore(p.redisContext, v0, v1, v2)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.IntNumberTerm(int(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEODIST(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geodist",
			Decl:             types.NewFunction(types.Args(types.String{}, types.String{}, types.String{}, types.String{}), types.Number{}),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			var v2 string
			if err := ast.As(terms[2].Value, &v2); err != nil {
				return nil, err
			}

			var v3 string
			if err := ast.As(terms[3].Value, &v3); err != nil {
				return nil, err
			}

			val := rdb.GeoDist(p.redisContext, v0, v1, v2, v3)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				term := ast.FloatNumberTerm(float64(r0))
				return term, nil
			default:
				return nil, err
			}
		},
	)
}

func registerGEOHASH(p *redisPlugin) {
	rego.RegisterBuiltinDyn(
		&rego.Function{
			Name:             "redis.geohash",
			Decl:             types.NewFunction(types.Args(types.String{}, types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
			Memoize:          true,
			Nondeterministic: true,
		},
		func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
			rdb, err := p.redisProxy.Get()
			if err != nil {
				return nil, err
			}

			var v0 string
			if err := ast.As(terms[0].Value, &v0); err != nil {
				return nil, err
			}

			var v1 []string
			if err := ast.As(terms[1].Value, &v1); err != nil {
				return nil, err
			}

			val := rdb.GeoHash(p.redisContext, v0, v1...)
			r0 := val.Val()
			switch err := val.Err(); err {
			case redis.Nil:
				return ast.NullTerm(), nil
			case nil:

				var tmp03113 []*ast.Term
				for _, v := range r0 {
					term := ast.StringTerm(v)
					tmp03113 = append(tmp03113, term)
				}
				term := ast.ArrayTerm(tmp03113...)

				return term, nil

			default:
				return nil, err
			}
		},
	)
}

func (p *redisPlugin) registerAutogenCommands() {
	registerCLIENTGETNAME(p)
	registerECHO(p)
	registerPING(p)
	registerQUIT(p)
	registerDEL(p)
	registerUNLINK(p)
	registerDUMP(p)
	registerEXISTS(p)
	registerEXPIRE(p)
	registerEXPIREAT(p)
	registerEXPIRENX(p)
	registerEXPIREXX(p)
	registerEXPIREGT(p)
	registerEXPIRELT(p)
	registerKEYS(p)
	registerMIGRATE(p)
	registerMOVE(p)
	registerOBJECTREFCOUNT(p)
	registerOBJECTENCODING(p)
	registerOBJECTIDLETIME(p)
	registerPERSIST(p)
	registerPEXPIRE(p)
	registerPEXPIREAT(p)
	registerPTTL(p)
	registerRANDOMKEY(p)
	registerRENAME(p)
	registerRENAMENX(p)
	registerRESTORE(p)
	registerRESTOREREPLACE(p)
	registerSORT(p)
	registerSORTSTORE(p)
	registerSORTINTERFACES(p)
	registerTOUCH(p)
	registerTTL(p)
	registerTYPE(p)
	registerAPPEND(p)
	registerDECR(p)
	registerDECRBY(p)
	registerGET(p)
	registerGETRANGE(p)
	registerGETSET(p)
	registerGETEX(p)
	registerGETDEL(p)
	registerINCR(p)
	registerINCRBY(p)
	registerINCRBYFLOAT(p)
	registerMGET(p)
	registerMSET(p)
	registerMSETNX(p)
	registerSET(p)
	registerSETARGS(p)
	registerSETEX(p)
	registerSETNX(p)
	registerSETXX(p)
	registerSETRANGE(p)
	registerSTRLEN(p)
	registerCOPY(p)
	registerGETBIT(p)
	registerSETBIT(p)
	registerBITCOUNT(p)
	registerBITOPAND(p)
	registerBITOPOR(p)
	registerBITOPXOR(p)
	registerBITOPNOT(p)
	registerBITPOS(p)
	registerBITFIELD(p)
	registerSCAN(p)
	registerSCANTYPE(p)
	registerSSCAN(p)
	registerHSCAN(p)
	registerZSCAN(p)
	registerHDEL(p)
	registerHEXISTS(p)
	registerHGET(p)
	registerHGETALL(p)
	registerHINCRBY(p)
	registerHINCRBYFLOAT(p)
	registerHKEYS(p)
	registerHLEN(p)
	registerHMGET(p)
	registerHSET(p)
	registerHMSET(p)
	registerHSETNX(p)
	registerHVALS(p)
	registerHRANDFIELD(p)
	registerBLPOP(p)
	registerBRPOP(p)
	registerBRPOPLPUSH(p)
	registerLINDEX(p)
	registerLINSERT(p)
	registerLINSERTBEFORE(p)
	registerLINSERTAFTER(p)
	registerLLEN(p)
	registerLPOP(p)
	registerLPOPCOUNT(p)
	registerLPOS(p)
	registerLPOSCOUNT(p)
	registerLPUSH(p)
	registerLPUSHX(p)
	registerLRANGE(p)
	registerLREM(p)
	registerLSET(p)
	registerLTRIM(p)
	registerRPOP(p)
	registerRPOPCOUNT(p)
	registerRPOPLPUSH(p)
	registerRPUSH(p)
	registerRPUSHX(p)
	registerLMOVE(p)
	registerBLMOVE(p)
	registerSADD(p)
	registerSCARD(p)
	registerSDIFF(p)
	registerSDIFFSTORE(p)
	registerSINTER(p)
	registerSINTERSTORE(p)
	registerSISMEMBER(p)
	registerSMISMEMBER(p)
	registerSMEMBERS(p)
	registerSMEMBERSMAP(p)
	registerSMOVE(p)
	registerSPOP(p)
	registerSPOPN(p)
	registerSRANDMEMBER(p)
	registerSRANDMEMBERN(p)
	registerSREM(p)
	registerSUNION(p)
	registerSUNIONSTORE(p)
	registerXADD(p)
	registerXDEL(p)
	registerXLEN(p)
	registerXRANGE(p)
	registerXRANGEN(p)
	registerXREVRANGE(p)
	registerXREVRANGEN(p)
	registerXREAD(p)
	registerXREADSTREAMS(p)
	registerXGROUPCREATE(p)
	registerXGROUPCREATEMKSTREAM(p)
	registerXGROUPSETID(p)
	registerXGROUPDESTROY(p)
	registerXGROUPCREATECONSUMER(p)
	registerXGROUPDELCONSUMER(p)
	registerXREADGROUP(p)
	registerXACK(p)
	registerXPENDING(p)
	registerXPENDINGEXT(p)
	registerXCLAIM(p)
	registerXCLAIMJUSTID(p)
	registerXAUTOCLAIM(p)
	registerXAUTOCLAIMJUSTID(p)
	registerXTRIM(p)
	registerXTRIMAPPROX(p)
	registerXTRIMMAXLEN(p)
	registerXTRIMMAXLENAPPROX(p)
	registerXTRIMMINID(p)
	registerXTRIMMINIDAPPROX(p)
	registerXINFOGROUPS(p)
	registerXINFOSTREAM(p)
	registerXINFOSTREAMFULL(p)
	registerXINFOCONSUMERS(p)
	registerBZPOPMAX(p)
	registerBZPOPMIN(p)
	registerZADD(p)
	registerZADDNX(p)
	registerZADDXX(p)
	registerZADDCH(p)
	registerZADDNXCH(p)
	registerZADDXXCH(p)
	registerZADDARGS(p)
	registerZADDARGSINCR(p)
	registerZINCR(p)
	registerZINCRNX(p)
	registerZINCRXX(p)
	registerZCARD(p)
	registerZCOUNT(p)
	registerZLEXCOUNT(p)
	registerZINCRBY(p)
	registerZINTER(p)
	registerZINTERWITHSCORES(p)
	registerZINTERSTORE(p)
	registerZMSCORE(p)
	registerZPOPMAX(p)
	registerZPOPMIN(p)
	registerZRANGE(p)
	registerZRANGEWITHSCORES(p)
	registerZRANGEBYSCORE(p)
	registerZRANGEBYLEX(p)
	registerZRANGEBYSCOREWITHSCORES(p)
	registerZRANGEARGS(p)
	registerZRANGEARGSWITHSCORES(p)
	registerZRANGESTORE(p)
	registerZRANK(p)
	registerZREM(p)
	registerZREMRANGEBYRANK(p)
	registerZREMRANGEBYSCORE(p)
	registerZREMRANGEBYLEX(p)
	registerZREVRANGE(p)
	registerZREVRANGEWITHSCORES(p)
	registerZREVRANGEBYSCORE(p)
	registerZREVRANGEBYLEX(p)
	registerZREVRANGEBYSCOREWITHSCORES(p)
	registerZREVRANK(p)
	registerZSCORE(p)
	registerZUNIONSTORE(p)
	registerZUNION(p)
	registerZUNIONWITHSCORES(p)
	registerZRANDMEMBER(p)
	registerZDIFF(p)
	registerZDIFFWITHSCORES(p)
	registerZDIFFSTORE(p)
	registerPFADD(p)
	registerPFCOUNT(p)
	registerPFMERGE(p)
	registerBGREWRITEAOF(p)
	registerBGSAVE(p)
	registerCLIENTKILL(p)
	registerCLIENTKILLBYFILTER(p)
	registerCLIENTLIST(p)
	registerCLIENTPAUSE(p)
	registerCLIENTID(p)
	registerCONFIGGET(p)
	registerCONFIGRESETSTAT(p)
	registerCONFIGSET(p)
	registerCONFIGREWRITE(p)
	registerDBSIZE(p)
	registerFLUSHALL(p)
	registerFLUSHALLASYNC(p)
	registerFLUSHDB(p)
	registerFLUSHDBASYNC(p)
	registerINFO(p)
	registerLASTSAVE(p)
	registerSAVE(p)
	registerSHUTDOWN(p)
	registerSHUTDOWNSAVE(p)
	registerSHUTDOWNNOSAVE(p)
	registerSLAVEOF(p)
	registerTIME(p)
	registerDEBUGOBJECT(p)
	registerREADONLY(p)
	registerREADWRITE(p)
	registerMEMORYUSAGE(p)
	registerEVAL(p)
	registerEVALSHA(p)
	registerSCRIPTEXISTS(p)
	registerSCRIPTFLUSH(p)
	registerSCRIPTKILL(p)
	registerSCRIPTLOAD(p)
	registerPUBLISH(p)
	registerPUBSUBCHANNELS(p)
	registerPUBSUBNUMSUB(p)
	registerPUBSUBNUMPAT(p)
	registerCLUSTERSLOTS(p)
	registerCLUSTERNODES(p)
	registerCLUSTERMEET(p)
	registerCLUSTERFORGET(p)
	registerCLUSTERREPLICATE(p)
	registerCLUSTERRESETSOFT(p)
	registerCLUSTERRESETHARD(p)
	registerCLUSTERINFO(p)
	registerCLUSTERKEYSLOT(p)
	registerCLUSTERGETKEYSINSLOT(p)
	registerCLUSTERCOUNTFAILUREREPORTS(p)
	registerCLUSTERCOUNTKEYSINSLOT(p)
	registerCLUSTERDELSLOTS(p)
	registerCLUSTERDELSLOTSRANGE(p)
	registerCLUSTERSAVECONFIG(p)
	registerCLUSTERSLAVES(p)
	registerCLUSTERFAILOVER(p)
	registerCLUSTERADDSLOTS(p)
	registerCLUSTERADDSLOTSRANGE(p)
	registerGEOADD(p)
	registerGEOPOS(p)
	registerGEORADIUS(p)
	registerGEORADIUSSTORE(p)
	registerGEORADIUSBYMEMBER(p)
	registerGEORADIUSBYMEMBERSTORE(p)
	registerGEOSEARCH(p)
	registerGEOSEARCHLOCATION(p)
	registerGEOSEARCHSTORE(p)
	registerGEODIST(p)
	registerGEOHASH(p)
}

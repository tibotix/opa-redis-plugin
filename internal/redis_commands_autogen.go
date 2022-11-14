// Code generated! DO NOT EDIT
package internal

import (
    "time"
    "github.com/tibotix/opa-redis-plugin/utils"

    "github.com/go-redis/redis/v8"
    "github.com/open-policy-agent/opa/ast"
    "github.com/open-policy-agent/opa/rego"
    "github.com/open-policy-agent/opa/types"
)


func registerCLIENTGETNAME(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.clientgetname",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClientGetName(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.echo",
            Decl: types.NewFunction(types.Args(types.Any{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Echo(p.redisContext,utils.Conv(v0)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.ping",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.Ping(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.quit",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.Quit(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.del",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Del(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.unlink",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Unlink(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.dump",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Dump(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.exists",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Exists(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.expire",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.Expire(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.expireat",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.ExpireAt(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.expirenx",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.ExpireNX(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.expirexx",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.ExpireXX(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.expiregt",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.ExpireGT(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.expirelt",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.ExpireLT(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.keys",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.Keys(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerMIGRATE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.migrate",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{},types.Number{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Migrate(p.redisContext,v0,v1,v2,v3,v4).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.move",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.Move(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.objectrefcount",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ObjectRefCount(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.objectencoding",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ObjectEncoding(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.objectidletime",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ObjectIdleTime(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val.Seconds())) 
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
            Name: "redis.persist",
            Decl: types.NewFunction(types.Args(types.String{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.Persist(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.pexpire",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.PExpire(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.pexpireat",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.PExpireAt(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.pttl",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.PTTL(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val.Milliseconds())) 
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
            Name: "redis.randomkey",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.RandomKey(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.rename",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Rename(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.renamenx",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.RenameNX(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.restore",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Restore(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.restorereplace",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.RestoreReplace(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.touch",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Touch(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.ttl",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.TTL(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val.Seconds())) 
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
            Name: "redis.type",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Type(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.append",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Append(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.decr",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Decr(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.decrby",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.DecrBy(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.get",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Get(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.getrange",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.GetRange(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.getset",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.GetSet(p.redisContext,v0,utils.Conv(v1)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.getex",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.GetEx(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.getdel",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.GetDel(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.incr",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Incr(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.incrby",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.IncrBy(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.incrbyfloat",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.IncrByFloat(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.FloatNumberTerm(float64(val)) 
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
            Name: "redis.mget",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Any{})),
            Memoize: true,
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
            

            val, err := rdb.MGet(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerMSET(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.mset",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Any{})), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.MSet(p.redisContext,utils.Conva(v0)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.msetnx",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Any{})), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.MSetNX(p.redisContext,utils.Conva(v0)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.set",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Set(p.redisContext,v0,utils.Conv(v1),v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.setex",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.SetEX(p.redisContext,v0,utils.Conv(v1),v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.setnx",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.SetNX(p.redisContext,v0,utils.Conv(v1),v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.setxx",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.SetXX(p.redisContext,v0,utils.Conv(v1),v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.setrange",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SetRange(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.strlen",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.StrLen(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.copy",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Number{},types.Boolean{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Copy(p.redisContext,v0,v1,v2,v3).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.getbit",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.GetBit(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.setbit",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SetBit(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.bitopand",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.BitOpAnd(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.bitopor",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.BitOpOr(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.bitopxor",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.BitOpXor(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.bitopnot",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.BitOpNot(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.bitpos",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.NewArray([]types.Type{}, types.Number{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.BitPos(p.redisContext,v0,v1,v2...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.bitfield",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.NewArray([]types.Type{}, types.Number{})),
            Memoize: true,
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
            

            val, err := rdb.BitField(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.IntNumberTerm(int(v))
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerHDEL(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.hdel",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.HDel(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.hexists",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.HExists(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.hget",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.HGet(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.hincrby",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.HIncrBy(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.hincrbyfloat",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.HIncrByFloat(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.FloatNumberTerm(float64(val)) 
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
            Name: "redis.hkeys",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.HKeys(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerHLEN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.hlen",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.HLen(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.hmget",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Any{})),
            Memoize: true,
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
            

            val, err := rdb.HMGet(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerHSET(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.hset",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.HSet(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.hmset",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.HMSet(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.hsetnx",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Any{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.HSetNX(p.redisContext,v0,v1,utils.Conv(v2)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.hvals",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.HVals(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerHRANDFIELD(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.hrandfield",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Boolean{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.HRandField(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerBLPOP(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.blpop",
            Decl: types.NewFunction(types.Args(types.Number{},types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.BLPop(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerBRPOP(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.brpop",
            Decl: types.NewFunction(types.Args(types.Number{},types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.BRPop(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerBRPOPLPUSH(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.brpoplpush",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.BRPopLPush(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.lindex",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.LIndex(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.linsert",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Any{},types.Any{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LInsert(p.redisContext,v0,v1,utils.Conv(v2),utils.Conv(v3)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.linsertbefore",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.Any{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LInsertBefore(p.redisContext,v0,utils.Conv(v1),utils.Conv(v2)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.linsertafter",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.Any{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LInsertAfter(p.redisContext,v0,utils.Conv(v1),utils.Conv(v2)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.llen",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LLen(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.lpop",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.LPop(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.lpopcount",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.LPopCount(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerLPUSH(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.lpush",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LPush(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.lpushx",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LPushX(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.lrange",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.LRange(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerLREM(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.lrem",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Any{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.LRem(p.redisContext,v0,v1,utils.Conv(v2)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.lset",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Any{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.LSet(p.redisContext,v0,v1,utils.Conv(v2)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.ltrim",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.LTrim(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.rpop",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.RPop(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.rpopcount",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.RPopCount(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerRPOPLPUSH(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.rpoplpush",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.RPopLPush(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.rpush",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.RPush(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.rpushx",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.RPushX(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.lmove",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.LMove(p.redisContext,v0,v1,v2,v3).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.blmove",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{},types.String{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.BLMove(p.redisContext,v0,v1,v2,v3,v4).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.sadd",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SAdd(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.scard",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SCard(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.sdiff",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.SDiff(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSDIFFSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.sdiffstore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SDiffStore(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.sinter",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.SInter(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSINTERSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.sinterstore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SInterStore(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.sismember",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.SIsMember(p.redisContext,v0,utils.Conv(v1)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.smismember",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.NewArray([]types.Type{}, types.Boolean{})),
            Memoize: true,
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
            

            val, err := rdb.SMIsMember(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.BooleanTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSMEMBERS(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.smembers",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.SMembers(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSMOVE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.smove",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Any{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.SMove(p.redisContext,v0,v1,utils.Conv(v2)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.spop",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.SPop(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.spopn",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.SPopN(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSRANDMEMBER(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.srandmember",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.SRandMember(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.srandmembern",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.SRandMemberN(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSREM(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.srem",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SRem(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.sunion",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.SUnion(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSUNIONSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.sunionstore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.SUnionStore(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xdel",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XDel(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xlen",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XLen(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xgroupcreate",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.XGroupCreate(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.xgroupcreatemkstream",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.XGroupCreateMkStream(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.xgroupsetid",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.XGroupSetID(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.xgroupdestroy",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XGroupDestroy(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xgroupcreateconsumer",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XGroupCreateConsumer(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xgroupdelconsumer",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XGroupDelConsumer(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xack",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XAck(p.redisContext,v0,v1,v2...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
                return term, nil
            default:
                return nil, err
            }
        },
    )
}


func registerXTRIM(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.xtrim",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XTrim(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xtrimapprox",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XTrimApprox(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xtrimmaxlen",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XTrimMaxLen(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xtrimmaxlenapprox",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XTrimMaxLenApprox(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xtrimminid",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XTrimMinID(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.xtrimminidapprox",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.XTrimMinIDApprox(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zadd",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZAdd(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zaddnx",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZAddNX(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zaddxx",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZAddXX(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zaddch",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZAddCh(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zaddnxch",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZAddNXCh(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zaddxxch",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZAddXXCh(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zcard",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZCard(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zcount",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZCount(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zlexcount",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZLexCount(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zincrby",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZIncrBy(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.FloatNumberTerm(float64(val)) 
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
            Name: "redis.zmscore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Number{})),
            Memoize: true,
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
            

            val, err := rdb.ZMScore(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.FloatNumberTerm(float64(v))
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrange",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.ZRange(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerZRANK(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrank",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZRank(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zrem",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZRem(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zremrangebyrank",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZRemRangeByRank(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zremrangebyscore",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZRemRangeByScore(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zremrangebylex",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZRemRangeByLex(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zrevrange",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.ZRevRange(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerZREVRANK(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrevrank",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZRevRank(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.zscore",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZScore(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.FloatNumberTerm(float64(val)) 
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
            Name: "redis.zrandmember",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Boolean{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.ZRandMember(p.redisContext,v0,v1,v2).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerZDIFF(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zdiff",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.ZDiff(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerZDIFFSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zdiffstore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ZDiffStore(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.pfadd",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Any{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.PFAdd(p.redisContext,v0,utils.Conva(v1)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.pfcount",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.PFCount(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.pfmerge",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.PFMerge(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.bgrewriteaof",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.BgRewriteAOF(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.bgsave",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.BgSave(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clientkill",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClientKill(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clientkillbyfilter",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ClientKillByFilter(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.clientlist",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClientList(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clientpause",
            Decl: types.NewFunction(types.Args(types.Number{}), types.Boolean{}),
            Memoize: true,
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
            

            val, err := rdb.ClientPause(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.BooleanTerm(val) 
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
            Name: "redis.clientid",
            Decl: types.NewFunction(types.Args(), types.Number{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClientID(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.configget",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.Any{})),
            Memoize: true,
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
            

            val, err := rdb.ConfigGet(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerCONFIGRESETSTAT(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.configresetstat",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ConfigResetStat(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.configset",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ConfigSet(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.configrewrite",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ConfigRewrite(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.dbsize",
            Decl: types.NewFunction(types.Args(), types.Number{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.DBSize(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.flushall",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.FlushAll(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.flushallasync",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.FlushAllAsync(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.flushdb",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.FlushDB(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.flushdbasync",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.FlushDBAsync(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.info",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.Info(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.lastsave",
            Decl: types.NewFunction(types.Args(), types.Number{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.LastSave(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.save",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.Save(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.shutdown",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.Shutdown(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.shutdownsave",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ShutdownSave(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.shutdownnosave",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ShutdownNoSave(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.slaveof",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.SlaveOf(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.time",
            Decl: types.NewFunction(types.Args(), types.Number{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.Time(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val.UnixMicro())) 
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
            Name: "redis.debugobject",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.DebugObject(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.readonly",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ReadOnly(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.readwrite",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ReadWrite(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.memoryusage",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Number{})), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.MemoryUsage(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.eval",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{}),types.NewArray([]types.Type{}, types.Any{})), types.Any{}),
            Memoize: true,
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
            

            val, err := rdb.Eval(p.redisContext,v0,v1,utils.Conva(v2)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                
            term := ast.NullTerm()
            if s, ok := val.(string); ok {
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
            Name: "redis.evalsha",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{}),types.NewArray([]types.Type{}, types.Any{})), types.Any{}),
            Memoize: true,
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
            

            val, err := rdb.EvalSha(p.redisContext,v0,v1,utils.Conva(v2)...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                
            term := ast.NullTerm()
            if s, ok := val.(string); ok {
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
            Name: "redis.scriptexists",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.Boolean{})),
            Memoize: true,
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
            

            val, err := rdb.ScriptExists(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.BooleanTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerSCRIPTFLUSH(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.scriptflush",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ScriptFlush(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.scriptkill",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ScriptKill(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.scriptload",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ScriptLoad(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.publish",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.Publish(p.redisContext,v0,utils.Conv(v1)).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.pubsubchannels",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.PubSubChannels(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerPUBSUBNUMPAT(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.pubsubnumpat",
            Decl: types.NewFunction(types.Args(), types.Number{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.PubSubNumPat(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.clusternodes",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClusterNodes(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clustermeet",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterMeet(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterforget",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterForget(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterreplicate",
            Decl: types.NewFunction(types.Args(types.String{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterReplicate(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterresetsoft",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClusterResetSoft(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterresethard",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClusterResetHard(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterinfo",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClusterInfo(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterkeyslot",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterKeySlot(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.clustergetkeysinslot",
            Decl: types.NewFunction(types.Args(types.Number{},types.Number{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.ClusterGetKeysInSlot(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerCLUSTERCOUNTFAILUREREPORTS(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.clustercountfailurereports",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterCountFailureReports(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.clustercountkeysinslot",
            Decl: types.NewFunction(types.Args(types.Number{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterCountKeysInSlot(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.clusterdelslots",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Number{})), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterDelSlots(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterdelslotsrange",
            Decl: types.NewFunction(types.Args(types.Number{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterDelSlotsRange(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clustersaveconfig",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClusterSaveConfig(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusterslaves",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.ClusterSlaves(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
            default:
                return nil, err
            }
        },
    )
}


func registerCLUSTERFAILOVER(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.clusterfailover",
            Decl: types.NewFunction(types.Args(), types.String{}),
            Memoize: true,
            Nondeterministic: true,
        },
        func(bctx rego.BuiltinContext, terms []*ast.Term) (*ast.Term, error) {
            rdb, err := p.redisProxy.Get()
            if err != nil {
                return nil, err
            }



            val, err := rdb.ClusterFailover(p.redisContext).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusteraddslots",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.Number{})), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterAddSlots(p.redisContext,v0...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.clusteraddslotsrange",
            Decl: types.NewFunction(types.Args(types.Number{},types.Number{}), types.String{}),
            Memoize: true,
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
            

            val, err := rdb.ClusterAddSlotsRange(p.redisContext,v0,v1).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.StringTerm(val) 
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
            Name: "redis.geoadd",
            Decl: types.NewFunction(types.Args(types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.GeoAdd(p.redisContext,v0).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.IntNumberTerm(int(val)) 
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
            Name: "redis.geodist",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.String{},types.String{}), types.Number{}),
            Memoize: true,
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
            

            val, err := rdb.GeoDist(p.redisContext,v0,v1,v2,v3).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
                term := ast.FloatNumberTerm(float64(val)) 
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
            Name: "redis.geohash",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val, err := rdb.GeoHash(p.redisContext,v0,v1...).Result()
            switch err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var ret []*ast.Term
        for _, v := range val {
            term := ast.StringTerm(v)
            ret = append(ret, term)
        }
        return ast.ArrayTerm(ret...), nil
                
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
    registerSETEX(p)
    registerSETNX(p)
    registerSETXX(p)
    registerSETRANGE(p)
    registerSTRLEN(p)
    registerCOPY(p)
    registerGETBIT(p)
    registerSETBIT(p)
    registerBITOPAND(p)
    registerBITOPOR(p)
    registerBITOPXOR(p)
    registerBITOPNOT(p)
    registerBITPOS(p)
    registerBITFIELD(p)
    registerHDEL(p)
    registerHEXISTS(p)
    registerHGET(p)
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
    registerSMOVE(p)
    registerSPOP(p)
    registerSPOPN(p)
    registerSRANDMEMBER(p)
    registerSRANDMEMBERN(p)
    registerSREM(p)
    registerSUNION(p)
    registerSUNIONSTORE(p)
    registerXDEL(p)
    registerXLEN(p)
    registerXGROUPCREATE(p)
    registerXGROUPCREATEMKSTREAM(p)
    registerXGROUPSETID(p)
    registerXGROUPDESTROY(p)
    registerXGROUPCREATECONSUMER(p)
    registerXGROUPDELCONSUMER(p)
    registerXACK(p)
    registerXTRIM(p)
    registerXTRIMAPPROX(p)
    registerXTRIMMAXLEN(p)
    registerXTRIMMAXLENAPPROX(p)
    registerXTRIMMINID(p)
    registerXTRIMMINIDAPPROX(p)
    registerZADD(p)
    registerZADDNX(p)
    registerZADDXX(p)
    registerZADDCH(p)
    registerZADDNXCH(p)
    registerZADDXXCH(p)
    registerZCARD(p)
    registerZCOUNT(p)
    registerZLEXCOUNT(p)
    registerZINCRBY(p)
    registerZMSCORE(p)
    registerZRANGE(p)
    registerZRANK(p)
    registerZREM(p)
    registerZREMRANGEBYRANK(p)
    registerZREMRANGEBYSCORE(p)
    registerZREMRANGEBYLEX(p)
    registerZREVRANGE(p)
    registerZREVRANK(p)
    registerZSCORE(p)
    registerZRANDMEMBER(p)
    registerZDIFF(p)
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
    registerPUBSUBNUMPAT(p)
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
    registerGEODIST(p)
    registerGEOHASH(p)
}

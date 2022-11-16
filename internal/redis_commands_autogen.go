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
            

            val := rdb.Echo(p.redisContext,utils.Conv(v0))
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
            

            val := rdb.Del(p.redisContext,v0...)
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
            

            val := rdb.Unlink(p.redisContext,v0...)
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
            

            val := rdb.Dump(p.redisContext,v0)
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
            

            val := rdb.Exists(p.redisContext,v0...)
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
            

            val := rdb.Expire(p.redisContext,v0,v1)
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
            

            val := rdb.ExpireAt(p.redisContext,v0,v1)
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
            

            val := rdb.ExpireNX(p.redisContext,v0,v1)
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
            

            val := rdb.ExpireXX(p.redisContext,v0,v1)
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
            

            val := rdb.ExpireGT(p.redisContext,v0,v1)
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
            

            val := rdb.ExpireLT(p.redisContext,v0,v1)
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
            

            val := rdb.Keys(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp27295 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp27295 = append(tmp27295, term)
        }
        term := ast.ArrayTerm(tmp27295...)
        
        return term,  nil
        
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
            

            val := rdb.Migrate(p.redisContext,v0,v1,v2,v3,v4)
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
            

            val := rdb.Move(p.redisContext,v0,v1)
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
            

            val := rdb.ObjectRefCount(p.redisContext,v0)
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
            

            val := rdb.ObjectEncoding(p.redisContext,v0)
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
            

            val := rdb.ObjectIdleTime(p.redisContext,v0)
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
            

            val := rdb.Persist(p.redisContext,v0)
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
            

            val := rdb.PExpire(p.redisContext,v0,v1)
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
            

            val := rdb.PExpireAt(p.redisContext,v0,v1)
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
            

            val := rdb.PTTL(p.redisContext,v0)
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
            

            val := rdb.Rename(p.redisContext,v0,v1)
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
            

            val := rdb.RenameNX(p.redisContext,v0,v1)
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
            

            val := rdb.Restore(p.redisContext,v0,v1,v2)
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
            

            val := rdb.RestoreReplace(p.redisContext,v0,v1,v2)
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
            Name: "redis.sort",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Order", types.String{}),types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
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
            
            var v1 *redis.Sort
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.Sort(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp38323 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp38323 = append(tmp38323, term)
        }
        term := ast.ArrayTerm(tmp38323...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerSORTSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.sortstore",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Order", types.String{}),types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.Number{}),
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
            
            var v2 *redis.Sort
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            

            val := rdb.SortStore(p.redisContext,v0,v1,v2)
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
            Name: "redis.sortinterfaces",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("By", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Get", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Order", types.String{}),types.NewStaticProperty("Alpha", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.Any{})),
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
            
            var v1 *redis.Sort
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.SortInterfaces(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp59721 []*ast.Term
        for _, v := range r0 {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            tmp59721 = append(tmp59721, term)
        }
        term := ast.ArrayTerm(tmp59721...)
        
        return term,  nil
        
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
            

            val := rdb.Touch(p.redisContext,v0...)
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
            

            val := rdb.TTL(p.redisContext,v0)
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
            

            val := rdb.Type(p.redisContext,v0)
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
            

            val := rdb.Append(p.redisContext,v0,v1)
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
            

            val := rdb.Decr(p.redisContext,v0)
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
            

            val := rdb.DecrBy(p.redisContext,v0,v1)
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
            

            val := rdb.Get(p.redisContext,v0)
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
            

            val := rdb.GetRange(p.redisContext,v0,v1,v2)
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
            

            val := rdb.GetSet(p.redisContext,v0,utils.Conv(v1))
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
            

            val := rdb.GetEx(p.redisContext,v0,v1)
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
            

            val := rdb.GetDel(p.redisContext,v0)
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
            

            val := rdb.Incr(p.redisContext,v0)
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
            

            val := rdb.IncrBy(p.redisContext,v0,v1)
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
            

            val := rdb.IncrByFloat(p.redisContext,v0,v1)
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
            

            val := rdb.MGet(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp90492 []*ast.Term
        for _, v := range r0 {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            tmp90492 = append(tmp90492, term)
        }
        term := ast.ArrayTerm(tmp90492...)
        
        return term,  nil
        
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
            

            val := rdb.MSet(p.redisContext,utils.Conva(v0)...)
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
            

            val := rdb.MSetNX(p.redisContext,utils.Conva(v0)...)
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
            

            val := rdb.Set(p.redisContext,v0,utils.Conv(v1),v2)
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
            Name: "redis.setargs",
            Decl: types.NewFunction(types.Args(types.String{},types.Any{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Mode", types.String{}),types.NewStaticProperty("TTL", types.Number{}),types.NewStaticProperty("ExpireAt", types.Number{}),types.NewStaticProperty("Get", types.Boolean{}),types.NewStaticProperty("KeepTTL", types.Boolean{})}, nil)), types.String{}),
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
            
            var v2 redis.SetArgs
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            

            val := rdb.SetArgs(p.redisContext,v0,utils.Conv(v1),v2)
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
            

            val := rdb.SetEX(p.redisContext,v0,utils.Conv(v1),v2)
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
            

            val := rdb.SetNX(p.redisContext,v0,utils.Conv(v1),v2)
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
            

            val := rdb.SetXX(p.redisContext,v0,utils.Conv(v1),v2)
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
            

            val := rdb.SetRange(p.redisContext,v0,v1,v2)
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
            

            val := rdb.StrLen(p.redisContext,v0)
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
            

            val := rdb.Copy(p.redisContext,v0,v1,v2,v3)
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
            

            val := rdb.GetBit(p.redisContext,v0,v1)
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
            

            val := rdb.SetBit(p.redisContext,v0,v1,v2)
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
            Name: "redis.bitcount",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Start", types.Number{}),types.NewStaticProperty("End", types.Number{})}, nil)), types.Number{}),
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
            
            var v1 *redis.BitCount
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.BitCount(p.redisContext,v0,v1)
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
            

            val := rdb.BitOpAnd(p.redisContext,v0,v1...)
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
            

            val := rdb.BitOpOr(p.redisContext,v0,v1...)
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
            

            val := rdb.BitOpXor(p.redisContext,v0,v1...)
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
            

            val := rdb.BitOpNot(p.redisContext,v0,v1)
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
            

            val := rdb.BitPos(p.redisContext,v0,v1,v2...)
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
            

            val := rdb.BitField(p.redisContext,v0,utils.Conva(v1)...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp49444 []*ast.Term
        for _, v := range r0 {
            term := ast.IntNumberTerm(int(v))
            tmp49444 = append(tmp49444, term)
        }
        term := ast.ArrayTerm(tmp49444...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerSCAN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.scan",
            Decl: types.NewFunction(types.Args(types.Number{},types.String{},types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}),types.Number{}}, types.Null{})),
            Memoize: true,
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
            

            val := rdb.Scan(p.redisContext,v0,v1,v2)
            r0,r1 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var tmp54047 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp54047 = append(tmp54047, term)
        }
        tr0 := ast.ArrayTerm(tmp54047...)
        tr1 := ast.UIntNumberTerm(uint64(r1))
        return ast.ArrayTerm(tr0,tr1), nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerSCANTYPE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.scantype",
            Decl: types.NewFunction(types.Args(types.Number{},types.String{},types.Number{},types.String{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}),types.Number{}}, types.Null{})),
            Memoize: true,
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
            

            val := rdb.ScanType(p.redisContext,v0,v1,v2,v3)
            r0,r1 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var tmp48347 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp48347 = append(tmp48347, term)
        }
        tr0 := ast.ArrayTerm(tmp48347...)
        tr1 := ast.UIntNumberTerm(uint64(r1))
        return ast.ArrayTerm(tr0,tr1), nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerSSCAN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.sscan",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{},types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}),types.Number{}}, types.Null{})),
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
            

            val := rdb.SScan(p.redisContext,v0,v1,v2,v3)
            r0,r1 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var tmp26155 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp26155 = append(tmp26155, term)
        }
        tr0 := ast.ArrayTerm(tmp26155...)
        tr1 := ast.UIntNumberTerm(uint64(r1))
        return ast.ArrayTerm(tr0,tr1), nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerHSCAN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.hscan",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{},types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}),types.Number{}}, types.Null{})),
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
            

            val := rdb.HScan(p.redisContext,v0,v1,v2,v3)
            r0,r1 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var tmp86672 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp86672 = append(tmp86672, term)
        }
        tr0 := ast.ArrayTerm(tmp86672...)
        tr1 := ast.UIntNumberTerm(uint64(r1))
        return ast.ArrayTerm(tr0,tr1), nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZSCAN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zscan",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.String{},types.Number{}), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}),types.Number{}}, types.Null{})),
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
            

            val := rdb.ZScan(p.redisContext,v0,v1,v2,v3)
            r0,r1 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var tmp73570 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp73570 = append(tmp73570, term)
        }
        tr0 := ast.ArrayTerm(tmp73570...)
        tr1 := ast.UIntNumberTerm(uint64(r1))
        return ast.ArrayTerm(tr0,tr1), nil
        
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
            

            val := rdb.HDel(p.redisContext,v0,v1...)
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
            

            val := rdb.HExists(p.redisContext,v0,v1)
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
            

            val := rdb.HGet(p.redisContext,v0,v1)
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
            

            val := rdb.HIncrBy(p.redisContext,v0,v1,v2)
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
            

            val := rdb.HIncrByFloat(p.redisContext,v0,v1,v2)
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
            

            val := rdb.HKeys(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp33540 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp33540 = append(tmp33540, term)
        }
        term := ast.ArrayTerm(tmp33540...)
        
        return term,  nil
        
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
            

            val := rdb.HLen(p.redisContext,v0)
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
            

            val := rdb.HMGet(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp24137 []*ast.Term
        for _, v := range r0 {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            tmp24137 = append(tmp24137, term)
        }
        term := ast.ArrayTerm(tmp24137...)
        
        return term,  nil
        
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
            

            val := rdb.HSet(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.HMSet(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.HSetNX(p.redisContext,v0,v1,utils.Conv(v2))
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
            

            val := rdb.HVals(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp61375 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp61375 = append(tmp61375, term)
        }
        term := ast.ArrayTerm(tmp61375...)
        
        return term,  nil
        
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
            

            val := rdb.HRandField(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp74838 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp74838 = append(tmp74838, term)
        }
        term := ast.ArrayTerm(tmp74838...)
        
        return term,  nil
        
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
            

            val := rdb.BLPop(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp66056 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp66056 = append(tmp66056, term)
        }
        term := ast.ArrayTerm(tmp66056...)
        
        return term,  nil
        
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
            

            val := rdb.BRPop(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp83590 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp83590 = append(tmp83590, term)
        }
        term := ast.ArrayTerm(tmp83590...)
        
        return term,  nil
        
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
            

            val := rdb.BRPopLPush(p.redisContext,v0,v1,v2)
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
            

            val := rdb.LIndex(p.redisContext,v0,v1)
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
            

            val := rdb.LInsert(p.redisContext,v0,v1,utils.Conv(v2),utils.Conv(v3))
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
            

            val := rdb.LInsertBefore(p.redisContext,v0,utils.Conv(v1),utils.Conv(v2))
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
            

            val := rdb.LInsertAfter(p.redisContext,v0,utils.Conv(v1),utils.Conv(v2))
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
            

            val := rdb.LLen(p.redisContext,v0)
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
            

            val := rdb.LPop(p.redisContext,v0)
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
            

            val := rdb.LPopCount(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp37835 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp37835 = append(tmp37835, term)
        }
        term := ast.ArrayTerm(tmp37835...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerLPOS(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.lpos",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Rank", types.Number{}),types.NewStaticProperty("MaxLen", types.Number{})}, nil)), types.Number{}),
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
            
            var v2 redis.LPosArgs
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            

            val := rdb.LPos(p.redisContext,v0,v1,v2)
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
            Name: "redis.lposcount",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.Number{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Rank", types.Number{}),types.NewStaticProperty("MaxLen", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.Number{})),
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
            
            var v3 redis.LPosArgs
            if err := ast.As(terms[3].Value, &v3); err != nil {
                return nil, err
            }
            

            val := rdb.LPosCount(p.redisContext,v0,v1,v2,v3)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp37359 []*ast.Term
        for _, v := range r0 {
            term := ast.IntNumberTerm(int(v))
            tmp37359 = append(tmp37359, term)
        }
        term := ast.ArrayTerm(tmp37359...)
        
        return term,  nil
        
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
            

            val := rdb.LPush(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.LPushX(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.LRange(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp35485 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp35485 = append(tmp35485, term)
        }
        term := ast.ArrayTerm(tmp35485...)
        
        return term,  nil
        
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
            

            val := rdb.LRem(p.redisContext,v0,v1,utils.Conv(v2))
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
            

            val := rdb.LSet(p.redisContext,v0,v1,utils.Conv(v2))
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
            

            val := rdb.LTrim(p.redisContext,v0,v1,v2)
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
            

            val := rdb.RPop(p.redisContext,v0)
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
            

            val := rdb.RPopCount(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp73402 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp73402 = append(tmp73402, term)
        }
        term := ast.ArrayTerm(tmp73402...)
        
        return term,  nil
        
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
            

            val := rdb.RPopLPush(p.redisContext,v0,v1)
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
            

            val := rdb.RPush(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.RPushX(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.LMove(p.redisContext,v0,v1,v2,v3)
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
            

            val := rdb.BLMove(p.redisContext,v0,v1,v2,v3,v4)
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
            

            val := rdb.SAdd(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.SCard(p.redisContext,v0)
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
            

            val := rdb.SDiff(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp77349 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp77349 = append(tmp77349, term)
        }
        term := ast.ArrayTerm(tmp77349...)
        
        return term,  nil
        
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
            

            val := rdb.SDiffStore(p.redisContext,v0,v1...)
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
            

            val := rdb.SInter(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp89853 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp89853 = append(tmp89853, term)
        }
        term := ast.ArrayTerm(tmp89853...)
        
        return term,  nil
        
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
            

            val := rdb.SInterStore(p.redisContext,v0,v1...)
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
            

            val := rdb.SIsMember(p.redisContext,v0,utils.Conv(v1))
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
            

            val := rdb.SMIsMember(p.redisContext,v0,utils.Conva(v1)...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp57920 []*ast.Term
        for _, v := range r0 {
            term := ast.BooleanTerm(v)
            tmp57920 = append(tmp57920, term)
        }
        term := ast.ArrayTerm(tmp57920...)
        
        return term,  nil
        
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
            

            val := rdb.SMembers(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp91369 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp91369 = append(tmp91369, term)
        }
        term := ast.ArrayTerm(tmp91369...)
        
        return term,  nil
        
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
            

            val := rdb.SMove(p.redisContext,v0,v1,utils.Conv(v2))
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
            

            val := rdb.SPop(p.redisContext,v0)
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
            

            val := rdb.SPopN(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp79958 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp79958 = append(tmp79958, term)
        }
        term := ast.ArrayTerm(tmp79958...)
        
        return term,  nil
        
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
            

            val := rdb.SRandMember(p.redisContext,v0)
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
            

            val := rdb.SRandMemberN(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp86254 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp86254 = append(tmp86254, term)
        }
        term := ast.ArrayTerm(tmp86254...)
        
        return term,  nil
        
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
            

            val := rdb.SRem(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.SUnion(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp72851 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp72851 = append(tmp72851, term)
        }
        term := ast.ArrayTerm(tmp72851...)
        
        return term,  nil
        
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
            

            val := rdb.SUnionStore(p.redisContext,v0,v1...)
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
            Name: "redis.xadd",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}),types.NewStaticProperty("NoMkStream", types.Boolean{}),types.NewStaticProperty("MaxLen", types.Number{}),types.NewStaticProperty("MaxLenApprox", types.Number{}),types.NewStaticProperty("MinID", types.String{}),types.NewStaticProperty("Approx", types.Boolean{}),types.NewStaticProperty("Limit", types.Number{}),types.NewStaticProperty("ID", types.String{}),types.NewStaticProperty("Values", types.Any{})}, nil)), types.String{}),
            Memoize: true,
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
            

            val := rdb.XAdd(p.redisContext,v0)
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
            

            val := rdb.XDel(p.redisContext,v0,v1...)
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
            

            val := rdb.XLen(p.redisContext,v0)
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
            

            val := rdb.XGroupCreate(p.redisContext,v0,v1,v2)
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
            

            val := rdb.XGroupCreateMkStream(p.redisContext,v0,v1,v2)
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
            

            val := rdb.XGroupSetID(p.redisContext,v0,v1,v2)
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
            

            val := rdb.XGroupDestroy(p.redisContext,v0,v1)
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
            

            val := rdb.XGroupCreateConsumer(p.redisContext,v0,v1,v2)
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
            

            val := rdb.XGroupDelConsumer(p.redisContext,v0,v1,v2)
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
            

            val := rdb.XAck(p.redisContext,v0,v1,v2...)
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


func registerXCLAIMJUSTID(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.xclaimjustid",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}),types.NewStaticProperty("Group", types.String{}),types.NewStaticProperty("Consumer", types.String{}),types.NewStaticProperty("MinIdle", types.Number{}),types.NewStaticProperty("Messages", types.NewArray([]types.Type{}, types.String{}))}, nil)), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val := rdb.XClaimJustID(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp79423 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp79423 = append(tmp79423, term)
        }
        term := ast.ArrayTerm(tmp79423...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerXAUTOCLAIMJUSTID(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.xautoclaimjustid",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Stream", types.String{}),types.NewStaticProperty("Group", types.String{}),types.NewStaticProperty("MinIdle", types.Number{}),types.NewStaticProperty("Start", types.String{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Consumer", types.String{})}, nil)), types.NewArray([]types.Type{types.NewArray([]types.Type{}, types.String{}),types.String{}}, types.Null{})),
            Memoize: true,
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
            

            val := rdb.XAutoClaimJustID(p.redisContext,v0)
            r0,r1 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        var tmp29908 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp29908 = append(tmp29908, term)
        }
        tr0 := ast.ArrayTerm(tmp29908...)
        tr1 := ast.StringTerm(r1)
        return ast.ArrayTerm(tr0,tr1), nil
        
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
            

            val := rdb.XTrim(p.redisContext,v0,v1)
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
            

            val := rdb.XTrimApprox(p.redisContext,v0,v1)
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
            

            val := rdb.XTrimMaxLen(p.redisContext,v0,v1)
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
            

            val := rdb.XTrimMaxLenApprox(p.redisContext,v0,v1,v2)
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
            

            val := rdb.XTrimMinID(p.redisContext,v0,v1)
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
            

            val := rdb.XTrimMinIDApprox(p.redisContext,v0,v1,v2)
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
            Name: "redis.xinfogroups",
            Decl: types.NewFunction(types.Args(types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}),types.NewStaticProperty("Consumers", types.Number{}),types.NewStaticProperty("Pending", types.Number{}),types.NewStaticProperty("LastDeliveredID", types.String{})}, nil))),
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
            

            val := rdb.XInfoGroups(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp95255 []*ast.Term
        for _, v := range r0 {
            
            name := ast.StringTerm(v.Name)
            
            consumers := ast.IntNumberTerm(int(v.Consumers))
            
            pending := ast.IntNumberTerm(int(v.Pending))
            
            lastdeliveredid := ast.StringTerm(v.LastDeliveredID)
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), name},[2]*ast.Term{ast.StringTerm("Consumers"), consumers},[2]*ast.Term{ast.StringTerm("Pending"), pending},[2]*ast.Term{ast.StringTerm("LastDeliveredID"), lastdeliveredid})
        
            tmp95255 = append(tmp95255, term)
        }
        term := ast.ArrayTerm(tmp95255...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerXINFOCONSUMERS(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.xinfoconsumers",
            Decl: types.NewFunction(types.Args(types.String{},types.String{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}),types.NewStaticProperty("Pending", types.Number{}),types.NewStaticProperty("Idle", types.Number{})}, nil))),
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
            

            val := rdb.XInfoConsumers(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp09508 []*ast.Term
        for _, v := range r0 {
            
            name := ast.StringTerm(v.Name)
            
            pending := ast.IntNumberTerm(int(v.Pending))
            
            idle := ast.IntNumberTerm(int(v.Idle))
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), name},[2]*ast.Term{ast.StringTerm("Pending"), pending},[2]*ast.Term{ast.StringTerm("Idle"), idle})
        
            tmp09508 = append(tmp09508, term)
        }
        term := ast.ArrayTerm(tmp09508...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerBZPOPMAX(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.bzpopmax",
            Decl: types.NewFunction(types.Args(types.Number{},types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{})}, nil)),
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
            

            val := rdb.BZPopMax(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
            key := ast.StringTerm(r0.Key)
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Key"), key})
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerBZPOPMIN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.bzpopmin",
            Decl: types.NewFunction(types.Args(types.Number{},types.NewArray([]types.Type{}, types.String{})), types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{})}, nil)),
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
            

            val := rdb.BZPopMin(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
            key := ast.StringTerm(r0.Key)
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Key"), key})
        
        return term,  nil
        
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
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAdd(p.redisContext,v0,v1...)
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
            Name: "redis.zaddnx",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddNX(p.redisContext,v0,v1...)
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
            Name: "redis.zaddxx",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddXX(p.redisContext,v0,v1...)
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
            Name: "redis.zaddch",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddCh(p.redisContext,v0,v1...)
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
            Name: "redis.zaddnxch",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddNXCh(p.redisContext,v0,v1...)
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
            Name: "redis.zaddxxch",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddXXCh(p.redisContext,v0,v1...)
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
            Name: "redis.zaddargs",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("NX", types.Boolean{}),types.NewStaticProperty("XX", types.Boolean{}),types.NewStaticProperty("LT", types.Boolean{}),types.NewStaticProperty("GT", types.Boolean{}),types.NewStaticProperty("Ch", types.Boolean{}),types.NewStaticProperty("Members", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil)))}, nil)), types.Number{}),
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
            
            var v1 redis.ZAddArgs
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddArgs(p.redisContext,v0,v1)
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
            Name: "redis.zaddargsincr",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("NX", types.Boolean{}),types.NewStaticProperty("XX", types.Boolean{}),types.NewStaticProperty("LT", types.Boolean{}),types.NewStaticProperty("GT", types.Boolean{}),types.NewStaticProperty("Ch", types.Boolean{}),types.NewStaticProperty("Members", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil)))}, nil)), types.Number{}),
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
            
            var v1 redis.ZAddArgs
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZAddArgsIncr(p.redisContext,v0,v1)
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
            Name: "redis.zincr",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
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
            
            var v1 *redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZIncr(p.redisContext,v0,v1)
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
            Name: "redis.zincrnx",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
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
            
            var v1 *redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZIncrNX(p.redisContext,v0,v1)
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
            Name: "redis.zincrxx",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil)), types.Number{}),
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
            
            var v1 *redis.Z
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZIncrXX(p.redisContext,v0,v1)
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
            

            val := rdb.ZCard(p.redisContext,v0)
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
            

            val := rdb.ZCount(p.redisContext,v0,v1,v2)
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
            

            val := rdb.ZLexCount(p.redisContext,v0,v1,v2)
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
            

            val := rdb.ZIncrBy(p.redisContext,v0,v1,v2)
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
            Name: "redis.zinter",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})),types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val := rdb.ZInter(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp38882 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp38882 = append(tmp38882, term)
        }
        term := ast.ArrayTerm(tmp38882...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZINTERWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zinterwithscores",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})),types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
            Memoize: true,
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
            

            val := rdb.ZInterWithScores(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp20093 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp20093 = append(tmp20093, term)
        }
        term := ast.ArrayTerm(tmp20093...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZINTERSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zinterstore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})),types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.Number{}),
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
            
            var v1 *redis.ZStore
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZInterStore(p.redisContext,v0,v1)
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
            

            val := rdb.ZMScore(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp21945 []*ast.Term
        for _, v := range r0 {
            term := ast.FloatNumberTerm(float64(v))
            tmp21945 = append(tmp21945, term)
        }
        term := ast.ArrayTerm(tmp21945...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZPOPMAX(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zpopmax",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Number{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            
            var v1 []int64
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZPopMax(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp15050 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp15050 = append(tmp15050, term)
        }
        term := ast.ArrayTerm(tmp15050...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZPOPMIN(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zpopmin",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.Number{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            
            var v1 []int64
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZPopMin(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp93144 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp93144 = append(tmp93144, term)
        }
        term := ast.ArrayTerm(tmp93144...)
        
        return term,  nil
        
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
            

            val := rdb.ZRange(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp71766 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp71766 = append(tmp71766, term)
        }
        term := ast.ArrayTerm(tmp71766...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGEWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangewithscores",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            

            val := rdb.ZRangeWithScores(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp72121 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp72121 = append(tmp72121, term)
        }
        term := ast.ArrayTerm(tmp72121...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGEBYSCORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangebyscore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}),types.NewStaticProperty("Max", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
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
            
            var v1 *redis.ZRangeBy
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRangeByScore(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp69767 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp69767 = append(tmp69767, term)
        }
        term := ast.ArrayTerm(tmp69767...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGEBYLEX(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangebylex",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}),types.NewStaticProperty("Max", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
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
            
            var v1 *redis.ZRangeBy
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRangeByLex(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp87122 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp87122 = append(tmp87122, term)
        }
        term := ast.ArrayTerm(tmp87122...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGEBYSCOREWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangebyscorewithscores",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}),types.NewStaticProperty("Max", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            
            var v1 *redis.ZRangeBy
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRangeByScoreWithScores(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp48707 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp48707 = append(tmp48707, term)
        }
        term := ast.ArrayTerm(tmp48707...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGEARGS(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangeargs",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}),types.NewStaticProperty("Start", types.Any{}),types.NewStaticProperty("Stop", types.Any{}),types.NewStaticProperty("ByScore", types.Boolean{}),types.NewStaticProperty("ByLex", types.Boolean{}),types.NewStaticProperty("Rev", types.Boolean{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val := rdb.ZRangeArgs(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp31257 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp31257 = append(tmp31257, term)
        }
        term := ast.ArrayTerm(tmp31257...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGEARGSWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangeargswithscores",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}),types.NewStaticProperty("Start", types.Any{}),types.NewStaticProperty("Stop", types.Any{}),types.NewStaticProperty("ByScore", types.Boolean{}),types.NewStaticProperty("ByLex", types.Boolean{}),types.NewStaticProperty("Rev", types.Boolean{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
            Memoize: true,
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
            

            val := rdb.ZRangeArgsWithScores(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp20044 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp20044 = append(tmp20044, term)
        }
        term := ast.ArrayTerm(tmp20044...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZRANGESTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrangestore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Key", types.String{}),types.NewStaticProperty("Start", types.Any{}),types.NewStaticProperty("Stop", types.Any{}),types.NewStaticProperty("ByScore", types.Boolean{}),types.NewStaticProperty("ByLex", types.Boolean{}),types.NewStaticProperty("Rev", types.Boolean{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.Number{}),
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
            
            var v1 redis.ZRangeArgs
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRangeStore(p.redisContext,v0,v1)
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
            

            val := rdb.ZRank(p.redisContext,v0,v1)
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
            

            val := rdb.ZRem(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.ZRemRangeByRank(p.redisContext,v0,v1,v2)
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
            

            val := rdb.ZRemRangeByScore(p.redisContext,v0,v1,v2)
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
            

            val := rdb.ZRemRangeByLex(p.redisContext,v0,v1,v2)
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
            

            val := rdb.ZRevRange(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp71650 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp71650 = append(tmp71650, term)
        }
        term := ast.ArrayTerm(tmp71650...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZREVRANGEWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrevrangewithscores",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{}), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            

            val := rdb.ZRevRangeWithScores(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp78040 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp78040 = append(tmp78040, term)
        }
        term := ast.ArrayTerm(tmp78040...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZREVRANGEBYSCORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrevrangebyscore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}),types.NewStaticProperty("Max", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
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
            
            var v1 *redis.ZRangeBy
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRevRangeByScore(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp43345 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp43345 = append(tmp43345, term)
        }
        term := ast.ArrayTerm(tmp43345...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZREVRANGEBYLEX(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrevrangebylex",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}),types.NewStaticProperty("Max", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
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
            
            var v1 *redis.ZRangeBy
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRevRangeByLex(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp66186 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp66186 = append(tmp66186, term)
        }
        term := ast.ArrayTerm(tmp66186...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZREVRANGEBYSCOREWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zrevrangebyscorewithscores",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Min", types.String{}),types.NewStaticProperty("Max", types.String{}),types.NewStaticProperty("Offset", types.Number{}),types.NewStaticProperty("Count", types.Number{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            
            var v1 *redis.ZRangeBy
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZRevRangeByScoreWithScores(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp40846 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp40846 = append(tmp40846, term)
        }
        term := ast.ArrayTerm(tmp40846...)
        
        return term,  nil
        
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
            

            val := rdb.ZRevRank(p.redisContext,v0,v1)
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
            

            val := rdb.ZScore(p.redisContext,v0,v1)
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
            Name: "redis.zunionstore",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})),types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.Number{}),
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
            
            var v1 *redis.ZStore
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.ZUnionStore(p.redisContext,v0,v1)
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
            Name: "redis.zunion",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})),types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
            Memoize: true,
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
            

            val := rdb.ZUnion(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp94449 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp94449 = append(tmp94449, term)
        }
        term := ast.ArrayTerm(tmp94449...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZUNIONWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zunionwithscores",
            Decl: types.NewFunction(types.Args(types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Keys", types.NewArray([]types.Type{}, types.String{})),types.NewStaticProperty("Weights", types.NewArray([]types.Type{}, types.Number{})),types.NewStaticProperty("Aggregate", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
            Memoize: true,
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
            

            val := rdb.ZUnionWithScores(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp28578 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp28578 = append(tmp28578, term)
        }
        term := ast.ArrayTerm(tmp28578...)
        
        return term,  nil
        
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
            

            val := rdb.ZRandMember(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp80263 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp80263 = append(tmp80263, term)
        }
        term := ast.ArrayTerm(tmp80263...)
        
        return term,  nil
        
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
            

            val := rdb.ZDiff(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp11916 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp11916 = append(tmp11916, term)
        }
        term := ast.ArrayTerm(tmp11916...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerZDIFFWITHSCORES(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.zdiffwithscores",
            Decl: types.NewFunction(types.Args(types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Score", types.Number{}),types.NewStaticProperty("Member", types.Any{})}, nil))),
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
            

            val := rdb.ZDiffWithScores(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp95321 []*ast.Term
        for _, v := range r0 {
            
            score := ast.FloatNumberTerm(float64(v.Score))
            
            
            member := ast.NullTerm()
            if s, ok := v.Member.(string); ok {
                member = ast.StringTerm(s)
            }
            
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Score"), score},[2]*ast.Term{ast.StringTerm("Member"), member})
        
            tmp95321 = append(tmp95321, term)
        }
        term := ast.ArrayTerm(tmp95321...)
        
        return term,  nil
        
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
            

            val := rdb.ZDiffStore(p.redisContext,v0,v1...)
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
            

            val := rdb.PFAdd(p.redisContext,v0,utils.Conva(v1)...)
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
            

            val := rdb.PFCount(p.redisContext,v0...)
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
            

            val := rdb.PFMerge(p.redisContext,v0,v1...)
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
            

            val := rdb.ClientKill(p.redisContext,v0)
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
            

            val := rdb.ClientKillByFilter(p.redisContext,v0...)
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
            

            val := rdb.ClientPause(p.redisContext,v0)
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
            

            val := rdb.ConfigGet(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp95777 []*ast.Term
        for _, v := range r0 {
            
            term := ast.NullTerm()
            if s, ok := v.(string); ok {
                term = ast.StringTerm(s)
            }
            
            tmp95777 = append(tmp95777, term)
        }
        term := ast.ArrayTerm(tmp95777...)
        
        return term,  nil
        
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
            

            val := rdb.ConfigSet(p.redisContext,v0,v1)
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
            

            val := rdb.Info(p.redisContext,v0...)
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
            

            val := rdb.SlaveOf(p.redisContext,v0,v1)
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
            

            val := rdb.DebugObject(p.redisContext,v0)
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
            

            val := rdb.MemoryUsage(p.redisContext,v0,v1...)
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
            

            val := rdb.Eval(p.redisContext,v0,v1,utils.Conva(v2)...)
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
            

            val := rdb.EvalSha(p.redisContext,v0,v1,utils.Conva(v2)...)
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
            

            val := rdb.ScriptExists(p.redisContext,v0...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp46423 []*ast.Term
        for _, v := range r0 {
            term := ast.BooleanTerm(v)
            tmp46423 = append(tmp46423, term)
        }
        term := ast.ArrayTerm(tmp46423...)
        
        return term,  nil
        
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
            

            val := rdb.ScriptLoad(p.redisContext,v0)
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
            

            val := rdb.Publish(p.redisContext,v0,utils.Conv(v1))
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
            

            val := rdb.PubSubChannels(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp61717 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp61717 = append(tmp61717, term)
        }
        term := ast.ArrayTerm(tmp61717...)
        
        return term,  nil
        
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
            Name: "redis.clusterslots",
            Decl: types.NewFunction(types.Args(), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Start", types.Number{}),types.NewStaticProperty("End", types.Number{}),types.NewStaticProperty("Nodes", types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("ID", types.String{}),types.NewStaticProperty("Addr", types.String{})}, nil)))}, nil))),
            Memoize: true,
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
                
        
        var tmp88665 []*ast.Term
        for _, v := range r0 {
            
            start := ast.IntNumberTerm(int(v.Start))
            
            end := ast.IntNumberTerm(int(v.End))
            
            
        var tmp45187 []*ast.Term
        for _, v := range v.Nodes {
            
            id := ast.StringTerm(v.ID)
            
            addr := ast.StringTerm(v.Addr)
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("ID"), id},[2]*ast.Term{ast.StringTerm("Addr"), addr})
        
            tmp45187 = append(tmp45187, term)
        }
        nodes := ast.ArrayTerm(tmp45187...)
        
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Start"), start},[2]*ast.Term{ast.StringTerm("End"), end},[2]*ast.Term{ast.StringTerm("Nodes"), nodes})
        
            tmp88665 = append(tmp88665, term)
        }
        term := ast.ArrayTerm(tmp88665...)
        
        return term,  nil
        
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
            

            val := rdb.ClusterMeet(p.redisContext,v0,v1)
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
            

            val := rdb.ClusterForget(p.redisContext,v0)
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
            

            val := rdb.ClusterReplicate(p.redisContext,v0)
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
            

            val := rdb.ClusterKeySlot(p.redisContext,v0)
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
            

            val := rdb.ClusterGetKeysInSlot(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp82903 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp82903 = append(tmp82903, term)
        }
        term := ast.ArrayTerm(tmp82903...)
        
        return term,  nil
        
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
            

            val := rdb.ClusterCountFailureReports(p.redisContext,v0)
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
            

            val := rdb.ClusterCountKeysInSlot(p.redisContext,v0)
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
            

            val := rdb.ClusterDelSlots(p.redisContext,v0...)
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
            

            val := rdb.ClusterDelSlotsRange(p.redisContext,v0,v1)
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
            

            val := rdb.ClusterSlaves(p.redisContext,v0)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp67311 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp67311 = append(tmp67311, term)
        }
        term := ast.ArrayTerm(tmp67311...)
        
        return term,  nil
        
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
            

            val := rdb.ClusterAddSlots(p.redisContext,v0...)
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
            

            val := rdb.ClusterAddSlotsRange(p.redisContext,v0,v1)
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
            Name: "redis.geoadd",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}),types.NewStaticProperty("Longitude", types.Number{}),types.NewStaticProperty("Latitude", types.Number{}),types.NewStaticProperty("Dist", types.Number{}),types.NewStaticProperty("GeoHash", types.Number{})}, nil))), types.Number{}),
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
            
            var v1 []*redis.GeoLocation
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.GeoAdd(p.redisContext,v0,v1...)
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
            Name: "redis.geopos",
            Decl: types.NewFunction(types.Args(types.String{},types.NewArray([]types.Type{}, types.String{})), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Longitude", types.Number{}),types.NewStaticProperty("Latitude", types.Number{})}, nil))),
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
            

            val := rdb.GeoPos(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp31353 []*ast.Term
        for _, v := range r0 {
            if v == nil {
                tmp31353 = append(tmp31353, ast.NullTerm())
                continue
            }
            
            
            longitude := ast.FloatNumberTerm(float64(v.Longitude))
            
            latitude := ast.FloatNumberTerm(float64(v.Latitude))
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Longitude"), longitude},[2]*ast.Term{ast.StringTerm("Latitude"), latitude})
        
            tmp31353 = append(tmp31353, term)
        }
        term := ast.ArrayTerm(tmp31353...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerGEORADIUS(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.georadius",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}),types.NewStaticProperty("Unit", types.String{}),types.NewStaticProperty("WithCoord", types.Boolean{}),types.NewStaticProperty("WithDist", types.Boolean{}),types.NewStaticProperty("WithGeoHash", types.Boolean{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Sort", types.String{}),types.NewStaticProperty("Store", types.String{}),types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}),types.NewStaticProperty("Longitude", types.Number{}),types.NewStaticProperty("Latitude", types.Number{}),types.NewStaticProperty("Dist", types.Number{}),types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
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
            
            var v2 float64
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            
            var v3 *redis.GeoRadiusQuery
            if err := ast.As(terms[3].Value, &v3); err != nil {
                return nil, err
            }
            

            val := rdb.GeoRadius(p.redisContext,v0,v1,v2,v3)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp53857 []*ast.Term
        for _, v := range r0 {
            
            name := ast.StringTerm(v.Name)
            
            longitude := ast.FloatNumberTerm(float64(v.Longitude))
            
            latitude := ast.FloatNumberTerm(float64(v.Latitude))
            
            dist := ast.FloatNumberTerm(float64(v.Dist))
            
            geohash := ast.IntNumberTerm(int(v.GeoHash))
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), name},[2]*ast.Term{ast.StringTerm("Longitude"), longitude},[2]*ast.Term{ast.StringTerm("Latitude"), latitude},[2]*ast.Term{ast.StringTerm("Dist"), dist},[2]*ast.Term{ast.StringTerm("GeoHash"), geohash})
        
            tmp53857 = append(tmp53857, term)
        }
        term := ast.ArrayTerm(tmp53857...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerGEORADIUSSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.georadiusstore",
            Decl: types.NewFunction(types.Args(types.String{},types.Number{},types.Number{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}),types.NewStaticProperty("Unit", types.String{}),types.NewStaticProperty("WithCoord", types.Boolean{}),types.NewStaticProperty("WithDist", types.Boolean{}),types.NewStaticProperty("WithGeoHash", types.Boolean{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Sort", types.String{}),types.NewStaticProperty("Store", types.String{}),types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.Number{}),
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
            
            var v2 float64
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            
            var v3 *redis.GeoRadiusQuery
            if err := ast.As(terms[3].Value, &v3); err != nil {
                return nil, err
            }
            

            val := rdb.GeoRadiusStore(p.redisContext,v0,v1,v2,v3)
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
            Name: "redis.georadiusbymember",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}),types.NewStaticProperty("Unit", types.String{}),types.NewStaticProperty("WithCoord", types.Boolean{}),types.NewStaticProperty("WithDist", types.Boolean{}),types.NewStaticProperty("WithGeoHash", types.Boolean{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Sort", types.String{}),types.NewStaticProperty("Store", types.String{}),types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}),types.NewStaticProperty("Longitude", types.Number{}),types.NewStaticProperty("Latitude", types.Number{}),types.NewStaticProperty("Dist", types.Number{}),types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
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
            
            var v2 *redis.GeoRadiusQuery
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            

            val := rdb.GeoRadiusByMember(p.redisContext,v0,v1,v2)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp74706 []*ast.Term
        for _, v := range r0 {
            
            name := ast.StringTerm(v.Name)
            
            longitude := ast.FloatNumberTerm(float64(v.Longitude))
            
            latitude := ast.FloatNumberTerm(float64(v.Latitude))
            
            dist := ast.FloatNumberTerm(float64(v.Dist))
            
            geohash := ast.IntNumberTerm(int(v.GeoHash))
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), name},[2]*ast.Term{ast.StringTerm("Longitude"), longitude},[2]*ast.Term{ast.StringTerm("Latitude"), latitude},[2]*ast.Term{ast.StringTerm("Dist"), dist},[2]*ast.Term{ast.StringTerm("GeoHash"), geohash})
        
            tmp74706 = append(tmp74706, term)
        }
        term := ast.ArrayTerm(tmp74706...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerGEORADIUSBYMEMBERSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.georadiusbymemberstore",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Radius", types.Number{}),types.NewStaticProperty("Unit", types.String{}),types.NewStaticProperty("WithCoord", types.Boolean{}),types.NewStaticProperty("WithDist", types.Boolean{}),types.NewStaticProperty("WithGeoHash", types.Boolean{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("Sort", types.String{}),types.NewStaticProperty("Store", types.String{}),types.NewStaticProperty("StoreDist", types.String{})}, nil)), types.Number{}),
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
            
            var v2 *redis.GeoRadiusQuery
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            

            val := rdb.GeoRadiusByMemberStore(p.redisContext,v0,v1,v2)
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
            Name: "redis.geosearch",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Member", types.String{}),types.NewStaticProperty("Longitude", types.Number{}),types.NewStaticProperty("Latitude", types.Number{}),types.NewStaticProperty("Radius", types.Number{}),types.NewStaticProperty("RadiusUnit", types.String{}),types.NewStaticProperty("BoxWidth", types.Number{}),types.NewStaticProperty("BoxHeight", types.Number{}),types.NewStaticProperty("BoxUnit", types.String{}),types.NewStaticProperty("Sort", types.String{}),types.NewStaticProperty("Count", types.Number{}),types.NewStaticProperty("CountAny", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.String{})),
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
            
            var v1 *redis.GeoSearchQuery
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.GeoSearch(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp77422 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp77422 = append(tmp77422, term)
        }
        term := ast.ArrayTerm(tmp77422...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerGEOSEARCHLOCATION(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.geosearchlocation",
            Decl: types.NewFunction(types.Args(types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("WithCoord", types.Boolean{}),types.NewStaticProperty("WithDist", types.Boolean{}),types.NewStaticProperty("WithHash", types.Boolean{})}, nil)), types.NewArray([]types.Type{}, types.NewObject([]*types.StaticProperty{types.NewStaticProperty("Name", types.String{}),types.NewStaticProperty("Longitude", types.Number{}),types.NewStaticProperty("Latitude", types.Number{}),types.NewStaticProperty("Dist", types.Number{}),types.NewStaticProperty("GeoHash", types.Number{})}, nil))),
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
            
            var v1 *redis.GeoSearchLocationQuery
            if err := ast.As(terms[1].Value, &v1); err != nil {
                return nil, err
            }
            

            val := rdb.GeoSearchLocation(p.redisContext,v0,v1)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp07256 []*ast.Term
        for _, v := range r0 {
            
            name := ast.StringTerm(v.Name)
            
            longitude := ast.FloatNumberTerm(float64(v.Longitude))
            
            latitude := ast.FloatNumberTerm(float64(v.Latitude))
            
            dist := ast.FloatNumberTerm(float64(v.Dist))
            
            geohash := ast.IntNumberTerm(int(v.GeoHash))
            
        term := ast.ObjectTerm([2]*ast.Term{ast.StringTerm("Name"), name},[2]*ast.Term{ast.StringTerm("Longitude"), longitude},[2]*ast.Term{ast.StringTerm("Latitude"), latitude},[2]*ast.Term{ast.StringTerm("Dist"), dist},[2]*ast.Term{ast.StringTerm("GeoHash"), geohash})
        
            tmp07256 = append(tmp07256, term)
        }
        term := ast.ArrayTerm(tmp07256...)
        
        return term,  nil
        
            default:
                return nil, err
            }
        },
    )
}


func registerGEOSEARCHSTORE(p *redisPlugin) {
    rego.RegisterBuiltinDyn(
        &rego.Function{
            Name: "redis.geosearchstore",
            Decl: types.NewFunction(types.Args(types.String{},types.String{},types.NewObject([]*types.StaticProperty{types.NewStaticProperty("StoreDist", types.Boolean{})}, nil)), types.Number{}),
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
            
            var v2 *redis.GeoSearchStoreQuery
            if err := ast.As(terms[2].Value, &v2); err != nil {
                return nil, err
            }
            

            val := rdb.GeoSearchStore(p.redisContext,v0,v1,v2)
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
            

            val := rdb.GeoDist(p.redisContext,v0,v1,v2,v3)
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
            

            val := rdb.GeoHash(p.redisContext,v0,v1...)
            r0 := val.Val()
            switch err := val.Err(); err {
            case redis.Nil:
                return ast.NullTerm(), nil
            case nil:
                
        
        var tmp59937 []*ast.Term
        for _, v := range r0 {
            term := ast.StringTerm(v)
            tmp59937 = append(tmp59937, term)
        }
        term := ast.ArrayTerm(tmp59937...)
        
        return term,  nil
        
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
    registerXGROUPCREATE(p)
    registerXGROUPCREATEMKSTREAM(p)
    registerXGROUPSETID(p)
    registerXGROUPDESTROY(p)
    registerXGROUPCREATECONSUMER(p)
    registerXGROUPDELCONSUMER(p)
    registerXACK(p)
    registerXCLAIMJUSTID(p)
    registerXAUTOCLAIMJUSTID(p)
    registerXTRIM(p)
    registerXTRIMAPPROX(p)
    registerXTRIMMAXLEN(p)
    registerXTRIMMAXLENAPPROX(p)
    registerXTRIMMINID(p)
    registerXTRIMMINIDAPPROX(p)
    registerXINFOGROUPS(p)
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

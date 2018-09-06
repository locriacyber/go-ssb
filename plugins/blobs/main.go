package blobs

import (
	"context"

	"github.com/cryptix/go/logging"
	"github.com/pkg/errors"

	"go.cryptoscope.co/muxrpc"

	"go.cryptoscope.co/luigi"
	"go.cryptoscope.co/sbot"
)

/*
blobs manifest.json except:
"get": "source",
"add": "sink",
"rm": "async",
"ls": "source",
"has": "async",
"want": "async",
"createWants": "source"

"size": "async",
"getSlice": "source",
"meta": "async",
"push": "async",
"changes": "source",
*/

var (
	_      sbot.Plugin = plugin{} // compile-time type check
	method             = muxrpc.Method{"blobs"}
)

func checkAndLog(log logging.Interface, err error) {
	if err != nil {
		if err := logging.LogPanicWithStack(log, "checkAndLog", err); err != nil {
			panic(err)
		}
	}
}

func New(log logging.Interface, bs sbot.BlobStore, wm sbot.WantManager) sbot.Plugin {
	rootHdlr := muxrpc.HandlerMux{}

	rootHdlr.Register(muxrpc.Method{"blobs", "get"}, getHandler{
		log: log,
		bs:  bs,
	})
	rootHdlr.Register(muxrpc.Method{"blobs", "add"}, addHandler{
		log: log,
		bs:  bs,
	})
	rootHdlr.Register(muxrpc.Method{"blobs", "list"}, listHandler{
		log: log,
		bs:  bs,
	})
	rootHdlr.Register(muxrpc.Method{"blobs", "has"}, hasHandler{
		log: log,
		bs:  bs,
	})
	rootHdlr.Register(muxrpc.Method{"blobs", "rm"}, rmHandler{
		log: log,
		bs:  bs,
	})
	rootHdlr.Register(muxrpc.Method{"blobs", "want"}, wantHandler{
		log: log,
		wm:  wm,
	})
	rootHdlr.Register(muxrpc.Method{"blobs", "createWants"}, &createWantsHandler{
		log:     log,
		bs:      bs,
		wm:      wm,
		sources: make(map[string]luigi.Source),
	})

	return plugin{
		h:   &rootHdlr,
		log: log,
	}
}

type plugin struct {
	h   muxrpc.Handler
	log logging.Interface
}

func (plugin) Name() string { return "blobs" }

func (plugin) Method() muxrpc.Method {
	return method
}

func (p plugin) Handler() muxrpc.Handler {
	return p.h
}

func (plugin) WrapEndpoint(edp muxrpc.Endpoint) interface{} {
	return endpoint{edp}
}

type endpoint struct {
	edp muxrpc.Endpoint
}

func (edp endpoint) Add(ctx context.Context) (sbot.MessageRef, error) {
	return sbot.MessageRef{}, errors.New("not implemented yet")
}
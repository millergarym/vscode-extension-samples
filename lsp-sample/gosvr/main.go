package main

import (
	"context"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/golangq/q"
	"golang.org/x/tools/jsonrpc2"
	"golang.org/x/tools/lsp/protocol"
)

type Server struct {
}

func (svr *Server) Initialize(ctx context.Context, req *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	q.Q(req)
	// type ServerCapabilities struct {
	// 	InnerServerCapabilities
	// 	ImplementationServerCapabilities
	// 	TypeDefinitionServerCapabilities
	// 	WorkspaceFoldersServerCapabilities
	// 	ColorServerCapabilities
	// 	FoldingRangeServerCapabilities
	// 	DeclarationServerCapabilities
	// 	SelectionRangeServerCapabilities
	// }
	q.Q(spew.Sdump(req))
	textDocumentSyncKind := protocol.Full
	ret := &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			InnerServerCapabilities: protocol.InnerServerCapabilities{
				CodeActionProvider: true,
				CompletionProvider: &protocol.CompletionOptions{
					TriggerCharacters: []string{"."},
				},
				DefinitionProvider:              true,
				DocumentFormattingProvider:      true,
				DocumentRangeFormattingProvider: true,
				DocumentSymbolProvider:          true,
				HoverProvider:                   true,
				DocumentHighlightProvider:       true,
				SignatureHelpProvider: &protocol.SignatureHelpOptions{
					TriggerCharacters: []string{"(", ","},
				},
				TextDocumentSync: &protocol.TextDocumentSyncOptions{
					Change:    textDocumentSyncKind,
					OpenClose: true,
				},
			},
			TypeDefinitionServerCapabilities: protocol.TypeDefinitionServerCapabilities{
				TypeDefinitionProvider: true,
			},
		},
	}

	// ret := &protocol.InitializeResult{
	// 	Capabilities: protocol.ServerCapabilities{},
	// 	Custom:       make(map[string]interface{}),
	// }
	return ret, nil
}
func (svr *Server) Initialized(ctx context.Context, req *protocol.InitializedParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) Shutdown(context.Context) error {
	q.Q("shutdown")
	return nil
}
func (svr *Server) Exit(context.Context) error {
	q.Q("exit")
	return nil
}
func (svr *Server) DidChangeWorkspaceFolders(ctx context.Context, req *protocol.DidChangeWorkspaceFoldersParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) DidChangeConfiguration(ctx context.Context, req *protocol.DidChangeConfigurationParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) DidChangeWatchedFiles(ctx context.Context, req *protocol.DidChangeWatchedFilesParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) Symbols(ctx context.Context, req *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) ExecuteCommand(ctx context.Context, req *protocol.ExecuteCommandParams) (interface{}, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DidOpen(ctx context.Context, req *protocol.DidOpenTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) DidChange(ctx context.Context, req *protocol.DidChangeTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) WillSave(ctx context.Context, req *protocol.WillSaveTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) WillSaveWaitUntil(ctx context.Context, req *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DidSave(ctx context.Context, req *protocol.DidSaveTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) DidClose(ctx context.Context, req *protocol.DidCloseTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *Server) Completion(ctx context.Context, req *protocol.CompletionParams) (*protocol.CompletionList, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) CompletionResolve(ctx context.Context, req *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) Hover(ctx context.Context, req *protocol.TextDocumentPositionParams) (*protocol.Hover, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) SignatureHelp(ctx context.Context, req *protocol.TextDocumentPositionParams) (*protocol.SignatureHelp, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) Definition(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) TypeDefinition(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) Implementation(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) References(ctx context.Context, req *protocol.ReferenceParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DocumentHighlight(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.DocumentHighlight, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DocumentSymbol(ctx context.Context, req *protocol.DocumentSymbolParams) ([]protocol.DocumentSymbol, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) CodeAction(ctx context.Context, req *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) CodeLens(ctx context.Context, req *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) CodeLensResolve(ctx context.Context, req *protocol.CodeLens) (*protocol.CodeLens, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DocumentLink(ctx context.Context, req *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DocumentLinkResolve(ctx context.Context, req *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) DocumentColor(ctx context.Context, req *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) ColorPresentation(ctx context.Context, req *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) Formatting(ctx context.Context, req *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) RangeFormatting(ctx context.Context, req *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) OnTypeFormatting(ctx context.Context, req *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) Rename(ctx context.Context, req *protocol.RenameParams) ([]protocol.WorkspaceEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *Server) FoldingRanges(ctx context.Context, req *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	q.Q(req)
	return nil, nil
}

func main() {
	q.Q("hw")
	q.Q(os.Args)

	stream := jsonrpc2.NewHeaderStream(os.Stdin, os.Stdout)
	srv := &Server{}
	connLSP, _, _ := protocol.NewServer(stream, srv)
	ctx := context.Background()
	q.Q(connLSP.Run(ctx))
}

package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *App) domReady(ctx context.Context) {
	runtime.WindowShow(ctx)
}

func (b *App) No() {
	runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "ты чё",
		Message: "пидора ответ",
	})
}

package worldedit

import (
	"sync"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	worldEdit "github.com/df-plus/weLib"
	"github.com/go-gl/mathgl/mgl64"
)

var cooldowns map[*player.Player]time.Time
var cooldownsMu sync.RWMutex

func setcooldown(p *player.Player) {
	cooldownsMu.Lock()
	defer cooldownsMu.Unlock()
	cooldowns[p] = time.Now().Add(time.Second / 2)
}

func cooldown(p *player.Player) bool {
	cooldownsMu.RLock()
	defer cooldownsMu.RUnlock()
	if c, ok := cooldowns[p]; ok {
		return c.Before(time.Now())
	}
	return false
}

type Handler struct {
	player.NopHandler
	p *player.Player
	m *worldEdit.Manager
}

func NewHandler(p *player.Player, m *worldEdit.Manager) *Handler {
	return &Handler{
		p: p,
		m: m,
	}
}
func (h *Handler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {

	if wandItem.Comparable(item.Stack{}) {
		return
	}
	heldItem, _ := h.p.HeldItems()
	if heldItem.Comparable(Wand()) {
		if !cooldown(h.p) {
			setcooldown(h.p)
			ctx.Cancel()
			h.m.SetPos1(h.p, clickPos)
		}
	}
}

func (h *Handler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drop *[]item.Stack) {
	if wandItem.Comparable(item.Stack{}) {
		return
	}
	heldItem, _ := h.p.HeldItems()
	if heldItem.Comparable(Wand()) {
		ctx.Cancel()
		h.m.SetPos2(h.p, pos.Vec3())
	}
}

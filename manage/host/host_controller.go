package host

import (
	"context"

	"github.com/valyala/fasthttp"
	"github.com/xiaofengshuyu/vpn-manager/manage/common"
)

// ConfigHandler is config handler
type ConfigHandler struct {
	common.BaseHandler
	Service ConfigService
}

// GetVPNConfig get personal vpn config
func (h *ConfigHandler) GetVPNConfig(ctx *fasthttp.RequestCtx) {
	user, ok := common.GlobalSession.GetUser(ctx.ID())
	if !ok {
		h.WriteJSON(ctx, nil, common.NewNotLoginError())
		return
	}

	conf, err := h.Service.GetVPNConfig(context.Background(), &user)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	h.WriteJSON(ctx, vpnConfigResponseEncode(conf), nil)
	return
}

// GetHostList get host list that user can used
func (h *ConfigHandler) GetHostList(ctx *fasthttp.RequestCtx) {
	user, ok := common.GlobalSession.GetUser(ctx.ID())
	if !ok {
		h.WriteJSON(ctx, nil, common.NewNotLoginError())
		return
	}

	conf, err := h.Service.GetVPNConfig(context.Background(), &user)
	if err != nil {
		h.WriteJSON(ctx, nil, err)
		return
	}
	data := vpnConfigResponseEncode(conf)
	h.WriteJSON(ctx, data.Hosts, nil)
	return
}

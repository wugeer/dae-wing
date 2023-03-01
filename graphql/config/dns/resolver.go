/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, v2rayA Organization <team@v2raya.org>
 */

package dns

import (
	"github.com/v2rayA/dae-wing/graphql/config"
	"github.com/v2rayA/dae-wing/graphql/config/routing"
	"github.com/v2rayA/dae/common"
	daeConfig "github.com/v2rayA/dae/config"
	"github.com/v2rayA/dae/pkg/config_parser"
)

type Resolver struct {
	*daeConfig.Dns
}

func (r *Resolver) Upstream() (rs []*config.ParamResolver) {
	for _, upstream := range r.Dns.Upstream {
		tag, afterTag := common.GetTagFromLinkLikePlaintext(string(upstream))
		rs = append(rs, &config.ParamResolver{Param: &config_parser.Param{
			Key: tag,
			Val: afterTag,
		}})
	}
	return rs
}

func (r *Resolver) Routing() *RoutingResolver {
	return &RoutingResolver{DnsRouting: &r.Dns.Routing}
}

type RoutingResolver struct {
	*daeConfig.DnsRouting
}

func (r *RoutingResolver) Request() *routing.Resolver {
	return &routing.Resolver{Routing: &daeConfig.Routing{
		Rules:    r.DnsRouting.Request.Rules,
		Fallback: r.DnsRouting.Request.Fallback,
	}}
}
func (r *RoutingResolver) Response() *routing.Resolver {
	return &routing.Resolver{Routing: &daeConfig.Routing{
		Rules:    r.DnsRouting.Response.Rules,
		Fallback: r.DnsRouting.Response.Fallback,
	}}
}

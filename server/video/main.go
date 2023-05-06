package main

func main() {
	//addr := "myserviceName:12345"
	//
	//// 使用 scheme dns:/// ，这样就会使用dns解析到 server 端的 pod ip
	//
	//
	//conn, err := grpc.Dial(
	//	"dns:///"+addr,
	//	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), //grpc.WithBalancerName(roundrobin.Name) 是旧版本写法，已废弃
	//	grpc.WithInsecure(),
	//	grpc.WithBlock(),
	//	grpc.WithResolvers(&myDnsBuilder{}),
	//)

}

type myDnsBuilder struct{}

func (m *myDnsBuilder) Scheme() string {
	return "dns"
}

//func (m *myDnsBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
//	host, port, err := parseTarget(target.Endpoint(), defaultPort)
//	if err != nil {
//		return nil, err
//	}
//
//	// IP address.
//	if ipAddr, ok := formatIP(host); ok {
//		addr := []resolver.Address{{Addr: ipAddr + ":" + port}}
//		cc.UpdateState(resolver.State{Addresses: addr})
//		return deadResolver{}, nil
//	}
//
//	// DNS address (non-IP).
//	ctx, cancel := context.WithCancel(context.Background())
//	d := &dnsResolver{
//		host:                 host,
//		port:                 port,
//		ctx:                  ctx,
//		cancel:               cancel,
//		cc:                   cc,
//		rn:                   make(chan struct{}, 1),
//		disableServiceConfig: opts.DisableServiceConfig,
//	}
//
//	if target.URL.Host == "" {
//		d.resolver = defaultResolver
//	} else {
//		d.resolver, err = customAuthorityResolver(target.URL.Host)
//		if err != nil {
//			return nil, err
//		}
//	}
//
//	d.wg.Add(1)
//	go d.watcher()
//	return d, nil
//}

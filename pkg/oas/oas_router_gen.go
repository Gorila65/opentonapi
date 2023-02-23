// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"net/http"
	"strings"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [2]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/v2/"
			if l := len("/v2/"); len(elem) >= l && elem[0:l] == "/v2/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "accounts/"
				if l := len("accounts/"); len(elem) >= l && elem[0:l] == "accounts/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "account_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleGetAccountRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'd': // Prefix: "dns/backresolve"
						if l := len("dns/backresolve"); len(elem) >= l && elem[0:l] == "dns/backresolve" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleDnsBackResolveRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'e': // Prefix: "events"
						if l := len("events"); len(elem) >= l && elem[0:l] == "events" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetEventsByAccountRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'j': // Prefix: "jettons"
						if l := len("jettons"); len(elem) >= l && elem[0:l] == "jettons" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetJettonsBalancesRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'n': // Prefix: "ntfs"
						if l := len("ntfs"); len(elem) >= l && elem[0:l] == "ntfs" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetNftItemsByOwnerRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 's': // Prefix: "subscriptions"
						if l := len("subscriptions"); len(elem) >= l && elem[0:l] == "subscriptions" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetSubscriptionsByAccountRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 't': // Prefix: "traces"
						if l := len("traces"); len(elem) >= l && elem[0:l] == "traces" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetTracesByAccountRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'b': // Prefix: "blockchain/"
				if l := len("blockchain/"); len(elem) >= l && elem[0:l] == "blockchain/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "accounts/"
					if l := len("accounts/"); len(elem) >= l && elem[0:l] == "accounts/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "account_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleGetRawAccountRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'm': // Prefix: "methods/"
							if l := len("methods/"); len(elem) >= l && elem[0:l] == "methods/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "method_name"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleExecGetMethodRequest([2]string{
										args[0],
										args[1],
									}, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						case 't': // Prefix: "transactions"
							if l := len("transactions"); len(elem) >= l && elem[0:l] == "transactions" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleGetAccountTransactionsRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				case 'b': // Prefix: "blocks/"
					if l := len("blocks/"); len(elem) >= l && elem[0:l] == "blocks/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "block_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleGetBlockRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/transactions"
						if l := len("/transactions"); len(elem) >= l && elem[0:l] == "/transactions" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetTransactionsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				case 'c': // Prefix: "config"
					if l := len("config"); len(elem) >= l && elem[0:l] == "config" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetConfigRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'm': // Prefix: "m"
					if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "asterchain-head"
						if l := len("asterchain-head"); len(elem) >= l && elem[0:l] == "asterchain-head" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetMasterchainHeadRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'e': // Prefix: "essage"
						if l := len("essage"); len(elem) >= l && elem[0:l] == "essage" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleSendMessageRequest([0]string{}, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/emulate"
							if l := len("/emulate"); len(elem) >= l && elem[0:l] == "/emulate" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "POST":
									s.handleEmulateMessageRequest([0]string{}, w, r)
								default:
									s.notAllowed(w, r, "POST")
								}

								return
							}
						}
					}
				case 't': // Prefix: "transactions/"
					if l := len("transactions/"); len(elem) >= l && elem[0:l] == "transactions/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "transaction_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetTransactionRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'v': // Prefix: "validators"
					if l := len("validators"); len(elem) >= l && elem[0:l] == "validators" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetValidatorsRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'd': // Prefix: "dns/"
				if l := len("dns/"); len(elem) >= l && elem[0:l] == "dns/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "auctions"
					if l := len("auctions"); len(elem) >= l && elem[0:l] == "auctions" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleGetAllAuctionsRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
				// Param: "domain_name"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "bids"
						if l := len("bids"); len(elem) >= l && elem[0:l] == "bids" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetDomainBidsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					case 'r': // Prefix: "resolve"
						if l := len("resolve"); len(elem) >= l && elem[0:l] == "resolve" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleDnsResolveRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'e': // Prefix: "events/"
				if l := len("events/"); len(elem) >= l && elem[0:l] == "events/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "event_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleGetEventRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'j': // Prefix: "jettons/"
				if l := len("jettons/"); len(elem) >= l && elem[0:l] == "jettons/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "account_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleGetJettonInfoRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'n': // Prefix: "nfts/"
				if l := len("nfts/"); len(elem) >= l && elem[0:l] == "nfts/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "collections"
					if l := len("collections"); len(elem) >= l && elem[0:l] == "collections" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleGetNftCollectionsRequest([0]string{}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "account_id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleGetNftCollectionRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
				// Param: "account_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleGetNftItemByAddressRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 's': // Prefix: "stacking/"
				if l := len("stacking/"); len(elem) >= l && elem[0:l] == "stacking/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'n': // Prefix: "nominator/"
					if l := len("nominator/"); len(elem) >= l && elem[0:l] == "nominator/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "account_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/pools"
						if l := len("/pools"); len(elem) >= l && elem[0:l] == "/pools" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handlePoolsByNominatorsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				case 'p': // Prefix: "pool/"
					if l := len("pool/"); len(elem) >= l && elem[0:l] == "pool/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "account_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleStackingPoolInfoRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 't': // Prefix: "traces/"
				if l := len("traces/"); len(elem) >= l && elem[0:l] == "traces/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "trace_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleGetTraceRequest([1]string{
							args[0],
						}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	count       int
	args        [2]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
func (s *Server) FindRoute(method, path string) (r Route, _ bool) {
	var (
		args = [2]string{}
		elem = path
	)
	r.args = args
	if elem == "" {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/v2/"
			if l := len("/v2/"); len(elem) >= l && elem[0:l] == "/v2/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "accounts/"
				if l := len("accounts/"); len(elem) >= l && elem[0:l] == "accounts/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "account_id"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "GetAccount"
						r.operationID = "getAccount"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'd': // Prefix: "dns/backresolve"
						if l := len("dns/backresolve"); len(elem) >= l && elem[0:l] == "dns/backresolve" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: DnsBackResolve
								r.name = "DnsBackResolve"
								r.operationID = "dnsBackResolve"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'e': // Prefix: "events"
						if l := len("events"); len(elem) >= l && elem[0:l] == "events" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetEventsByAccount
								r.name = "GetEventsByAccount"
								r.operationID = "getEventsByAccount"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'j': // Prefix: "jettons"
						if l := len("jettons"); len(elem) >= l && elem[0:l] == "jettons" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetJettonsBalances
								r.name = "GetJettonsBalances"
								r.operationID = "getJettonsBalances"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'n': // Prefix: "ntfs"
						if l := len("ntfs"); len(elem) >= l && elem[0:l] == "ntfs" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetNftItemsByOwner
								r.name = "GetNftItemsByOwner"
								r.operationID = "getNftItemsByOwner"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 's': // Prefix: "subscriptions"
						if l := len("subscriptions"); len(elem) >= l && elem[0:l] == "subscriptions" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetSubscriptionsByAccount
								r.name = "GetSubscriptionsByAccount"
								r.operationID = "getSubscriptionsByAccount"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 't': // Prefix: "traces"
						if l := len("traces"); len(elem) >= l && elem[0:l] == "traces" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetTracesByAccount
								r.name = "GetTracesByAccount"
								r.operationID = "getTracesByAccount"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'b': // Prefix: "blockchain/"
				if l := len("blockchain/"); len(elem) >= l && elem[0:l] == "blockchain/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "accounts/"
					if l := len("accounts/"); len(elem) >= l && elem[0:l] == "accounts/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "account_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "GetRawAccount"
							r.operationID = "getRawAccount"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'm': // Prefix: "methods/"
							if l := len("methods/"); len(elem) >= l && elem[0:l] == "methods/" {
								elem = elem[l:]
							} else {
								break
							}

							// Param: "method_name"
							// Leaf parameter
							args[1] = elem
							elem = ""

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: ExecGetMethod
									r.name = "ExecGetMethod"
									r.operationID = "execGetMethod"
									r.args = args
									r.count = 2
									return r, true
								default:
									return
								}
							}
						case 't': // Prefix: "transactions"
							if l := len("transactions"); len(elem) >= l && elem[0:l] == "transactions" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: GetAccountTransactions
									r.name = "GetAccountTransactions"
									r.operationID = "getAccountTransactions"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				case 'b': // Prefix: "blocks/"
					if l := len("blocks/"); len(elem) >= l && elem[0:l] == "blocks/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "block_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "GetBlock"
							r.operationID = "getBlock"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/transactions"
						if l := len("/transactions"); len(elem) >= l && elem[0:l] == "/transactions" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetTransactions
								r.name = "GetTransactions"
								r.operationID = "getTransactions"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				case 'c': // Prefix: "config"
					if l := len("config"); len(elem) >= l && elem[0:l] == "config" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetConfig
							r.name = "GetConfig"
							r.operationID = "getConfig"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				case 'm': // Prefix: "m"
					if l := len("m"); len(elem) >= l && elem[0:l] == "m" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'a': // Prefix: "asterchain-head"
						if l := len("asterchain-head"); len(elem) >= l && elem[0:l] == "asterchain-head" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetMasterchainHead
								r.name = "GetMasterchainHead"
								r.operationID = "getMasterchainHead"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'e': // Prefix: "essage"
						if l := len("essage"); len(elem) >= l && elem[0:l] == "essage" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = "SendMessage"
								r.operationID = "sendMessage"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/emulate"
							if l := len("/emulate"); len(elem) >= l && elem[0:l] == "/emulate" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "POST":
									// Leaf: EmulateMessage
									r.name = "EmulateMessage"
									r.operationID = "emulateMessage"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}
						}
					}
				case 't': // Prefix: "transactions/"
					if l := len("transactions/"); len(elem) >= l && elem[0:l] == "transactions/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "transaction_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetTransaction
							r.name = "GetTransaction"
							r.operationID = "getTransaction"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
				case 'v': // Prefix: "validators"
					if l := len("validators"); len(elem) >= l && elem[0:l] == "validators" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetValidators
							r.name = "GetValidators"
							r.operationID = "getValidators"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'd': // Prefix: "dns/"
				if l := len("dns/"); len(elem) >= l && elem[0:l] == "dns/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'a': // Prefix: "auctions"
					if l := len("auctions"); len(elem) >= l && elem[0:l] == "auctions" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: GetAllAuctions
							r.name = "GetAllAuctions"
							r.operationID = "getAllAuctions"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
				// Param: "domain_name"
				// Match until "/"
				idx := strings.IndexByte(elem, '/')
				if idx < 0 {
					idx = len(elem)
				}
				args[0] = elem[:idx]
				elem = elem[idx:]

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'b': // Prefix: "bids"
						if l := len("bids"); len(elem) >= l && elem[0:l] == "bids" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetDomainBids
								r.name = "GetDomainBids"
								r.operationID = "getDomainBids"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					case 'r': // Prefix: "resolve"
						if l := len("resolve"); len(elem) >= l && elem[0:l] == "resolve" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: DnsResolve
								r.name = "DnsResolve"
								r.operationID = "dnsResolve"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'e': // Prefix: "events/"
				if l := len("events/"); len(elem) >= l && elem[0:l] == "events/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "event_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: GetEvent
						r.name = "GetEvent"
						r.operationID = "getEvent"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			case 'j': // Prefix: "jettons/"
				if l := len("jettons/"); len(elem) >= l && elem[0:l] == "jettons/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "account_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: GetJettonInfo
						r.name = "GetJettonInfo"
						r.operationID = "getJettonInfo"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			case 'n': // Prefix: "nfts/"
				if l := len("nfts/"); len(elem) >= l && elem[0:l] == "nfts/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'c': // Prefix: "collections"
					if l := len("collections"); len(elem) >= l && elem[0:l] == "collections" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "GetNftCollections"
							r.operationID = "getNftCollections"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						// Param: "account_id"
						// Leaf parameter
						args[0] = elem
						elem = ""

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: GetNftCollection
								r.name = "GetNftCollection"
								r.operationID = "getNftCollection"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
				// Param: "account_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: GetNftItemByAddress
						r.name = "GetNftItemByAddress"
						r.operationID = "getNftItemByAddress"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			case 's': // Prefix: "stacking/"
				if l := len("stacking/"); len(elem) >= l && elem[0:l] == "stacking/" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'n': // Prefix: "nominator/"
					if l := len("nominator/"); len(elem) >= l && elem[0:l] == "nominator/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "account_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case '/': // Prefix: "/pools"
						if l := len("/pools"); len(elem) >= l && elem[0:l] == "/pools" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: PoolsByNominators
								r.name = "PoolsByNominators"
								r.operationID = "poolsByNominators"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				case 'p': // Prefix: "pool/"
					if l := len("pool/"); len(elem) >= l && elem[0:l] == "pool/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "account_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: StackingPoolInfo
							r.name = "StackingPoolInfo"
							r.operationID = "stackingPoolInfo"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
				}
			case 't': // Prefix: "traces/"
				if l := len("traces/"); len(elem) >= l && elem[0:l] == "traces/" {
					elem = elem[l:]
				} else {
					break
				}

				// Param: "trace_id"
				// Leaf parameter
				args[0] = elem
				elem = ""

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: GetTrace
						r.name = "GetTrace"
						r.operationID = "getTrace"
						r.args = args
						r.count = 1
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}

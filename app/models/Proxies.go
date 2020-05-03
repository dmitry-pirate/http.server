package models

import (
	"database/sql"
	"time"
	"vpn_api/app/store"
)

//Proxies Model From DB
type Proxies struct {
	Id                     int            `db:"id"`
	CountryId              int            `db:"country_id"`
	Name                   string         `db:"name"`
	Ip                     string         `db:"ip"`
	Port                   string         `db:"port"`
	CreatedAt              time.Time      `db:"created_at"`
	UpdatedAt              time.Time      `db:"updated_at"`
	OvpnPort               string         `db:"ovpn_port"`
	Endpoint               string         `db:"endpoint"`
	Status                 string         `db:"status"`
	Premium                int            `db:"premium"`
	Streaming              int            `db:"streaming"`
	Schema                 string         `db:"schema"`
	UserVersion            string         `db:"user_version"`
	RemoteIdentifier       sql.NullString `db:"remote_identifier"`
	Psk                    sql.NullString `db:"psk"`
	Region                 sql.NullString `db:"region"`
	MaxNetworkSpeed        int            `db:"max_network_speed"`
	Interface              string         `db:"interface"`
	ServiceState           int            `db:"service_state"`
	ReceiveLoadPercentage  float64        `db:"receive_load_percentage"`
	TransmitLoadPercentage float64        `db:"transmit_load_percentage"`
	HostGroup              sql.NullString `db:"host_group"`
}

//Response Proxies Model
type ProxiesJson struct {
	CountryCode    string      `json:"country_code"`
	CountryName    string      `json:"country_name"`
	Region         string      `json:"region"`
	OriginalRegion string      `json:"original_region"`
	Mode           int         `json:"mode"`
	Streaming      int         `json:"streaming"`
	Lock           bool        `json:"lock"`
	Icon           ProxyIcon   `json:"icon"`
	Nodes          []ProxyNode `json:"nodes"`
}

//Proxy Server Icon
type ProxyIcon struct {
	X1 string `json:"x1"`
	X2 string `json:"x2"`
	X3 string `json:"x3"`
}

//Proxy Server Node
type ProxyNode struct {
	Ip               string `json:"ip"`
	Path             string `json:"path"`
	Port             string `json:"port"`
	OvpnPort         string `json:"ovpn_port"`
	Schema           string `json:"schema"`
	Psk              string `json:"psk"`
	RemoteIdentifier string `json:"remote_identifier"`
}

//Store proxies pac results to used response format
func PrepareProxiesPacResults(store *store.Store, proxy *Proxies, country *Countries) ProxiesJson {
	var nodes []ProxyNode
	nodes = append(nodes, ProxyNode{
		Ip:               proxy.Ip,
		Path:             proxy.Endpoint,
		Port:             proxy.Port,
		OvpnPort:         proxy.OvpnPort,
		Schema:           proxy.Schema,
		Psk:              proxy.Psk.String,
		RemoteIdentifier: proxy.RemoteIdentifier.String,
	})
	return ProxiesJson{
		CountryCode:    country.CountryCode,
		CountryName:    country.Name,
		Region:         proxy.Region.String,
		OriginalRegion: proxy.Region.String,
		Mode:           proxy.Premium,
		Streaming:      proxy.Streaming,
		Lock:           false,
		Icon: ProxyIcon{
			X1: store.GetConfig().Site.ManageUrl + "/images/flags/" + country.CountryCode + ".imageset/" + country.CountryCode + ".png",
			X2: store.GetConfig().Site.ManageUrl + "/images/flags/" + country.CountryCode + ".imageset/" + country.CountryCode + "@2x.png",
			X3: store.GetConfig().Site.ManageUrl + "/images/flags/" + country.CountryCode + ".imageset/" + country.CountryCode + "@3x.png",
		},
		Nodes: nodes,
	}
}

//Store proxies ping results to used response format
func PrepareProxiesPingResults(store *store.Store, proxies *[]Proxies) []ProxiesJson {
	var proxiesJson []ProxiesJson
	for _, proxy := range *proxies {
		var country Countries
		err := store.GetConnection().Get(&country, "select * from countries where id = ? limit 1", proxy.CountryId)
		if err != nil {
			panic(err)
		}
		var nodes []ProxyNode
		proxiesJson = append(proxiesJson, ProxiesJson{
			CountryCode:    country.CountryCode,
			CountryName:    country.Name,
			Region:         proxy.Region.String,
			OriginalRegion: proxy.Region.String,
			Mode:           proxy.Premium,
			Streaming:      proxy.Streaming,
			Lock:           false,
			Icon: ProxyIcon{
				X1: store.GetConfig().Site.ManageUrl + "/images/flags/" + country.CountryCode + ".imageset/" + country.CountryCode + ".png",
				X2: store.GetConfig().Site.ManageUrl + "/images/flags/" + country.CountryCode + ".imageset/" + country.CountryCode + "@2x.png",
				X3: store.GetConfig().Site.ManageUrl + "/images/flags/" + country.CountryCode + ".imageset/" + country.CountryCode + "@3x.png",
			},
			Nodes: nodes,
		})
	}
	return proxiesJson
}

package metadata

import (
	"github.com/denverdino/aliyungo/metadata"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/metadata",

	"DNS_NAMESERVERS": metadata.DNS_NAMESERVERS,
	"EIPV4":           metadata.EIPV4,
	"ENDPOINT":        metadata.ENDPOINT,
	"HOSTNAME":        metadata.HOSTNAME,
	"IMAGE_ID":        metadata.IMAGE_ID,
	"INSTANCE_ID":     metadata.INSTANCE_ID,
	"MAC":             metadata.MAC,
	"META_VERSION_LATEST": metadata.META_VERSION_LATEST,
	"NETWORK_TYPE":        metadata.NETWORK_TYPE,
	"NTP_CONF_SERVERS":    metadata.NTP_CONF_SERVERS,
	"OWNER_ACCOUNT_ID":    metadata.OWNER_ACCOUNT_ID,
	"PRIVATE_IPV4":        metadata.PRIVATE_IPV4,
	"REGION":              metadata.REGION,
	"RS_TYPE_META_DATA":   metadata.RS_TYPE_META_DATA,
	"RS_TYPE_USER_DATA":   metadata.RS_TYPE_USER_DATA,
	"SERIAL_NUMBER":       metadata.SERIAL_NUMBER,
	"SOURCE_ADDRESS":      metadata.SOURCE_ADDRESS,
	"VPC_CIDR_BLOCK":      metadata.VPC_CIDR_BLOCK,
	"VPC_ID":              metadata.VPC_ID,
	"VSWITCH_CIDR_BLOCK":  metadata.VSWITCH_CIDR_BLOCK,
	"VSWITCH_ID":          metadata.VSWITCH_ID,
	"ZONE":                metadata.ZONE,

	"metaData": metadata.NewMetaData,
	"Request":  qlang.StructOf((*metadata.Request)(nil)),
}

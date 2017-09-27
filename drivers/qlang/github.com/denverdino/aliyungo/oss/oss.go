package oss

import (
	"github.com/denverdino/aliyungo/oss"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/denverdino/aliyungo/oss",

	"APNorthEast1":       oss.APNorthEast1,
	"APSouthEast1":       oss.APSouthEast1,
	"APSouthEast2":       oss.APSouthEast2,
	"AuthenticatedRead":  oss.AuthenticatedRead,
	"Beijing":            oss.Beijing,
	"BucketOwnerFull":    oss.BucketOwnerFull,
	"BucketOwnerRead":    oss.BucketOwnerRead,
	"DefaultContentType": oss.DefaultContentType,
	"DefaultRegion":      oss.DefaultRegion,
	"EUCentral1":         oss.EUCentral1,
	"Hangzhou":           oss.Hangzhou,
	"HeaderOSSPrefix":    oss.HeaderOSSPrefix,
	"Hongkong":           oss.Hongkong,
	"MEEast1":            oss.MEEast1,
	"Private":            oss.Private,
	"PublicRead":         oss.PublicRead,
	"PublicReadWrite":    oss.PublicReadWrite,
	"Qingdao":            oss.Qingdao,
	"Shanghai":           oss.Shanghai,
	"Shenzhen":           oss.Shenzhen,
	"USEast1":            oss.USEast1,
	"USWest1":            oss.USWest1,
	"Zhangjiakou":        oss.Zhangjiakou,
	"ACL":                func(v string) oss.ACL { return oss.ACL(v) },
	"Region":             func(v string) oss.Region { return oss.Region(v) },

	"authenticateCallBack": oss.AuthenticateCallBack,
	"setAttemptStrategy":   oss.SetAttemptStrategy,
	"setListMultiMax":      oss.SetListMultiMax,
	"setListPartsMax":      oss.SetListPartsMax,

	"AccessControlPolicy":       qlang.StructOf((*oss.AccessControlPolicy)(nil)),
	"Bucket":                    qlang.StructOf((*oss.Bucket)(nil)),
	"BucketInfo":                qlang.StructOf((*oss.BucketInfo)(nil)),
	"newOSSClient":              oss.NewOSSClient,
	"newOSSClientForAssumeRole": oss.NewOSSClientForAssumeRole,
	"CopyObjectResult":          qlang.StructOf((*oss.CopyObjectResult)(nil)),
	"CopyOptions":               qlang.StructOf((*oss.CopyOptions)(nil)),
	"Delete":                    qlang.StructOf((*oss.Delete)(nil)),
	"ErrorDocument":             qlang.StructOf((*oss.ErrorDocument)(nil)),
	"GetBucketInfoResp":         qlang.StructOf((*oss.GetBucketInfoResp)(nil)),
	"GetLocationResp":           qlang.StructOf((*oss.GetLocationResp)(nil)),
	"GetServiceResp":            qlang.StructOf((*oss.GetServiceResp)(nil)),
	"IndexDocument":             qlang.StructOf((*oss.IndexDocument)(nil)),
	"Key":                       qlang.StructOf((*oss.Key)(nil)),
	"ListResp":                  qlang.StructOf((*oss.ListResp)(nil)),
	"Multi":                     qlang.StructOf((*oss.Multi)(nil)),
	"Object":                    qlang.StructOf((*oss.Object)(nil)),
	"Options":                   qlang.StructOf((*oss.Options)(nil)),
	"Owner":                     qlang.StructOf((*oss.Owner)(nil)),
	"Part":                      qlang.StructOf((*oss.Part)(nil)),
	"RedirectAllRequestsTo":     qlang.StructOf((*oss.RedirectAllRequestsTo)(nil)),
	"RoutingRule":               qlang.StructOf((*oss.RoutingRule)(nil)),
	"WebsiteConfiguration":      qlang.StructOf((*oss.WebsiteConfiguration)(nil)),
}

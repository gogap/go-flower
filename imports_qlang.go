package main

import (
	"qlang.io/cl/qlang"

	_ "github.com/gogap/flow-contrib/handler/qlang"

	qlcdn "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/cdn"
	qlcms "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/cms"
	qlcommon "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/common"
	qlcrm "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/crm"
	qlcs "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/cs"
	qldm "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/dm"
	qldns "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/dns"
	qlecs "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/ecs"
	qless "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/ess"
	qllocation "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/location"
	qlmetadata "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/metadata"
	qlmns "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/mns"
	qlmq "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/mq"
	qlnas "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/nas"
	qlopensearch "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/opensearch"
	qloss "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/oss"
	qlpush "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/push"
	qlram "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/ram"
	qlrds "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/rds"
	qlslb "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/slb"
	qlsls "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/sls"
	qlsms "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/sms"
	qlsts "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/sts"
	qlutil "github.com/gogap/go-flower/drivers/qlang/github.com/denverdino/aliyungo/util"

	qlstructs "github.com/gogap/go-flower/drivers/qlang/github.com/fatih/structs"
)

func init() {
	qlang.Import("aliyun_cdn", qlcdn.Exports)
	qlang.Import("aliyun_cms", qlcms.Exports)
	qlang.Import("aliyun_common", qlcommon.Exports)
	qlang.Import("aliyun_crm", qlcrm.Exports)
	qlang.Import("aliyun_cs", qlcs.Exports)
	qlang.Import("aliyun_dm", qldm.Exports)
	qlang.Import("aliyun_dns", qldns.Exports)
	qlang.Import("aliyun_ecs", qlecs.Exports)
	qlang.Import("aliyun_ess", qless.Exports)
	qlang.Import("aliyun_location", qllocation.Exports)
	qlang.Import("aliyun_metadata", qlmetadata.Exports)
	qlang.Import("aliyun_mns", qlmns.Exports)
	qlang.Import("aliyun_mq", qlmq.Exports)
	qlang.Import("aliyun_nas", qlnas.Exports)
	qlang.Import("aliyun_opensearch", qlopensearch.Exports)
	qlang.Import("aliyun_oss", qloss.Exports)
	qlang.Import("aliyun_push", qlpush.Exports)
	qlang.Import("aliyun_ram", qlram.Exports)
	qlang.Import("aliyun_rds", qlrds.Exports)
	qlang.Import("aliyun_slb", qlslb.Exports)
	qlang.Import("aliyun_sls", qlsls.Exports)
	qlang.Import("aliyun_sms", qlsms.Exports)
	qlang.Import("aliyun_sts", qlsts.Exports)
	qlang.Import("aliyun_util", qlutil.Exports)
	qlang.Import("structs", qlstructs.Exports)
}

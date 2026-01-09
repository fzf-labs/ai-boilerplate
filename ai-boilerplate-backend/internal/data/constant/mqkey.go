package constant

import (
	"github.com/fzf-labs/kratos-contrib/pkg/mq"
)

var mqKey = mq.NewMessageConfigManager(mq.MQTypeAsynq)

var MQTest = mqKey.Register(&mq.MessageConfig{
	Key: "MQ_TEST",
	Metadata: map[mq.MetaKey]string{
		mq.MetaKeyAsynqQueue: "MQ_TEST",
	},
})

package service

import (
	"math/rand"
	"time"
)


var defaultAvatars = []string{
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/zoe.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/william.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/walter.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/thomas.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/taylor.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/sophia.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/sam.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/ryan.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/ruby.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/quinn.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/paul.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/owen.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/olivia.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/norman.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/nora.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/natalie.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/naomi.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/miley.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/mike.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/lucas.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/kylie.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/julia.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/joshua.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/john.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/jane.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/jackson.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/ivy.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/isaac.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/henry.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/harry.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/harold.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/hanna.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/grace.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/george.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/freddy.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/frank.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/finn.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/emma.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/emily.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/edward.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/clara.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/claire.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/chloe.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/audrey.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/arthur.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/anna.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/andy.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/alfred.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/alexa.png",
	"https://mozde.oss-cn-hangzhou.aliyuncs.com/public/avatar/default/abigail.png",
}

func GetRandomAvatar() string {
	rand.Seed(time.Now().UnixMicro())
	return defaultAvatars[rand.Intn(len(defaultAvatars))]
}

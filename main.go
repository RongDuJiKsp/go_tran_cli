package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	terrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
	"os"
)

func main() {
	client := createTencentClient()
	req := createRequestByCommand()
	res, err := client.TextTranslate(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(*res.Response.TargetText)
}
func createTencentClient() *tmt.Client {
	credential := common.NewCredential(
		os.Getenv("TCC_SECRET_ID"),
		os.Getenv("TCC_SECRET_KEY"),
	)
	client, err := tmt.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	if err != nil {
		handleErr(err)
	}
	return client
}
func createRequestByCommand() *tmt.TextTranslateRequest {
	req := tmt.NewTextTranslateRequest()
	p := int64(88)
	req.ProjectId = &p
	req.SourceText = flag.String("text", "", "需要翻译的文本")
	req.Source = flag.String("from", "en", "需要翻译的源语言，默认为en")
	req.Target = flag.String("lang", "zh", usageOfLangParam())
	flag.Parse()
	return req
}
func handleErr(err error) {
	var e *terrors.TencentCloudSDKError
	if errors.As(err, &e) {
		fmt.Printf("An API error has returned: %s", e.Message)
	} else {
		panic(err)
	}
}
func usageOfLangParam() string {
	return `
目标语言，各源语言的目标语言支持列表如下

zh（简体中文）：zh-TW（繁体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）
zh-TW（繁体中文）：zh（简体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）
en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）
ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）
ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）
fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）
es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）
it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）
de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）
tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）
ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）
pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）
vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）
id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）
th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）
ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）
ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）
hi（印地语）：en（英语）

示例值：zh`
}

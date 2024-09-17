package main

import (
	"fmt"
	"slices"
	"sort"
	"time"
)

type resource struct {
	key           string    // resourceのkey
	tagKeys       []string  // resourceに紐付けられているtagKey
	firstTaggedAt time.Time // 最初にタグが付けられた時間
}

type resourceTag struct {
	resourceKey string    // resourceのkey
	tagKey      string    // tagのkey
	createdAt   time.Time // タグの作成時間
}

// 問： 変数rTagsを、resourceKey毎に整理してresource型のスライスに整理してください
//
//	　重複タグは排除し、1番最初のタグが作られた時間をfirstTaggedAtに入れます。
//	　＊playgorundでは時間が固定ですが、実際に時間が経過している環境でも動作するよう実装してください。
//
// （下記、出力結果も参考にしてください）

func main() {
	rTags := []resourceTag{
		{
			resourceKey: "resource1",
			tagKey:      "tag1",
			createdAt:   time.Now(),
		},
		{
			resourceKey: "resource1",
			tagKey:      "tag2",
			createdAt:   time.Now().Add(30 * time.Minute),
		},
		{
			resourceKey: "resource1",
			tagKey:      "tag1",
			createdAt:   time.Now().Add(20 * time.Minute),
		},
		{
			resourceKey: "resource2",
			tagKey:      "tag2",
			createdAt:   time.Now(),
		},
		{
			resourceKey: "resource2",
			tagKey:      "tag3",
			createdAt:   time.Now().Add(10 * time.Minute),
		},
	}

	// rTagsをresourceKey毎のmapに変換
	resourceMap := make(map[string][]resourceTag)
	for _, rTag := range rTags {
		resourceMap[rTag.resourceKey] = append(resourceMap[rTag.resourceKey], rTag)
	}

	// resource型への変換
	var result []resource
	for key, resources := range resourceMap {
		// tagKeysの取得
		var tagKeys []string
		for _, r := range resources {
			if !slices.Contains(tagKeys, r.tagKey) { // 重複しているtagKeyは除外
				tagKeys = append(tagKeys, r.tagKey)
			}
		}

		// firstTaggedAtの取得
		sort.Slice(resources, func(i, j int) bool {
			return resources[i].createdAt.Before(resources[j].createdAt)
		})
		firstTaggedAt := resources[0].createdAt // 1番最初のタグが作られた時間

		// resourceの作成
		result = append(result, resource{
			key:           key,
			tagKeys:       tagKeys,
			firstTaggedAt: firstTaggedAt,
		})
	}

	fmt.Printf("%+v", result)

	// 出力結果
	// [{key:resource1 tagKeys:[tag1 tag2] firstTaggedAt:{wall:13454110244426743808 ext:1 loc:0x53c960}} {key:resource2 tagKeys:[tag2 tag3] firstTaggedAt:{wall:13454110244426743808 ext:1 loc:0x53c960}}]
}

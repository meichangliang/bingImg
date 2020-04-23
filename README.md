# bingImg

golang 实现获取 bing 每日美图,每日凌晨自动更新爬取

## 获取 cn.bing.com 官网当天的图片地址

URL

<https://api.meichangliang.com/getbingimg>

使用方法：

get 请求

返回的数据格式 `json`

```json
[
  "https://cn.bing.com/th?id=OHR.KingEider_ZH-CN3559595357_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.KauriTree_ZH-CN3695568740_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.GPS_ZH-CN5160095061_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.BluebellWood_ZH-CN8128422960_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.NeistPoint_ZH-CN3115403132_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.VernalFalls_ZH-CN2664125316_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.AlgonquinGrouse_ZH-CN2514966091_1920x1080.jpg",
  "https://cn.bing.com/th?id=OHR.NBNMSipapu_ZH-CN2293681419_1920x1080.jpg"
]
```

## 直接返回图片

请求地址

```url

https://api.meichangliang.com/bz

```

使用方法

```html
<img src="https://api.meichangliang.com/bz" alt="" />
```

效果如下图

![bing.com](https://api.meichangliang.com/bz "这里是标题 hover 后显示")

请求参数

| 属性 | 类型 | 必填 | 默认 |
| :--: | :--: | :--: | :--: |
| idx  | int  |  否  |  -1  |

> -1 则为随机返回某一张图片，0-7 则返回对应的壁纸, 例如 0 返回当前 cn.bing.com 的壁纸

使用方法

```html
<img src="https://api.meichangliang.com/bz?idx=0" alt="" />
<img src="https://api.meichangliang.com/bz?idx=1" alt="" />
```

效果如下图

![bing.com](https://api.meichangliang.com/bz?idx=0 "这里是标题 hover 后显示")

> 流量有限，希望大家可以省着点用

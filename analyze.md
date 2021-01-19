## 接口分析


### 请求 Header 信息
```JSON
"User-Agent:Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0",
"Accept-Language:zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2",
"Sign:error",
"DNT:1",
"TE:Trailers",
"Pragma:no-cache",
"Content-Type:application/json",
```

### 图片下载链接

Get 
下载地址
https://w.wallhaven.cc/full/q6/wallhaven-q6oxqq.jpg
https://w.wallhaven.cc/full/p8/wallhaven-p8g3le.jpg
https://w.wallhaven.cc/full/96/wallhaven-96gqz8.jpg
https://w.wallhaven.cc/full/r2/wallhaven-r2g7rm.jpg
https://w.wallhaven.cc/full/5d/wallhaven-5dekk8.jpg
https://w.wallhaven.cc/full/zm/wallhaven-zm9kpy.jpg
https://w.wallhaven.cc/full/wy/wallhaven-wye1qr.jpg
https://w.wallhaven.cc/full/96/wallhaven-96w8e8.png
https://w.wallhaven.cc/full/39/wallhaven-392zpd.jpg
https://w.wallhaven.cc/full/5w/wallhaven-5wkvz8.png 1920 × 1382

Get 
缩略图
https://th.wallhaven.cc/small/r2/r222om.jpg
https://th.wallhaven.cc/small/43/438w60.jpg
https://th.wallhaven.cc/small/5w/5wkvz8.jpg 300 x 200




### 获取图片列表:

Post https://api.zzzmh.cn/bz/getJsonByType
返回结果
json 数据
```JSON
{"msg":"success","result":{"current":1,"total":99976,"pages":834,"size":120,"records":[{"t":"p","x":3840,"i":"r2e391","y":2160},{"t":"j","x":4000,"i":"nkd1zm","y":2588},{"t":"j","x":2560,"i":"g8wvm7","y":1453},{"t":"j","x":2395,"i":"2ex21g","y":1597},{"t":"j","x":3840,"i":"nk1314","y":2160},{"t":"j","x":7680,"i":"j5yyww","y":4320},{"t":"j","x":2560,"i":"q6zr3q","y":1707},{"t":"j","x":3840,"i":"x1xy9l","y":2160},{"t":"j","x":3840,"i":"eolmrl","y":2160},{"t":"p","x":8185,"i":"73qmky","y":5787},{"t":"j","x":7952,"i":"xl3v9d","y":4480},{"t":"j","x":3840,"i":"vge72l","y":2160},{"t":"j","x":7475,"i":"x4g5qn","y":2615},{"t":"j","x":5071,"i":"r22221","y":3369},{"t":"j","x":2560,"i":"4yxpl7","y":1600}],"searchCount":true,"orders":[]},"code":0}
```
- records 中的参数内容 
> t : 图片的类型 [p(.png),j(.jpg)]  
> x,y : 图片的分辨率  
> i : 图片的编号  
- total : 图片的总数
- pages : 页面总数
- size : 每页的个数
- current : 当前页面


1. 筛选接口

请求参数
```JSON
{"sort":0,"color":-1,"dimensionX":null,"dimensionY":null,"fileType":-1,"createAt":-1,"category":-1,"target":"classify","pageNum":1}
```

2. 二次元接口

请求参数
```JSON
{"target":"anime","pageNum":1}
```

3. 人物接口

请求参数
```JSON
{"target":"people","pageNum":1}
```

4. 精选接口

请求参数
```JSON
{"target":"index","pageNum":1}
```
## 极简壁纸下载

[极简壁纸](https://bz.zzzmh.cn/) 官方的壁纸下载程序，批量下载。
## 网站接口分析

<details>
    <summary>图片列表</summary>
Post https://api.zzzmh.cn/bz/getJsonByType
</details>

<details>
    <summary>图片下载</summary>
Get https://w.wallhaven.cc/full/5w/wallhaven-5wkvz8.png 1920 × 1382   
 
Get https://w.wallhaven.cc/full/39/wallhaven-392zpd.jpg  
</details>

[接口详细内容](./analyze.md)

## 操作步骤

命令行执行  
> 可执行文件  
>
> - linux 当前文件下的 bg_jingjan 
> 
> - windows 当前文件下的 bg_jingjan.exe  
>
> - mac 当前文件下的 bg_jingjan.app  
>
> 不能执行可以运行 make build  

```shell script

$ ./bg_jijian.exe
-t 0, -o ./images, -s 1 , -a false, -v false
 可执行文件路径: E:\gowork\_github\bg-jijian\bg_jijian.exe:
  -a    是否下载后续的所有页面, 默认只下载一页
  -h    帮助信息
  -o string
        下载的图片保存路径 (default "./images")
  -s uint
        下载第几页的图片 (default 1)
  -t uint
        下载的图片的类型[0 : 二次元, 1 : 人物, 2 : 精选]
  -v    是否显示日志信息, 默认不显示
作者 : 这个程序不太乖 ,
 GitHub: https://github.com/getbuguai
 BiliBili: https://space.bilibili.com/278413353

```
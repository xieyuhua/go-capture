# gocapture

抓包以及流量统计,并在 web 进行可视化  _另外有 cli 分支，提供单文件打包的纯命令行版本抓包与流量统计工具_

运行前安装 winpcap（windows）或者 libpcap-dev（linux） linux改为静态链接libpcap库，不需要再安装libpcap了。

当前可视化默认支持本地的 8080 端口(如果使用其他端口以及修改城市经纬度，需要编辑 static/js/config.js)，
![gif](https://github.com/aoyouer/gocapture/raw/main/gif/CPT2106080056-800x385.gif)

通过 http://localhost:8080 访问流量地图

web 接口 /str 或者 /json 可以获取文本数据。

*自行在linux下编译的时候请先删掉nac.syso*

![image](https://user-images.githubusercontent.com/29120060/169734307-cb937321-378e-4db8-9661-a5ee8ff3c058.png)

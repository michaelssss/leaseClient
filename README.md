# Client端  
##  使用说明：
    go build main.go
    nohup ./main accessKey accessId severAccessKey serverIP:port
其中accessKey,accessId由阿里云处获取severAccessKey用于与Server通信校验，serverIp:port是Server的ip和端口用冒号连接  
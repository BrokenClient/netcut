# netcut
command-line tool for netcut.cn 为了便捷使用netcut.cn 网站而做的命令行工具


# 介绍
[netcut.cn](https://netcut.cn)\
最终程序有三种打开方式
1. netcut                  (输出网络剪切板的内容)
2. netcut message          (将message上传到网络剪切板)
3. ls/cat xx | netcut      (将管道传输的字符上传到网络剪切板)


# 开始前
根据抓包信息替换代码中的note_id值\
使用命令`go build netcut.go`\
最好`mv ./netcut /bin`方便使用(linux环境)

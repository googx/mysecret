# 说明

此项目公开 存储 一些加密的数据,放在 github 当中 ,使用 `gpg2` 加密该数据  


加密流程: 

1. 在 指定的git 库中 明文 存储 相关的数据 并提交 变更 
2. 运行脚本 将 git库 进行 tar 压缩文件. 并将 按照每一个压缩分片文件最大为 1MB 进行分片   
3. 使用 `gpg2` 进行对每一个分片文件进行加密, 加密后 保存为 普通文本文件 进行在此库中存储 


有一些特殊的文件 不需要在 git 库中存储,而是 直接加密保存在 目录中的,例如密钥什么的 ,方便 程序脚本中该目录中 读取
加密的文件,并在脚本中 直接进行解密读取访问 


通过 gpg2 来解密 磁盘
gpg2 --decrypt --quiet -r <gpg-key-id>  <encrypt-file> | sudo cryptsetup open <block-device-path> <vulumn name>

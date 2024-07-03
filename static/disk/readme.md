此目录记录 硬盘分区加密的 历史变化记录

此目录规则为: disk/{此硬盘序列号}/{加密分区的uuid}

要获取此硬盘密钥信息需要通过以下几个步骤:

0. 用 `key/master.gpg.key` 解密出 [aes_disk_key.gpg.json](aes_disk_key.gpg.json) 中的信息,拿到 `aes_disk.gpg.key` 密钥
1. 通过 `aes_disk.gpg.key` 密钥解密指定磁盘的 `dat`文件,该 `dat`文件是一串加密的`json`.
2. 该 `json` 信息中 包含了 使用 指定 `gpg`密钥加密的 `gpg消息`
3. 使用 `gpg` 解密该 `gpg消息` 得到 硬盘的解密密钥.即可解密硬盘  
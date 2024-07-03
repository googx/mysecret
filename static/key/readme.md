# 相关流程示例

填写完以上问题后,按照以下流程,计算出正确的 `AES` 密钥即可获得 `GPG`主密钥

1. 获取项目的初始提交hash,得到 加密块的初始向量 `IV`
   > 1. 提取 项目初始提交hash
   > 2. 进行计算 md5sum,即为 `IV`
    ```shell
    export TMP_IV=$( echo -n "##$(git log --reverse --format=%H)##" | md5sum | cut -d" " -f1 )
    echo -n ${TMP_IV}
    ```
   > 输出以下结果:
   > e408c596e915c963fa833b3b3ece5c2f
2. 使用类似的算法,计算此文件的 `hash`,得到  `AES-256` 的加密密钥 `KEY`
   > 输出的AES密钥 用来解密 [master.gpg.key](master.gpg.key) 文件,得到 `GPG` 密钥
    ```shell
    export TMP_KEY=$(sha256sum master.key.md | cut -d' ' -f1)
    echo -n ${TMP_KEY}
    ```
   > 输出以下结果:
   > 4c693f1cdaa0f14f67271c72f7b9d957ec08d23893b3f03effcb8c6c5ebe8a1f
3. 使用 以上 计算得到的 `IV` 和 `KEY` 使用以下命令 解密 `master.gpg.key` 文件
   ```shell
      echo -n ${TMP_IV}
      echo -n ${TMP_KEY}
   
      openssl aes-256-cbc -d -a -nosalt \
      -K ${TMP_KEY} \
      -iv ${TMP_IV} \
      -in master.gpg.key \
      -out master.gpg.key.dec
   ```
4. 也可以用 以下命令 进行 加密文件
   ```shell
      echo -n ${TMP_IV}
      echo -n ${TMP_KEY}
   
      openssl aes-256-cbc -e -a -nosalt \
      -K ${TMP_KEY} \
      -iv ${TMP_IV} \
      -in master.gpg.key.dec \
      -out master.gpg.key.enc
   ```
5. 接着 就可以 使用 `gpg --import master.gpg.key.dec` 导入密钥了,注意此密钥还有常用密码进行基础保护.
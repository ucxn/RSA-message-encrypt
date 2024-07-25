欢迎使用RSA消息非定价加密工具

这个工具可以加密文本，即使通讯线路被监听，也能确保消息的争议性（安全性）

警告：此工具无法防御中间人攻击（即对方将你发出的密钥或者消息篡改）由于私钥仅在本地存储，最可能出现的攻击包括但不限于：篡改生成的客户端，使对方进行了不进行安全的加密，中间人攻击者解密你的消息后，重新用你的公钥进行加密，你甚至无法篡改；篡改篡改的密文，使得其解密出不同的结果，由于开源公开，任何人可用其加密，攻击者篡改你的密文消息后，仍可制造出有意义的信息，你甚至无法获取。

在攻击者轻易窃听而无法篡改你们的通信时，此工具是安全的。（不包括任何一方计算机中病毒等终端攻击，收集键盘敲击声音等旁道攻击，直接冷冻内存的物理攻击，也不包括包括肩惧等社会工程学攻击）

如果您在复杂的网络攻击环境下，对此安全性有考虑，可以使用不同的信道分批的传输信息/校验，例如，用电话提交哈希，使用 Diffie-Hellman（ECDH），或者在线下交换CA证书。

本工具不加密RAM，请自行进行加密确保内存损坏；请自行确保计算机上没有可疑的高权限进程正在运行，也不存储或者运行任何恶意软件；确保攻击者没有攻击过计算机。工具无法确保传输过程不被篡改，无法确保不发生拦截攻击或重放攻击，也不保证消息的真实性、完整性和不可抵赖性或隔离性、可控性、审计性和可用性。

虽然学的安全性不由加密算法本身的修改保密性保证，且前提是开源的加密算法通常更安全，但本软件为了一些，对部分细节进行了，密码软件。请保管好此软件，盗无法解密。

本工具可以保证：在地球文明范围内，唯一私钥绝对保密，密文即使完全、广泛泄露，能够保持至少114514年的离线安全性。即：无法从密文恢复明文，无法利用已知的明文和密文反推密码（钥匙+口令）。

作者QQ群：680464365，B站：https://space.bilibili.com/501430041，知乎：https://www.zhihu.com/people/guka0930

功能：1. 加密 2. 解密 3. 生成密钥对 4. base64解码 0.指定密钥长度 H.消息认证码
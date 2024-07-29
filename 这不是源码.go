由于本程序仍然需要跨平台优化，暂不开源，此文件仅用于标记代码类型

欢迎使用RSA消息非对称加密工具

此工具可以加密文本，即使通讯全程被监听，也能确保消息的私密性（security）

警告：此工具无法防范中间人攻击（即对方将你发出去的密钥或者消息篡改）由于私钥只在本地存储，最有可能出现的攻击包括但不限于：篡改生成的公钥，使对方进行了不安全的加密，中间人攻击者解密你的消息后，重新用你的公钥进行加密，你甚至无法察觉；篡改发送的密文，使得其解码出不同的结果，由于公钥公开，任何人可用其加密，攻击者篡改你的密文消息后，仍可制造出有意义的信息，你甚至无法察觉。

在攻击者单纯窃听而无法篡改你们的通信时，此工具是安全的。（不包括任何一方计算机中病毒等终端攻击，收集键盘敲击声音等旁信道攻击，直接冷冻内存的物理攻击，也不包括肩窥等社会工程学攻击）

如果你处于复杂的网络攻击环境下，对此工具的安全性有担忧，可以使用多种不同的信道分批次的传输信息/校验，例如，用电话传递hash，使用Diffie-Hellman（ECDH），或者在线下交换CA证书。

本工具不加密RAM，请自行进行加密或者确保内存完整性；请自行确保计算机上没有可疑的高权限进程正在运行，也不存储或者运行任何恶意软件；确保攻击者曾经没有物理攻击过计算机。本工具无法确保传输过程不被篡改，无法确保不发生拦截攻击或重放攻击，也不保证消息的真实性、完整性和不可抵赖性或隔离性、可控性、审计性和可用性。

虽然密码学的安全性不由加密算法本身的保密性保证，且开源的加密算法通常更安全，但本软件为了便于使用，对部分细节进行了一些修改，属于专有软件。请妥善保管好此软件，以避免将来无法解密。

本工具可以保证：在地球文明范围内，只要私钥绝对保密，密文即使完全、广泛泄露，能够保持至少114514年的离线安全性。即：无法从密文恢复明文，无法利用已知的明文和密文反推密码（密钥+口令）。

作者QQ群：680464365 B站：https://space.bilibili.com/501430041 知乎：https://www.zhihu.com/people/guka0930

功能：1. 加密 2.解密 3.生成密钥对 4.base64解密 0.指定密钥长度 H.消息认证码

本程序支持，加密解密，生成密钥对，base64解码，自定义加密长度（8-2^64）

本软件独创把pem格式编码成单行格式方便传输。本软件对于新手非常友好，一次一密容易生成，用完就换。接收方向对方提供公钥后可以一键用私钥解密，真正实现简单易用。任何有Windows基础并且熟练操作键盘的人都可以掌握。跨平台，可以在Linux上使用。

本软件采用编程库中的随机数生成器调用，Microsoft Windows crypo随机API，各种win32句柄，网络接口统计信息，电脑的内存容量和频率，电脑CPU型号，无线网卡型号（如果有），操作系统的小版本号等各种尽可能调用作为信息熵，极大提高随机性。

对于对称加密软件，建议直接参考我的文章
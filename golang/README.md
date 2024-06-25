# 标准库

1. archive/tar：实现对tar存档的访问
2. archive/zip：支持读取和写入zip存档
3. arena：提供了为GO值集合分配内存的能力，并可以一次安全地手动释放这些空间；
4. bufio：实现缓冲I/O，包装io.Reader和io.Writer对象，创建另一个对象；
5. builtin：提供了GO预定义标识符的文档；
6. bytes：实现字节片操作的函数；
7. compress/bzip2：实现bzip2解压缩；
8. compress/flate：实现RFC1951中描述的DEFLATE压缩数据格式；
9. compress/gzip：按照RFC1952中的规定，实现gzip格式压缩文件的读写；
10. container/heap：实现heap.Interface的任何类型提供堆操作；
11. container/list：实现了一个双链接列表；
12. container/ring：实现了环形链表的操作；
13. context：上下文定义类型，它跨API边界和进程之间携带截止提起、取消信号和其他请求范围值；
14. crypto：收集常用的加密常数；
15. crypto/aes：实现AES加密(以前称为RijnDeal)
16. crypto/boring：公开仅在使用Go+BoringCrypto构建时可用的函数；
17. crypto/cipher：实现了标准的分组密码模式，这些模式可以封装在低级分组密码实现中；
18. crypto/des：实现数据加密标准DES和三重数据加密算法TDEA；
19. crypto/hmac：实现了密钥散列消息认证码hmac
20. crypto/md5：实现RFC1321中定义的MD5哈希算法；
21. crypto/rand：实现密码安全的随机数生成器；
22. crypto/rsa：按照PKCS#1和RFC8017中的规定实现RSA加密；
23. crypto/sha1：实现RFC3174中定义的SHA-1哈希算法；
24. crypto/sha256：实现了FIPS180-4中定义的SHA224和SHA256散列算法；
25. crypto/tls：部分实现了RFC5246中规定的TLS1.2和RFC8446中指定的TLS1.3；
26. crypto/x509：实现X.509标准的自己；
27. database/sql：围绕SQL或类似SQL数据库提供通用接口；
28. database/sql/driver：定义要有包sql使用的数据库驱动程序实现的接口；
29. debug/buildinfo：buildinfo提供了对嵌入在Go的二进制文件中的信息的访问，这些信息是如何构建的；
30. debug/dwarf：dwarf提供对从可执行文件加载的dwarf调试信息的访问，如dwarf2.0标准中所定义；
31. debug/elf：elf实现对elf对象文件的访问；
32. debug/pe：pe实现对pe(MicrosoftWindows可移植可执行文件)文件的访问；
33. debug/plan9obj：plan9obj实现对Plan9 a.out对象文件的访问；
34. embed：提供对运行的Go程序中嵌入的文件的访问；
35. encoding/ascii85：实现btoa工具和Adobe的PostScript和PDF文档格式中使用的ascii85数据编码；
36. encoding/base32：实现RFC4648指定的base32编码；
37. encoding/base64：实现RFC4648指定的base64编码；
38. encoding/binary：实现数字和字节序列之间的简单转换以及变量的编码和解码；
39. encoding/csv：读取和写入逗号分割值CSV文件；
40. encoding/gob：管理gobs流编码器(发射器)和解码器(接收器)之间交换的二进制值；
41. encoding/hex：实现十六进制编码和解码；
42. encoding/json：实现RFC7159中定义的JSON编码和解码；
43. encoding/pem：实现PEM数据编码，该编码源自隐私增强邮件；
44. encoding/xml：实现一个简单的XML1.0解析器，他可以理解XML名称空间；
45. errors：实现处理错误的函数；
46. expvar：为公共变量(如服务器中的操作计数器)提供标准化接口；
47. flag：实现命令行标志解析；
48. fmt：使用类似C的printf和scanf的函数实现格式化的I/O；
49. go/ast：申明用于表示Go包的语法树类型；
50. go/build：收集有关Go软件包的信息；
51. go/build/constraint：实现构建约束行的解析和求值；
52. go/constant：实现表示非类型化Go常量及其相应操作的值；
53. go/doc：从Go AST中提取源代码文档；
54. go/comment：实现Go文档注释的解析和重新格式化，这些注释直接位于包、const、func、type或var的顶级申明之前；
55. go/format：实现go源的标准格式；
56. go/importer：提供对导入数据导入期的访问；
57. go/parser：实现Go源的解析器；
58. go/printer：实现AST节点的打印；
59. go/scanner：实现GO源文本的扫描程序；
60. go/token：定义表示go编程语言的词汇标记的常量和标记的基本操作；
61. go/types：申明数据类型并实现Go包的类型检查算法；
62. hash：为哈希函数提供的接口；
63. hash/crc32：实现32位循环冗余校验或CRC-32校验和；
64. hash/crc64：实现64位循环冗余校验或CRC-64校验和；
65. hash/fnv：实现了FNV-1和FNV-1a创建的非加密散列函数；
66. hash/maphash：提供字节序列的哈希函数；
67. html：提供了转移和取消html文本的功能；
68. html/template：实现了数据驱动的模板，用于生成HTML输出，防止代码注入；
69. image：实现基本的二维图像库
70. image/color：实现基本颜色库；
71. image/draw：提供图像合成功能；
72. image/gif：实现GIF图像解码器和编码器；
73. image/jpeg：实现JPEG图像解码器和编码器；
74. image/png：实现PNG图像解码器和编码器；
75. index/suffixarray：使用内存中的后缀数组在对数时间内实现子字符串搜索；
76. io：提供I/O原语的基本接口；
77. io/fs：定义文件系统的基本接口；
78. ioutil：实现一些I/O实用程序功能；
79. log：实现一个简单的日志记录包；
80. log/syslog：提供系统日志服务的简单界面；
81. math：提供基本常数和数学函数；
82. math/big：实现任意精度算术(大数字)；
83. math/bits：为预先申明的无符号整数类型实现位技术和操作函数；
84. math/cmplx：为复数提供基本常数和数学函数；
85. math/rand：实现不适合安全敏感工作的伪随机数生成器；
86. mime：实现MIME规范的部分；
87. mime/multipart：实现MIME多部分解析，如RFC2046中所定义；
88. mime/quoteprintable：实现RFC2045指定的引用可打印编码；
89. net：为网络I/O提供可移植的接口，包括TCP/IP、UDP、域名解析和Unix域套接字；
90. net/http：提供http客户端和服务器实现；
91. net/http/cgi：实现RFC3875中规定的CGI(公共网关接口);
92. net/http/cookiejar：实现符合内存RFC6265的http.Cookiejar;
93. net/http/fcgi：实现FastCGI协议；
94. net/http/httptest：提供用于http测试的实用程序；
95. net/http/httptrace：提供跟踪HTTP客户端请求中事件的机制；
96. net/http/httputil：提供了HTTP实用程序功能，补充了net/HTTP包中更常见的功能；
97. net/http/pprof：通过其HTTP服务器运行时分析数据，以pprof可视化工具预期的格式提供服务；
98. net/mail：实现邮件消息的解析；
99. net/netip：定义一个小值类型的IP地址类型；
100. net/rpc：通过网络或其他I/O连接访问对象的导出方法；
101. net/rpc/jsonrpc：为RPC包实现JSON-RPC1.0 ClientCode和ServerCodec；
102. net/smtp：实现RFC5321中定义的简单邮件传输协议；
103. net/textproto：以HTTP、NNTP和SMTP的形式实现对基于文本的请求/相应协议的通用支持；
104. net/url：解析URL并实现查询转义；
105. os：为操作系统功能提供独立于平台的接口；
106. os/exec：运行外部命令；
107. os/signal：实现对输入信号的访问；
108. os/user：允许按名称或id查找用户账户；
109. path：实现用于操作斜杠分割路径的实用程序例程；
110. path/filepath：实现实用程序例程，以与目标操作系统定义的文件路径兼容的方式操作文件名路径；
111. plugin：实现Go插件的加载和符号解析；
112. reflect：实现运行时反射，允许程序操作具有任意类型的对象；
113. regexp：实现正则表达式搜索；
114. regexp/syntax：将正则表达式解析为解析数，并将解析数编译成程序；
115. runtime：包含与GO运行时系统交互的操作，例如控制goroutine的函数；
116. runtime/cgo：包含对cgo工具生成的代码的运行时支持；
117. runtime/coverage：覆盖率数据操作；
118. runtime/debug：包含程序运行时自行调试的工具；
119. runtime/metrics：提供了一个稳定的接口来访问Go运行时导出的实现定义的度量；
120. runtime/pprof：以pprof可视化工具预期的格式写入运行时分析数据；
121. runtime/race：实现数据竞争检测逻辑；
122. runtime/trace：包含程序生成Go执行跟踪程序跟踪的工具；
123. sort：提供用于排序切片和用户定义集合的原语；
124. strconv：基本数据类型与字符串之间的转换；
125. strings：操作UTF-8编码字符串的简单函数；
126. sync：提供基本的同步原语，如互斥锁；
127. sync/atomic：提供用于实现同步算法的低级原子内存原语；
128. syscall：包含低级操作系统原语的接口，系统调用；
129. syscall/js：使用js/wasm体系结构时，可以访问WebAssembly主机环境；
130. testing：为Go包的自动测试提供支持；
131. testing/fstest：实现对文件系统的测试实现和用户的支持；
132. testing/iotest：主要用于测试的读写器；
133. testing/quick：实用程序函数以帮助进行黑盒测试；
134. text/scanner：为utf-8编码文本提供扫描仪和标记器；
135. text/tabwriter：实现写入筛选器，将输入中耳朵选项卡列转换为正确对齐的文本；
136. text/template：用于生成文本输出的数据驱动模板；
137. text/template/parse：为text/template和html/template定义的模板构建解析树；
138. time：提供测量和显示时间的功能；
139. time/tzdata：提供时区数据库的嵌入副本；
140. unicode：提供数据和函数来测试Unicode代码点的某些属性；
141. unicode/utf16：实现utf-16序列的编码和解码；
142. unicode/utf8：实现函数和常量以支持utf-8编码的文本；
143. unsafe：包含绕过Go程序类型安全的操作；


## 常用包

1. context
2. encoding/json
3. encoding/base64
4. encoding/hex
5. errors
6. flag
7. fmt
8. io
9. io/util
10. log
11. math
12. net
13. net/http
14. net/url
15. reflect
16. regexp
17. sort
18. strconv
19. strings
20. sync
21. time





















go部署合约

第一步, 安装go开发环境

        i、先做配置工作(Mac)
            vim ~/.zshrc  进入文件
                      写入一下四行,
                        export GOROOT="/usr/local/go"
                        export GOPATH="/Users/chengjian/go"  这里改成自己的目录
                        export GOBIN="$GOROOT/bin"
                        export PATH="$PATH:$GOBIN"
                      然后esc,:wq,回车保存
            source ~/.zshrc
        mac或者linux
        ii、注意设置go mod以及安装时的代理问题
        iii、配置 GOPROXY 环境变量
                export GOPROXY=https://goproxy.io,direct
            还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
                export GOPRIVATE=git.mycompany.com,github.com/my/private
        windows
        iiii、配置 GOPROXY 环境变量
            $env:GOPROXY = "https://goproxy.io,direct"
            还可以设置不走 proxy 的私有仓库或组，多个用逗号相隔（可选）
            $env:GOPRIVATE = "git.mycompany.com,github.com/my/private"
  
  
第二步, 安装go-ethereum. 

        i、git clone https://github.com/ethereum/go-ethereum.git
        ii、install solc and abigen(有点容易出bug)

第三步, 把sol合约文件用 solc命令生成abi文件. 

        i、solc --abi inbox.sol 
        ii、把inbox.sol copy到GOROOT下面的inbox文件夹中去

第四步, 把abi文件用 abigen命令生成go文件, 这样deploy.go可以导入.go的合约文件,来部署它. 

        i、abigen --abi=inbox_sol_inbox.abi --pkg=inbox --out=inbox.go
        ii、以上命令在GOROOT inbox文件夹中执行

第五步, 跑deploy.go, 获得合约地址和交易哈希 (我用的kovan,测试链的浏览器打不开, 只能具体调用看结果). 

        i、go run deploy.go
        i、deploy.go文件可以放在GOPATH src文件夹中

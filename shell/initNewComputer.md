# 安装依赖
```sh
apt-get install wget curl
```
# 1.安装neovim
```sh
echo "install neovim"
echo "download neovim"
wget "https://github.com/neovim/neovim/releases/download/nightly/nvim-linux64.tar.gz"
tar -xzvf nvim-linux64.tar.gz ./
mv nvim-linux64/ /usr/local/
ln -s /usr/local/nvim-linux64/bin/nvim /usr/loca/bin/
```

# 安装coc插件
```sh
echo "hosts 199.232.28.133 raw.githubusercontent.com" >> /etc/hosts
curl -fLo ~/.config/nvim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
# 如果网络问题上面的都不行，我们可以直接去`https://github.com/junegunn/vim-plug/blob/master/plug.vim`复制
# copy配置文件
mkdir -p ~/.config/nvim
cp ./init.vim ~/.config/nvim/init.vim
```

# 安装gotags
```sh
go get -u github.com/jstemmer/gotags
sudo apt install ctags
```

# golang 安装二进制工具
```sh
nvim -c 'GoInstallBinaries|q'
```
# 安装各种所需的coc组件
```sh
nvim -c 'CocInstall -sync coc-go coc-rls coc-html coc-css coc-json coc-rust-analyzer|q'
```

# rust安装组件
```sh
rustup component add rls rust-analysis rust-src
cp ./coc-settings.json ~/.config/nvim/coc-settings.json
```

# coc 需要用到nodejs，但是系统apt的node版本过低，要去官网下载安装
# 安装node
```
wget https://nodejs.org/dist/v14.16.0/node-v14.16.0-linux-x64.tar.xz
tar -xvJf node-v14.16.0-linux-x64.tar.xz
sudo mv node-v14.16.0-linux-x64 /usr/local/
sudo ln -s /usr/local/node-v14.16.0-linux-x64/bin/npm /usr/local/bin/
sudo ln -s /usr/local/node-v14.16.0-linux-x64/bin/node /usr/local/bin/
```


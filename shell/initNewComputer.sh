
apt-get install wget curl
# 1.安装neovim
echo "install neovim"
echo "download neovim"
wget "https://github.com/neovim/neovim/releases/download/nightly/nvim-linux64.tar.gz"
tar -xzvf nvim-linux64.tar.gz ./
mv nvim-linux64/ /usr/local/
ln -s /usr/local/nvim-linux64/bin/nvim /usr/loca/bin/

# 安装coc插件
echo "hosts 199.232.28.133 raw.githubusercontent.com" >> /etc/hosts
curl -fLo ~/.config/nvim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
# 如果网络问题上面的都不行，我们可以直接去`https://github.com/junegunn/vim-plug/blob/master/plug.vim`复制
# copy配置文件
mkdir -p ~/.config/nvim
cp ./init.vim ~/.config/nvim/init.vim

# golang 安装二进制工具
GoInstallBinaries
# rust安装组件
rustup component add rls rust-analysis rust-src

cp ./coc-settings.json ~/.config/nvim/coc-settings.json
# coc 需要用到nodejs，但是系统apt的node版本过低，要去官网下载安装

